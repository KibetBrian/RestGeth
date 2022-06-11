package main

import (
	"log"

	"github.com/KibetBrian/geth/core"
	"github.com/KibetBrian/geth/utils"
)
const (
	address = "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B"
)

func main(){
	utils.ConfigureEnv()
	balance := core.GetBalance(address)
	log.Println(balance)
}