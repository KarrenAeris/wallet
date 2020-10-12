package main

import (
	"github.com/KarrenAeris/wallet/pkg/wallet"
)


func main() {
	svc := &wallet.Service{}
	svc.RegisterAccount("+992930000000")
	svc.RegisterAccount("+992931111111")
	svc.RegisterAccount("+992932222222")
	svc.ExportToFile("data/export.txt")
}