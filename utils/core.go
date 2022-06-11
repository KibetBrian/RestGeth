package utils

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

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