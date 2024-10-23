Run this in root :
go mod init <your_module_name>
go get github.com/ethereum/go-ethereum@latest
go get github.com/joho/godotenv

Run this in root:
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
export PATH=$PATH:$(go env GOPATH)/bin
source ~/.zshrc  # or ~/.bashrc depending on the shell you're using
solc-select install <version>
solc-select use <version>

Run this in contracts:
solc --abi <contract_name>.sol -o build
solc --bin <contract_name>.sol -o build
abigen --bin=./build/<contract_name>.bin --abi=./build/<contract_name>.abi --pkg=<contract_name> --out=<contract_name>.go
