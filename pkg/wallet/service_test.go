package wallet

import (
	"fmt"
	"testing"
)

func TestService_FindAccountByID_success(t *testing.T) {
	svc := Service{}
	account, err := svc.RegisterAccount("+9920000001")
	if err != nil {
		fmt.Println(account)
	}

	account, err = svc.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}

func TestService_FindAccountByID_notFound(t *testing.T) {
	svc := Service{}
	account, err := svc.RegisterAccount("+9920000001")
	if err != nil {
		fmt.Println(account)
	}

	account, err = svc.FindAccountByID(3)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}