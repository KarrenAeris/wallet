package wallet

import (
	"fmt"
	"testing"
)

func TestService_FindAccountById_success(t *testing.T) {
	svc := Service{}
	account, err := svc.RegisterAccount("+9920000001")
	if err != nil {
		fmt.Println(account)
	}

	account, err = svc.FindAccountById(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}

func TestService_FindAccountById_notFound(t *testing.T) {
	svc := Service{}
	account, err := svc.RegisterAccount("+9920000001")
	if err != nil {
		fmt.Println(account)
	}

	account, err = svc.FindAccountById(3)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}