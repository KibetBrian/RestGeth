run:
	go run main.go
tidy:
	go mod tidy

#Compile to get bin and abi
compile:
	solc --bin --abi ./contracts/contract.sol -o ./build

#Generate golang code from the abi
abigengo:
	abigen --bin=./build/Election.bin --abi ./build/Election.abi --pkg=Election --out=generated/Election.go