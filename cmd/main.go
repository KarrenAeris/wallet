package main

import (
	"github.com/KarrenAeris/wallet/pkg/wallet"
)


func main() {
	svc := &wallet.Service{}
	
	svc.ExportToFile("../data/export.txt")
	svc.ImportFromFile("../data/import.txt")
}