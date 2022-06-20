package utils

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

//Returs the chain id
func GetNetworkId(client *ethclient.Client,ctx context.Context) *big.Int{
	chainId, err := client.ChainID(ctx)
	if err != nil{
		log.Fatalf("\nFailed to get network id. Err %v\n", err)
	}
	return chainId
}

//Returns nonce of an address in a pending state 
func GetPendingNonce(client *ethclient.Client,ctx context.Context, pubAddr common.Address) uint64{
	nonce, err := client.PendingNonceAt(ctx, pubAddr)
	if err != nil{
		log.Fatalf("\n Failed to get pending nonce.Error: %v", err)
	}
	return nonce

}

//Suggests the optimum gas price for transaction execution
func SuggestGasPrice(client *ethclient.Client, ctx context.Context) *big.Int{
	gasPrice, err := client.SuggestGasPrice(ctx);
	if err != nil {
		log.Fatalf("\nFailed to suggest gas price.Error %v", err)
	}
	return gasPrice
}

//Converts private key from *ecdsa.private to string
func PrivateKeyToString(p *ecdsa.PrivateKey) string{
	privateKeyBytes := crypto.FromECDSA(p)
	return hexutil.Encode(privateKeyBytes)
}

//Converts public key from ecdsa.PublicKey to string
func PublicKeyToString(p ecdsa.PublicKey) string{
	publicKeyBytes := crypto.FromECDSAPub(&p)
	return hexutil.Encode(publicKeyBytes)
}

//Converts public address to string
func PublicAddressToString(p common.Address) string{
	return p.Hex()
}

//String to common.Address
func StringToAddress(address string) common.Address{
	return common.HexToAddress(address)
}