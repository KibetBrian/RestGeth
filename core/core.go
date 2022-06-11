package core

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/KibetBrian/geth/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//Connects client to the main-net or test-net
func Connect() (*ethclient.Client, context.Context){
	ctx := context.Background();
	networkUrl := os.Getenv("networkUrl")

	client, err := ethclient.DialContext(ctx, networkUrl)
	if err != nil {
		log.Fatalf("Error occured while dialing to the network. Err: %v\n", err)
	}

	return client, ctx
}


//This function gets the balance at a particular address
func GetBalance(address string) *big.Float{
	client, ctx := Connect()

	addr := common.HexToAddress(address)
	balanceInWei, err := client.BalanceAt(ctx,addr, nil)
	if err != nil {
		log.Printf("Failed to get balance: Err: %v", err)
	}

	balanceInEther := utils.WeiToEth(balanceInWei)
	return balanceInEther
}