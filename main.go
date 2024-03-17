package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/2ff306b1c7834dc89508b4ab646851d5"
var ganacheURL = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)

	if err != nil {
		log.Fatal("dial context", err)
	}

	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)

	if err != nil {
		log.Fatal("block error", err)
	}

	fmt.Println("block:", block.Number())

	addr := "0xb151bbd2Fd06776E6394410C56579705E0D7498B"

	addres := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), addres, nil)
	if err != nil {
		log.Fatal("get balance error", err)
	}
	fmt.Println("balance", balance)

	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	balanceEth := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(balanceEth)

	request, _ := http.Get("https://api.coinbase.com/v2/exchange-rates?currency=ETH")

	var data map[string]interface{}
	body, _ := io.ReadAll(request.Body)
	er := json.Unmarshal(body, &data)

	if er != nil {
		log.Fatal("get balance error", er)
	}

	fmt.Println("1ETH USD:", data["data"].(map[string]interface{})["rates"].(map[string]interface{})["USD"])
	transactionValuedUSD, _, _ := new(big.Float).Parse(data["data"].(map[string]interface{})["rates"].(map[string]interface{})["USD"].(string), 10)
	transactionValuedUSD.Mul(transactionValuedUSD, balanceEth)
	fmt.Println("transaction value to USD", transactionValuedUSD.String())

	fmt.Println("Taxa para BRL:", data["data"].(map[string]interface{})["rates"].(map[string]interface{})["BRL"])
	transactionValuedBRL, _, _ := new(big.Float).Parse(data["data"].(map[string]interface{})["rates"].(map[string]interface{})["BRL"].(string), 10)
	transactionValuedBRL.Mul(transactionValuedBRL, balanceEth)
	fmt.Println("transaction value to BRL", transactionValuedBRL.String())

	fmt.Println("Taxa para BTC:", data["data"].(map[string]interface{})["rates"].(map[string]interface{})["BTC"])
	transactionValuedBTC, _, _ := new(big.Float).Parse(data["data"].(map[string]interface{})["rates"].(map[string]interface{})["BTC"].(string), 10)
	transactionValuedBTC.Mul(transactionValuedBTC, balanceEth)
	fmt.Println("transaction value to BTC", transactionValuedBTC.String())
}
