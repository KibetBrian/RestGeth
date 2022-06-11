package core

import (
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	PrivateKey    string
	PublicKey     string
	PublicAddress string
}

func CreateWallet() *Wallet {
	//Create private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Printf("Failed to generate private key. Err: %v", privateKey)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHexAddressString := hexutil.Encode(privateKeyBytes)

	//Generates public key from private key
	publicKey := privateKey.PublicKey
	publicKeyBytes := crypto.FromECDSAPub(&publicKey)
	publicKeyAddressString := hexutil.Encode(publicKeyBytes)

	//Generates public adress from public key
	publicAddress := crypto.PubkeyToAddress(publicKey)
	publicAddressString := publicAddress.Hex()

	//Create Wallet struct
	newWallet := &Wallet{
		PrivateKey:    privateKeyHexAddressString,
		PublicKey:     publicKeyAddressString,
		PublicAddress: publicAddressString,
	}
	return newWallet
}
