package utils

import (
	"encoding/json"
	"log"
	"math"
	"math/big"
)

type Error struct{
	message string
	err error
}

//This function converts wei to eth
func WeiToEth(w *big.Int) *big.Float {

	fp := big.NewFloat(math.Pow10(18))
	fw := new(big.Float).SetInt(w)

	v := new(big.Float).Quo(fw, fp)

	return v
}

//This function converts eth to wei
func EthToWei(eth *big.Float) *big.Int {

	fp := big.NewFloat(math.Pow10(18))
	ew := eth.Mul(eth, fp)

	res := new(big.Int)
	ew.Int(res)
	return res
}

//This function takes error and message as input and prints it on the terminal
func HandleError(err error, message string){
	if err != nil {
		newErr := Error{
			message:message,
			err:err,
		}
		bytes, _ := json.MarshalIndent(newErr," ","")
		log.Fatalln(string(bytes))
	}	

}