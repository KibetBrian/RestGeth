package core

import (
	"crypto/ecdsa"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)
const (
	path = "keydir"
)

type RawWallet struct{
	PrivateKey *ecdsa.PrivateKey
	PublicKey ecdsa.PublicKey
	PublicAddress common.Address
}

//Creates new keystore and stores it in the keystores directory
func GenerateKeyStore(passphrase string) {
	keystore := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := keystore.NewAccount(passphrase)
	if err != nil{
		log.Printf("Failed to generate new key from keystore. Err: %v\n", err)
	}
	log.Println("Generated keystore with addresss:  ",account.Address.Hex())	
}

//This function decrypts the file to get the addresses
func DecryptKeystore(passphrase string, keyNum int) *RawWallet{

	fileInfo, err := ioutil.ReadDir(path)
	if err != nil{
		log.Printf("Failed to read the directory. Err: %v",err)
	}
	fileName := fileInfo[keyNum-1].Name()
	filePath := filepath.Join(path,fileName)
	
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil{
		log.Printf("Failed to read file. Err: %v",err)
	}

	key, err := keystore.DecryptKey(fileBytes, passphrase)
	if err != nil{
		log.Printf("Failed to read file. Err: %v",err)
	}
	privateKey := key.PrivateKey
	publicKey := privateKey.PublicKey
	publicAddress := crypto.PubkeyToAddress(publicKey)

	newRawWallet := RawWallet{
		PrivateKey: privateKey,
		PublicKey: publicKey,
		PublicAddress: publicAddress,
	}
	log.Println("The public address of the wallet is: ", publicAddress.Hex())
	return &newRawWallet
}

