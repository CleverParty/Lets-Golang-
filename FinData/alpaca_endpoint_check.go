package main

import (
	"fmt"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
)

func initAlpacaCreds() {
	var valueToBeAsserted = "test"
	fmt.Println("Value is :", valueToBeAsserted)
	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
}

func init() {
	os.Setenv(common.EnvApiKeyID, "")
	os.Setenv(common.EnvApiSecretKey, "")

	fmt.Printf("Running w/ credentials [%v %v]\n", common.Credentials().ID, common.Credentials().Secret)

	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
}

func main() {
	alpacaClient := alpaca.NewClient(common.Credentials())
	acct, err := alpacaClient.GetAccount()
	if err != nil {
		panic(err)
	}

	fmt.Println(*acct)
}

// windows file modification:
// warning: LF will be replaced by CRLF in test.go.
// The file will have its original line endings in your working directory
// warning: LF will be replaced by CRLF in FinData/alpaca_endpoint_check.go.
// The file will have its original line endings in your working directory
