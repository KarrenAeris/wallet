package main

import (
	"github.com/KarrenAeris/wallet/pkg/wallet"
)


func main() {
	svc := &wallet.Service{}
	svc.ImportFromFile("data/import.txt")
}