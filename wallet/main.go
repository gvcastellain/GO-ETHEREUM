package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	pvk, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal(err)
	}

	privateData := crypto.FromECDSA(pvk)
	fmt.Println(hexutil.Encode(privateData))

	publicData := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(publicData))

	fmt.Println(crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}
