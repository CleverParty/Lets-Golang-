package main

import (
	"fmt"
	"testing"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

func init() {
	valueToBePrinted = "test"
	fmt.Println("Value is :", *valueToBePrinted)
	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
}

func TestRun(t *testing.T) {
	expected := "Setup Travis CI for Golang project"
	got := run()
	if expected != got {
		t.Fatalf("expected %v got %v", expected, got)
	}
}
