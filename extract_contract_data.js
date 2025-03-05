const fs = require('fs');
const path = require('path');

// Contract names
const contracts = [
  'CosineToken',
  'WalletLinking',
  'CreditVerification',
  'Bridge'
];

// Directories
const artifactsDir = path.join(__dirname, 'contracts', 'artifacts', 'contracts', 'cosine');
const outputDir = path.join(__dirname, 'contracts', 'abigen_input');

// Create output directory if it doesn't exist
if (!fs.existsSync(outputDir)) {
  fs.mkdirSync(outputDir, { recursive: true });
}

// Extract ABI and bytecode for each contract
contracts.forEach(contract => {
  const artifactPath = path.join(artifactsDir, contract, `${contract}.json`);
  console.log(`Processing ${artifactPath}...`);
  
  try {
    const artifactData = JSON.parse(fs.readFileSync(artifactPath, 'utf8'));
    
    // Extract ABI
    const abiPath = path.join(outputDir, `${contract}.abi`);
    fs.writeFileSync(abiPath, JSON.stringify(artifactData.abi));
    
    // Extract bytecode 
    const binPath = path.join(outputDir, `${contract}.bin`);
    fs.writeFileSync(binPath, artifactData.bytecode);
    
    console.log(`Successfully extracted data for ${contract}`);
  } catch (error) {
    console.error(`Error processing ${contract}: ${error.message}`);
  }
});

console.log('Extraction complete!');