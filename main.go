package main

import (
	"log"

	"github.com/KibetBrian/geth/core"
	"github.com/KibetBrian/geth/utils"
)

const (
	address = "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B"
)

func main() {

	utils.ConfigureEnv()
	wallet1 := core.DecryptKeystore("password1", 1)
	wallet2 := core.DecryptKeystore("password2", 2)
	address1 := utils.PublicAddressToString(wallet1.PublicAddress)
	address2 := utils.PublicAddressToString(wallet2.PublicAddress)
	core.TranferEther(address2, 0.5, "password1")
	balance1 := core.GetBalance(address1)
	balance2 := core.GetBalance(address2)

	log.Println("Address1: ", address1)
	log.Println("Address2: ", address2)
	log.Println("Balance1: ", balance1)
	log.Println("Balance2: ", balance2)
}
