package core

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/KibetBrian/geth/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//Connects client to the main-net or test-net
func Connect() (*ethclient.Client, context.Context) {
	ctx := context.Background()
	networkUrl := os.Getenv("networkUrl")

	client, err := ethclient.DialContext(ctx, networkUrl)
	if err != nil {
		log.Fatalf("Error occured while dialing to the network. Err: %v\n", err)
	}

	return client, ctx
}

//This function gets the balance at a particular address
func GetBalance(address string) *big.Float {
	client, ctx := Connect()

	addr := common.HexToAddress(address)
	balanceInWei, err := client.BalanceAt(ctx, addr, nil)
	if err != nil {
		log.Printf("Failed to get balance: Err: %v\n", err)
	}

	balanceInEther := utils.WeiToEth(balanceInWei)
	return balanceInEther
}

func TranferEther(addr string, amount float64, password string) {
	client, ctx := Connect()
	Key := DecryptKeystore(password, 1)

	addressTo := common.HexToAddress(addr)
	addressFrom := Key.PublicAddress

	nonce, err := client.NonceAt(ctx, addressFrom, nil)
	if err != nil {
		log.Printf("Failed to get nonce at %v\n address. Err: %v\n", addr, err)
	}

	//Convert amount from int to wei
	amountInBigFloat := new(big.Float).SetFloat64(amount)
	amountInWei := utils.EthToWei(amountInBigFloat)

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("Failed to suggest the gas price. Err: %v\n", err)
	}
	tx := types.NewTransaction(nonce, addressTo, amountInWei, 21000, gasPrice, nil)

	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Printf("Failed to get chain id. Err %v\n", err)
	}

	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainId), Key.PrivateKey)
	if err != nil {
		log.Fatalf("Transaction failed. Err: %v\n", err)
	}

	err = client.SendTransaction(ctx, tx)
	if err != nil {
		log.Fatalf("Failed to send transaction to the pending pool of execution. Err: %v\n", err)
	}
	log.Println("Tx Completed.", tx.Hash().Hex())
}
