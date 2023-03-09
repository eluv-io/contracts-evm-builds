## Description: Generate types and wrappers for Payment Spliter and ElvToken
VERSION=v0.0.1
mkdir -p ./dist/$VERSION

echo "Generating Commerce Types and Wrappers for Ethers v5"
echo "---------------------------------------------------"

abi-types-generator ../dist/v0.0.1/commerce/abi/Address.abi --output=./src/ --name=Address --provider=ethers_v5
abi-types-generator ../dist/v0.0.1/commerce/abi/Context.abi --output=./src/ --name=Context --provider=ethers_v5
abi-types-generator ../dist/v0.0.1/commerce/abi/IERC20.abi --output=./src/ --name=ERC20 --provider=ethers_v5
abi-types-generator ../dist/v0.0.1/commerce/abi/IERC20Permit.abi --output=./src/ --name=ERC20Permit --provider=ethers_v5
abi-types-generator ../dist/v0.0.1/commerce/abi/Payment.abi --output=./src/ --name=Payment --provider=ethers_v5
abi-types-generator ../dist/v0.0.1/commerce/abi/PaymentSplitter.abi --output=./src/ --name=PaymentSplitter --provider=ethers_v5
abi-types-generator ../dist/v0.0.1/commerce/abi/SafeERC20.abi --output=./src/ --name=SafeERC20 --provider=ethers_v5

abi-types-generator ../../elv-live-js/contracts/v4/ElvToken.abi --output=./src/ --name=ElvToken --provider=ethers_v5

## TODO
# generate v3 Type and Wrappers

## copy to dist
cp ./src/* ./dist/$VERSION

