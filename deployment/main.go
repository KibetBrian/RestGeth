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

func main (){
	utils.ConfigureEnv()

	wallet := core.DecryptKeystore(os.Getenv("wallet1passphrase"), 1)
	
	client,ctx := core.Connect()

	nonce := utils.GetPendingNonce(client, ctx, wallet.PublicAddress)

	gasPrice := utils.SuggestGasPrice(client, ctx)

	networkId := utils.GetNetworkId(client, ctx)

	auth, err := bind.NewKeyedTransactorWithChainID(wallet.PrivateKey,networkId)
	if err != nil {
		log.Fatalf("Failed to create transaction signer: %v\n", err)
	}
	auth.Nonce=big.NewInt(int64(nonce))
	auth.Value=big.NewInt(0)
	auth.GasLimit=uint64(3000000)
	auth.GasPrice=gasPrice

	address, transaction, election, err := election.DeployElection(auth,client)

	if err != nil{
		log.Fatalf("Failde to deploy the contract: %v\n", err)
	}
	log.Println("Contract Address: ",address)
	log.Println("Contact Transaction: ",transaction)
	log.Println("Contract Election: ",election)
}