package main

import (
	"log"
	"math/big"
	"os"

	"github.com/KibetBrian/geth/core"
	"github.com/KibetBrian/geth/election"
	"github.com/KibetBrian/geth/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const (
	address = "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B"
)

func main() {
	utils.ConfigureEnv()
	
	wallet := core.DecryptKeystore(os.Getenv("wallet1passphrase"),1)

	client, ctx := core.Connect()

	netId := utils.GetNetworkId(client, ctx)

	nonce := utils.GetPendingNonce(client, ctx, wallet.PublicAddress)

	gasPrice := utils.SuggestGasPrice(client, ctx)

	contractAddress :=utils.StringToAddress("0x964a58e67f6070131932cc060f285ac88a9a32e3")
	election, err := election.NewElection(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create election instance: %v\n", err)
	}

	//Transaction signer
	tx , err := bind.NewKeyedTransactorWithChainID(wallet.PrivateKey, netId)
	if err != nil{
		log.Fatalf("Failed to create Transaction Signer: %v\n", err)
	}
	tx.GasLimit=3000000
	tx.GasPrice=gasPrice
	tx.Nonce=big.NewInt(int64(nonce))

	//Transactor
	voterAddress := utils.StringToAddress("c50cd042edf5cf1d31cd5f9f04b499acd7e33a67")
	transaction, err := election.AddVoter(tx, voterAddress)
	if err != nil {
		log.Fatalf("Failed to add voter: %v", err)
	}

	//Caller
	voters, err := election.GetVoters(&bind.CallOpts{From: wallet.PublicAddress,Context: ctx})
	if err != nil{
		log.Printf("Failed to get voters. Err: %v\n", err)
	}
	log.Println("Voters: ", voters)
	log.Println("Transaction: ", transaction)
	
}
