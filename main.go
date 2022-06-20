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
	
	balance1 := core.GetBalance("c50cd042edf5cf1d31cd5f9f04b499acd7e33a67")
	balance2 := core.GetBalance("00deb43c56cca15ead73b21847115dc769df767c")
	log.Println("Balance1: ", balance1)
	log.Println("Balance2: ", balance2)
	
}
