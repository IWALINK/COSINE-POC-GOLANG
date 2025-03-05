
const fs = require('fs');
const path = require('path');
const solc = require('solc');

// Project directories
const ROOT_DIR = __dirname;
const CONTRACTS_DIR = path.join(ROOT_DIR, 'cosine');
const INTERFACES_DIR = path.join(ROOT_DIR, 'interfaces');
const BUILD_DIR = path.join(ROOT_DIR, 'artifacts/contracts');
const NODE_MODULES_DIR = path.join(ROOT_DIR, 'node_modules');

// Create directories if they don't exist
if (!fs.existsSync(BUILD_DIR)) {
  fs.mkdirSync(BUILD_DIR, { recursive: true });
}

// Debug directory structure
console.log('Directory structure:');
console.log(`ROOT_DIR: ${ROOT_DIR}`);
console.log(`CONTRACTS_DIR: ${CONTRACTS_DIR}`);
console.log(`INTERFACES_DIR: ${INTERFACES_DIR}`);

// Find all Solidity files in our contracts and interfaces directories
function findSolidityFiles() {
  const files = [];
  
  // Function to recursively find .sol files in a directory
  function findInDir(dir, baseDir) {
    if (!fs.existsSync(dir)) {
      console.log(`Directory does not exist: ${dir}`);
      return;
    }
    
    const items = fs.readdirSync(dir);
    for (const item of items) {
      const itemPath = path.join(dir, item);
      const stats = fs.statSync(itemPath);
      
      if (stats.isDirectory()) {
        findInDir(itemPath, baseDir);
      } else if (stats.isFile() && item.endsWith('.sol')) {
        // Use a path relative to the baseDir
        const relativePath = path.relative(baseDir, itemPath);
        files.push({
          fullPath: itemPath,
          relativePath: relativePath
        });
      }
    }
  }
  
  // Find in contracts directory - store relative to contract dir
  findInDir(CONTRACTS_DIR, ROOT_DIR);
  
  // Find in interfaces directory - store relative to root dir
  findInDir(INTERFACES_DIR, ROOT_DIR);
  
  return files;
}

// Custom import resolver with debug logging
function findImports(importPath) {
  console.log(`Resolving import: ${importPath}`);
  
  try {
    // List of possible locations to check for the file
    const possibleLocations = [
      // Check if it's a node_modules path
      importPath.startsWith('@') ? path.join(NODE_MODULES_DIR, importPath) : null,
      
      // Check relative to ROOT_DIR
      path.join(ROOT_DIR, importPath),
      
      // Check relative to CONTRACTS_DIR
      path.join(CONTRACTS_DIR, importPath),
      
      // Check in INTERFACES_DIR directly for interface imports
      importPath.includes('ICosine') ? path.join(INTERFACES_DIR, path.basename(importPath)) : null,
      
      // Check with different directory structure assumptions
      path.join(ROOT_DIR, '..', importPath),
      path.join(CONTRACTS_DIR, '..', importPath)
    ].filter(Boolean); // Remove null entries
    
    // Try each location
    for (const location of possibleLocations) {
      console.log(`Trying location: ${location}`);
      if (fs.existsSync(location)) {
        console.log(`Found at: ${location}`);
        return { contents: fs.readFileSync(location, 'utf8') };
      }
    }
    
    throw new Error(`Could not find ${importPath} in any of the search locations`);
  } catch (e) {
    console.error(`Error resolving import ${importPath}: ${e.message}`);
    return { error: `Error resolving import ${importPath}: ${e.message}` };
  }
}

// List all .sol files
const solidityFiles = findSolidityFiles();
console.log('Found Solidity files:');
solidityFiles.forEach(file => console.log(`  ${file.relativePath}`));

// Prepare the sources input for the solc compiler
const sources = {};
for (const file of solidityFiles) {
  sources[file.relativePath] = { content: fs.readFileSync(file.fullPath, 'utf8') };
}

// Configure the solc input
const input = {
  language: 'Solidity',
  sources,
  settings: {
    outputSelection: {
      '*': {
        '*': ['abi', 'evm.bytecode', 'evm.deployedBytecode', 'evm.methodIdentifiers']
      }
    },
    optimizer: {
      enabled: true,
      runs: 200
    }
  }
};

// Compile the solidity files
console.log('Compiling contracts...');
const output = JSON.parse(solc.compile(JSON.stringify(input), { import: findImports }));

// Check for errors
if (output.errors) {
  const hasError = output.errors.some(error => error.severity === 'error');
  if (hasError) {
    console.error('Failed to compile:');
    console.error(output.errors.map(err => err.formattedMessage).join('\n'));
    process.exit(1);
  } else {
    console.warn('Compilation warnings:');
    console.warn(output.errors.map(err => err.formattedMessage).join('\n'));
  }
}

// Write the output to files
console.log('Writing compiled contracts to artifacts directory...');
for (const sourceName in output.contracts) {
  const contracts = output.contracts[sourceName];
  
  // Create directory for the source file
  const sourcePathParts = sourceName.split('/');
  const fileName = sourcePathParts[sourcePathParts.length - 1];
  const contractName = fileName.replace('.sol', '');
  
  const outputDir = path.join(BUILD_DIR, contractName);
  
  if (!fs.existsSync(outputDir)) {
    fs.mkdirSync(outputDir, { recursive: true });
  }
  
  for (const contractName in contracts) {
    const contract = contracts[contractName];
    const contractFileName = path.join(outputDir, `${contractName}.json`);
    
    // Create the artifact JSON
    const artifact = {
      contractName,
      abi: contract.abi,
      bytecode: contract.evm.bytecode.object,
      deployedBytecode: contract.evm.deployedBytecode.object,
      metadata: JSON.stringify({
        compiler: {
          version: `0.8.19+commit.7dd6d404`
        },
        language: 'Solidity',
        output: {
          abi: contract.abi,
          devdoc: {},
          userdoc: {}
        },
        settings: {
          compilationTarget: {
            [sourceName]: contractName
          },
          evmVersion: 'london',
          optimizer: {
            enabled: true,
            runs: 200
          },
          libraries: {}
        }
      }),
      methodIdentifiers: contract.evm.methodIdentifiers
    };
    
    fs.writeFileSync(contractFileName, JSON.stringify(artifact, null, 2));
    console.log(`Written: ${contractFileName}`);
  }
}

console.log('Compilation completed successfully!');