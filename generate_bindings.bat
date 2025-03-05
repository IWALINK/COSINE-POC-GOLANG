@echo off
echo Generating Go bindings for COSINE smart contracts...

:: Create directory (don't use -p flag on Windows)
if not exist pkg\smartcontracts\bindings mkdir pkg\smartcontracts\bindings

:: Use the extracted ABI and bytecode files
abigen --bin=contracts\abigen_input\CosineToken.bin --abi=contracts\abigen_input\CosineToken.abi --pkg=bindings --out=pkg\smartcontracts\bindings\cosine_token.go --type=CosineToken

abigen --bin=contracts\abigen_input\WalletLinking.bin --abi=contracts\abigen_input\WalletLinking.abi --pkg=bindings --out=pkg\smartcontracts\bindings\wallet_linking.go --type=WalletLinking

abigen --bin=contracts\abigen_input\CreditVerification.bin --abi=contracts\abigen_input\CreditVerification.abi --pkg=bindings --out=pkg\smartcontracts\bindings\credit_verification.go --type=CreditVerification

abigen --bin=contracts\abigen_input\Bridge.bin --abi=contracts\abigen_input\Bridge.abi --pkg=bindings --out=pkg\smartcontracts\bindings\bridge.go --type=Bridge

echo Go bindings generated successfully!