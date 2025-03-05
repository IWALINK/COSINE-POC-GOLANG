const fs = require('fs');
const path = require('path');
// Use relative path to the ethers module
const ethers = require('../contracts/node_modules/ethers');

async function main() {
  // Connect to the local network
  const provider = new ethers.JsonRpcProvider('http://localhost:8545');

  // Use the first private key from Ganache after running ganache -d --db ./ganache_data because it has enough eth for a deployment in this case it is the one below
  const wallet = new ethers.Wallet('0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d', provider);
  
  console.log("Deploying contracts with the account:", wallet.address);
  const balance = await provider.getBalance(wallet.address);
  console.log("Account balance:", ethers.formatEther(balance));

  // Load contract artifacts with updated paths
  const loadArtifact = (name) => {
    const artifactPath = path.resolve(__dirname, `../contracts/artifacts/contracts/cosine/${name}/${name}.json`);
    console.log(`Loading artifact from: ${artifactPath}`);
    
    try {
      return JSON.parse(fs.readFileSync(artifactPath, 'utf8'));
    } catch (error) {
      console.error(`Error loading artifact ${name}: ${error.message}`);
      throw error;
    }
  };

  // Deploy CosineToken
  console.log("Deploying CosineToken...");
  const cosineTokenArtifact = loadArtifact('CosineToken');
  const CosineToken = new ethers.ContractFactory(
    cosineTokenArtifact.abi,
    cosineTokenArtifact.bytecode,
    wallet
  );
  const cosineToken = await CosineToken.deploy();
  await cosineToken.waitForDeployment();
  const cosineTokenAddress = await cosineToken.getAddress();
  console.log("CosineToken deployed to:", cosineTokenAddress);

  // Deploy WalletLinking
  console.log("\nDeploying WalletLinking...");
  const walletLinkingArtifact = loadArtifact('WalletLinking');
  const WalletLinking = new ethers.ContractFactory(
    walletLinkingArtifact.abi,
    walletLinkingArtifact.bytecode,
    wallet
  );
  const walletLinking = await WalletLinking.deploy();
  await walletLinking.waitForDeployment();
  const walletLinkingAddress = await walletLinking.getAddress();
  console.log("WalletLinking deployed to:", walletLinkingAddress);

  // Deploy CreditVerification
  console.log("\nDeploying CreditVerification...");
  const creditVerificationArtifact = loadArtifact('CreditVerification');
  const CreditVerification = new ethers.ContractFactory(
    creditVerificationArtifact.abi,
    creditVerificationArtifact.bytecode,
    wallet
  );
  const creditVerification = await CreditVerification.deploy(
    cosineTokenAddress,
    walletLinkingAddress
  );
  await creditVerification.waitForDeployment();
  const creditVerificationAddress = await creditVerification.getAddress();
  console.log("CreditVerification deployed to:", creditVerificationAddress);

  // Deploy Bridge
  console.log("\nDeploying Bridge...");
  const bridgeArtifact = loadArtifact('Bridge');
  const Bridge = new ethers.ContractFactory(
    bridgeArtifact.abi,
    bridgeArtifact.bytecode,
    wallet
  );
  const chainId = 31337; // Local chain ID
  const bridge = await Bridge.deploy(
    cosineTokenAddress,
    walletLinkingAddress,
    chainId
  );
  await bridge.waitForDeployment();
  const bridgeAddress = await bridge.getAddress();
  console.log("Bridge deployed to:", bridgeAddress);

  // Save the contract addresses
  console.log("\nContract Addresses:");
  console.log(`CosineToken: ${cosineTokenAddress}`);
  console.log(`WalletLinking: ${walletLinkingAddress}`);
  console.log(`CreditVerification: ${creditVerificationAddress}`);
  console.log(`Bridge: ${bridgeAddress}`);

  // Grant BRIDGE_ROLE to the Bridge contract
  console.log("\nGranting BRIDGE_ROLE to Bridge contract...");
  const cosineTokenContract = new ethers.Contract(
    cosineTokenAddress,
    cosineTokenArtifact.abi,
    wallet
  );
  const BRIDGE_ROLE = await cosineTokenContract.BRIDGE_ROLE();
  const grantRoleTx = await cosineTokenContract.grantRole(BRIDGE_ROLE, bridgeAddress);
  await grantRoleTx.wait();
  console.log("BRIDGE_ROLE granted to Bridge contract");

  // Save the addresses to a file for easy access
  const addresses = {
    CosineToken: cosineTokenAddress,
    WalletLinking: walletLinkingAddress,
    CreditVerification: creditVerificationAddress,
    Bridge: bridgeAddress
  };
  fs.writeFileSync(
    path.resolve(__dirname, '../contract_addresses.json'),
    JSON.stringify(addresses, null, 2)
  );
  console.log("\nContract addresses saved to contract_addresses.json");

  console.log("\nDeployment complete!");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });