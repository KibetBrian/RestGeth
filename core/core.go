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
		log.Printf("Failed to get balance: Err: %v\n", err)
	}

	balanceInEther := utils.WeiToEth(balanceInWei)
	return balanceInEther
}

func TranferEther(addr string, amount int, password string){
		client, ctx := Connect()
		
		address := common.HexToAddress(addr)
		addressFrom := common.HexToAddress("0xc50Cd042edf5CF1D31CD5F9F04b499acd7e33A67")

		nonce, err := client.NonceAt(ctx,addressFrom, nil)
		if err != nil{
			log.Printf("Failed to get nonce at %v\n address. Err: %v\n",addr, err)
		}
		
		//Convert amount from int to wei
		amountInBigInt := big.NewInt(int64(amount));
		amountInBigFloat := new(big.Float).SetInt(amountInBigInt)
		amountInWei := utils.EthToWei(amountInBigFloat)

		
		gasPrice, err := client.SuggestGasPrice(ctx)
		if err != nil{
			log.Printf("Failed to suggest the gas price. Err: %v\n",err)
		}
		tx := types.NewTransaction(nonce, address,amountInWei,21000,gasPrice,nil)

		privakeKey := DecryptKeystore(password,1).PrivateKey

		chainId, err := client.ChainID(ctx)
		if err != nil {
			log.Printf("Failed to get chain id. Err %v\n",err)
		}

		tx, err =types.SignTx(tx, types.NewEIP155Signer(chainId),privakeKey)
		if err != nil{
			log.Fatalf("Transaction failed. Err: %v\n",err)
		}
}