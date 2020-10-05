package wallet

import (
	"fmt"
	"testing"
)

// Автотесты для FindAccountByID
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

// Автотесты для Reject
func TestService_Reject_success(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("+9920000001")

	account, err := svc.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = svc.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	payment, err := svc.Pay(account.ID, 100_00, "food")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	pay, err := svc.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = svc.Reject(pay.ID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}

func TestService_Reject_fail(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("+9920000001")

	account, err := svc.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = svc.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	payment, err := svc.Pay(account.ID, 100_00, "food")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	pay, err := svc.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	wrongPayID := pay.ID + "14"
	err = svc.Reject(wrongPayID)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}

// Автотесты для Repeat
func TestService_Repeat_success(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("+9920000001")
  
	account, err := svc.FindAccountByID(1)
	if err != nil {
	  t.Errorf("\ngot > %v \nwant > nil", err)
	}
  
	err = svc.Deposit(account.ID, 1000_00)
	if err != nil {
	  t.Errorf("\ngot > %v \nwant > nil", err)
	}
  
	payment, err := svc.Pay(account.ID, 100_00, "auto")
	if err != nil {
	  t.Errorf("\ngot > %v \nwant > nil", err)
	}
  
	pay, err := svc.FindPaymentByID(payment.ID)
	if err != nil {
	  t.Errorf("\ngot > %v \nwant > nil", err)
	}
  
	pay, err = svc.Repeat(pay.ID)
	if err != nil {
	  t.Errorf("Repeat(): Error(): can't pay for an account(%v): %v", pay.ID, err)
	}
}

// Автотесты для PayFromFavorite
func TestService_Favorite_success_user(t *testing.T) {
	svc := Service{}

	account, err := svc.RegisterAccount("+992000000001")
	if err != nil {
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	err = svc.Deposit(account.ID, 100_00)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	payment, err := svc.Pay(account.ID, 10_00, "food")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	favorite, err := svc.FavoritePayment(payment.ID, "Tcell")
	if err != nil {
		t.Errorf("FavoritePayment(): Error(): can't find the favorite(%v): %v", favorite, err)
	}

	paymentFavorite, err := svc.PayFromFavorite(favorite.ID)
	if err != nil {
		t.Errorf("PayFromFavorite(): Error(): can't pay from the favorite(%v): %v", paymentFavorite, err)
	}
}