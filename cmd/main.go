package main

import (
	"fmt"

	"github.com/KarrenAeris/wallet/pkg/wallet"
	
)


func main() {
	svc := &wallet.Service{}
	account, err := svc.RegisterAccount("+992930000001")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(account.Balance)
}