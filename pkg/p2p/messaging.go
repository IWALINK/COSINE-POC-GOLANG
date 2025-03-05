package p2p

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// MessageType represents the type of message being sent.
type MessageType string

// Message types
const (
	TypeVote          MessageType = "vote"
	TypeCreditScore   MessageType = "credit-score"
	TypeValidatorInfo MessageType = "validator-info"
	TypePing          MessageType = "ping"
	TypePong          MessageType = "pong"
)

// Message represents a message in the COSINE P2P network.
type Message struct {
	Type      MessageType         `json:"type"`
	Timestamp int64               `json:"timestamp"`
	Sender    string              `json:"sender"`
	Target    string              `json:"target,omitempty"`
	Payload   map[string]interface{} `json:"payload"`
}

// MessageHandler is a function that handles incoming messages.
type MessageHandler func(msg *Message, from peer.ID) error

// MessagingService manages P2P messaging for COSINE validators.
type MessagingService struct {
	network         *NetworkManager
	logger          *utils.Logger
	config          *utils.ConfigManager
	metrics         *utils.MetricsCollector
	pubsub          *pubsub.PubSub
	topics          map[string]*pubsub.Topic
	subscriptions   map[string]*pubsub.Subscription
	handlers        map[MessageType][]MessageHandler
	ctx             context.Context
	cancel          context.CancelFunc
	handlersMu      sync.RWMutex
	topicsMu        sync.RWMutex
	subscriptionsMu sync.RWMutex
}

// NewMessagingService creates a new P2P messaging service.
func NewMessagingService(
	network *NetworkManager,
	config *utils.ConfigManager,
	logger *utils.Logger,
) (*MessagingService, error) {
	if network == nil {
		return nil, fmt.Errorf("network manager is required")
	}
	if config == nil {
		return nil, fmt.Errorf("configuration manager is required")
	}
	if logger == nil {
		logger = utils.GetGlobalLogger()
	}

	ctx, cancel := context.WithCancel(context.Background())

	// Create messaging service instance
	service := &MessagingService{
		network:       network,
		logger:        logger.WithComponent("p2p-messaging"),
		config:        config,
		metrics:       utils.GetGlobalMetrics(),
		topics:        make(map[string]*pubsub.Topic),
		subscriptions: make(map[string]*pubsub.Subscription),
		handlers:      make(map[MessageType][]MessageHandler),
		ctx:           ctx,
		cancel:        cancel,
	}

	// Initialize pubsub
	if err := service.initPubSub(); err != nil {
		cancel()
		return nil, err
	}

	// Setup standard topics
	if err := service.setupStandardTopics(); err != nil {
		cancel()
		return nil, err
	}

	return service, nil
}

//FOR TESTS

// StringToPeerID converts a string to a peer.ID
func (ms *MessagingService) StringToPeerID(peerIDStr string) (peer.ID, error) {
	return peer.Decode(peerIDStr)
}

// initPubSub initializes the pubsub system.
func (ms *MessagingService) initPubSub() error {
    // Create new pubsub instance with gossipsub protocol
    ps, err := pubsub.NewGossipSub(ms.ctx, ms.network.GetHost())
    if err != nil {
        return fmt.Errorf("failed to create pubsub: %w", err)
    }

    ms.pubsub = ps
    ms.logger.Info("PubSub initialized with GossipSub protocol")
    return nil
}

// setupStandardTopics creates and subscribes to standard protocol topics.
func (ms *MessagingService) setupStandardTopics() error {
    // Default topics the service will subscribe to
    standardTopics := []string{
        "cosine/validator/votes",
        "cosine/validator/credit-scores",
        "cosine/validator/announcements",
    }

    // Join each topic and set up subscription
    for _, topicName := range standardTopics {
        if err := ms.JoinTopic(topicName); err != nil {
            return fmt.Errorf("failed to join topic %s: %w", topicName, err)
        }
    }

    // Set up a direct message protocol for peer-to-peer messages
    protocolID := protocol.ID("/cosine/1.0.0")
    ms.network.GetHost().SetStreamHandler(protocolID, ms.handleDirectMessage)

    return nil
}

// JoinTopic joins a pubsub topic and sets up subscription.
func (ms *MessagingService) JoinTopic(topicName string) error {
	ms.topicsMu.Lock()
	defer ms.topicsMu.Unlock()

	// Check if already joined
	if _, exists := ms.topics[topicName]; exists {
		return nil
	}

	// Join the topic
	topic, err := ms.pubsub.Join(topicName)
	if err != nil {
		return fmt.Errorf("failed to join topic: %w", err)
	}

	// Subscribe to the topic
	subscription, err := topic.Subscribe()
	if err != nil {
		return fmt.Errorf("failed to subscribe to topic: %w", err)
	}

	// Store topic and subscription
	ms.topics[topicName] = topic
	
	ms.subscriptionsMu.Lock()
	ms.subscriptions[topicName] = subscription
	ms.subscriptionsMu.Unlock()

	// Start message handling goroutine
	go ms.handleSubscription(topicName, subscription)

	ms.logger.Info("Joined topic", "topic", topicName)
	return nil
}

// LeaveTopic leaves a pubsub topic and cancels the subscription.
func (ms *MessagingService) LeaveTopic(topicName string) error {
	ms.topicsMu.Lock()
	defer ms.topicsMu.Unlock()

	// Check if joined
	topic, exists := ms.topics[topicName]
	if !exists {
		return nil
	}

	// Cancel subscription
	ms.subscriptionsMu.Lock()
	if subscription, exists := ms.subscriptions[topicName]; exists {
		subscription.Cancel()
		delete(ms.subscriptions, topicName)
	}
	ms.subscriptionsMu.Unlock()

	// Close the topic
	if err := topic.Close(); err != nil {
		return fmt.Errorf("failed to close topic: %w", err)
	}

	delete(ms.topics, topicName)
	ms.logger.Info("Left topic", "topic", topicName)
	return nil
}

// PublishToTopic publishes a message to a pubsub topic.
func (ms *MessagingService) PublishToTopic(topicName string, msg *Message) error {
	ms.topicsMu.RLock()
	topic, exists := ms.topics[topicName]
	ms.topicsMu.RUnlock()

	if !exists {
		return fmt.Errorf("not subscribed to topic %s", topicName)
	}

	// Set timestamp if not set
	if msg.Timestamp == 0 {
		msg.Timestamp = time.Now().UnixNano()
	}

	// Set sender if not set
	if msg.Sender == "" {
		msg.Sender = ms.network.GetHost().ID().String()
	}

	// Marshal message to JSON
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// Publish message
	if err := topic.Publish(ms.ctx, msgBytes); err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	// Update metrics
	ms.metrics.IncCounter("p2p_messages_published_total", string(msg.Type))

	ms.logger.Debug("Published message to topic",
		"topic", topicName,
		"type", msg.Type,
		"target", msg.Target)

	return nil
}

// SendDirect sends a direct message to a specific peer.
func (ms *MessagingService) SendDirect(peerID peer.ID, msg *Message) error {
	// Set timestamp if not set
	if msg.Timestamp == 0 {
		msg.Timestamp = time.Now().UnixNano()
	}

	// Set sender if not set
	if msg.Sender == "" {
		msg.Sender = ms.network.GetHost().ID().String()
	}

	// Marshal message to JSON
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// Open a stream to the peer
	stream, err := ms.network.GetHost().NewStream(ms.ctx, peerID, protocol.ID("/cosine/1.0.0"))
	if err != nil {
		return fmt.Errorf("failed to open stream to peer: %w", err)
	}
	defer stream.Close()

	// Write message length as a 4-byte prefix (simple length-prefixed protocol)
	msgLen := len(msgBytes)
	lenBuf := []byte{
		byte(msgLen >> 24),
		byte(msgLen >> 16),
		byte(msgLen >> 8),
		byte(msgLen),
	}

	if _, err := stream.Write(lenBuf); err != nil {
		return fmt.Errorf("failed to write message length: %w", err)
	}

	// Write message body
	if _, err := stream.Write(msgBytes); err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	// Update metrics
	ms.metrics.IncCounter("p2p_direct_messages_sent_total", string(msg.Type))

	ms.logger.Debug("Sent direct message",
		"peer", peerID.String(),
		"type", msg.Type,
		"target", msg.Target)

	return nil
}

// RegisterHandler registers a handler function for a specific message type.
func (ms *MessagingService) RegisterHandler(msgType MessageType, handler MessageHandler) {
	ms.handlersMu.Lock()
	defer ms.handlersMu.Unlock()

	ms.handlers[msgType] = append(ms.handlers[msgType], handler)
	ms.logger.Debug("Registered message handler", "type", msgType)
}

// handleSubscription processes messages from a subscription.
func (ms *MessagingService) handleSubscription(topicName string, subscription *pubsub.Subscription) {
	ms.logger.Debug("Starting message handler for topic", "topic", topicName)

	for {
		// Get the next message
		msg, err := subscription.Next(ms.ctx)
		if err != nil {
			if err == context.Canceled || err == context.DeadlineExceeded {
				ms.logger.Debug("Subscription handler stopped", "topic", topicName)
				return
			}
			ms.logger.Error("Failed to get next message from subscription",
				"topic", topicName,
				"error", err)
			continue
		}

		// Skip messages from self
		if msg.ReceivedFrom == ms.network.GetHost().ID() {
			continue
		}

		// Process the message
		go ms.processMessage(msg.Data, msg.ReceivedFrom)
	}
}

// handleDirectMessage handles incoming direct messages from peers.
func (ms *MessagingService) handleDirectMessage(stream network.Stream) {
    defer stream.Close()

    // Read message length (4-byte prefix)
    lenBuf := make([]byte, 4)
    if _, err := stream.Read(lenBuf); err != nil {
        ms.logger.Error("Failed to read message length", "error", err)
        return
    }

    // Calculate message length
    msgLen := int(lenBuf[0])<<24 | int(lenBuf[1])<<16 | int(lenBuf[2])<<8 | int(lenBuf[3])

    // Read message body
    msgBuf := make([]byte, msgLen)
    if _, err := stream.Read(msgBuf); err != nil {
        ms.logger.Error("Failed to read message body", "error", err)
        return
    }

    // Process the message
    ms.processMessage(msgBuf, stream.Conn().RemotePeer())
}

// processMessage processes an incoming message.
func (ms *MessagingService) processMessage(data []byte, from peer.ID) {
	// Parse the message
	var msg Message
	if err := json.Unmarshal(data, &msg); err != nil {
		ms.logger.Error("Failed to unmarshal message", "error", err)
		return
	}

	// Update metrics
	ms.metrics.IncCounter("p2p_messages_received_total", string(msg.Type))

	// Handle the message
	ms.handlersMu.RLock()
	handlers, exists := ms.handlers[msg.Type]
	ms.handlersMu.RUnlock()

	if !exists || len(handlers) == 0 {
		ms.logger.Debug("No handler registered for message type", "type", msg.Type)
		return
	}

	// Call all registered handlers
	for _, handler := range handlers {
		if err := handler(&msg, from); err != nil {
			ms.logger.Error("Message handler failed",
				"type", msg.Type,
				"error", err)
		}
	}
}

// BroadcastVote broadcasts a vote to the network.
func (ms *MessagingService) BroadcastVote(
	walletAddress string,
	voteValue int,
	voterReputation float64,
) error {
	msg := &Message{
		Type:   TypeVote,
		Target: walletAddress,
		Payload: map[string]interface{}{
			"value":      voteValue,
			"reputation": voterReputation,
		},
	}

	return ms.PublishToTopic("cosine/validator/votes", msg)
}

// BroadcastCreditScore broadcasts a credit score update to the network.
func (ms *MessagingService) BroadcastCreditScore(
	walletAddress string,
	score float64,
	reason string,
) error {
	msg := &Message{
		Type:   TypeCreditScore,
		Target: walletAddress,
		Payload: map[string]interface{}{
			"score":  score,
			"reason": reason,
		},
	}

	return ms.PublishToTopic("cosine/validator/credit-scores", msg)
}

// SendValidatorInfo sends validator information to a specific peer.
func (ms *MessagingService) SendValidatorInfo(
	peerID peer.ID,
	validatorID string,
	stake float64,
	performance float64,
) error {
	msg := &Message{
		Type: TypeValidatorInfo,
		Payload: map[string]interface{}{
			"validator_id": validatorID,
			"stake":        stake,
			"performance":  performance,
		},
	}

	return ms.SendDirect(peerID, msg)
}

// Ping sends a ping message to a peer and waits for pong response.
func (ms *MessagingService) Ping(peerID peer.ID) (time.Duration, error) {
	// Create ping message
	pingMsg := &Message{
		Type:      TypePing,
		Timestamp: time.Now().UnixNano(),
		Payload:   map[string]interface{}{},
	}

	// Create a channel to receive the pong response
	pongCh := make(chan *Message, 1)
	errCh := make(chan error, 1)

	// Register temporary handler for pong response
	ms.handlersMu.Lock()
	pongHandlers := ms.handlers[TypePong]
	ms.handlers[TypePong] = append(ms.handlers[TypePong], func(msg *Message, from peer.ID) error {
		if from == peerID {
			select {
			case pongCh <- msg:
			default:
			}
		}
		return nil
	})
	ms.handlersMu.Unlock()

	// Remove temporary handler after timeout
	defer func() {
		ms.handlersMu.Lock()
		ms.handlers[TypePong] = pongHandlers
		ms.handlersMu.Unlock()
	}()

	// Send ping message
	pingStart := time.Now()
	if err := ms.SendDirect(peerID, pingMsg); err != nil {
		return 0, fmt.Errorf("failed to send ping: %w", err)
	}

	// Wait for pong response with timeout
	select {
	case <-pongCh:
		return time.Since(pingStart), nil
	case err := <-errCh:
		return 0, err
	case <-time.After(10 * time.Second):
		return 0, fmt.Errorf("ping timeout")
	}
}

// Close shuts down the messaging service and all subscriptions.
func (ms *MessagingService) Close() error {
	ms.cancel()

	// Cancel all subscriptions
	ms.subscriptionsMu.Lock()
	for _, subscription := range ms.subscriptions {
		subscription.Cancel()
	}
	ms.subscriptions = make(map[string]*pubsub.Subscription)
	ms.subscriptionsMu.Unlock()

	// Close all topics
	ms.topicsMu.Lock()
	for topicName, topic := range ms.topics {
		if err := topic.Close(); err != nil {
			ms.logger.Error("Failed to close topic", "topic", topicName, "error", err)
		}
	}
	ms.topics = make(map[string]*pubsub.Topic)
	ms.topicsMu.Unlock()

	ms.logger.Info("Messaging service closed")
	return nil
}