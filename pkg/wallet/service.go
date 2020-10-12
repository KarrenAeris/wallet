package wallet

import (
	"errors"
	"os"
	"log"
	"strconv"
	"github.com/KarrenAeris/wallet/pkg/types"
	"github.com/google/uuid"
)

// распространённые возможные ошибки
var ErrPhoneRegistered = errors.New("phone already registred")
var ErrAmountMustBePositive = errors.New("amount must be greater than zero")
var ErrAccountNotFound = errors.New("account not found")
var ErrNotEnoughtBalance = errors.New("account not enough balance")
var ErrPaymentNotFound = errors.New("payment not found")
var ErrFavoriteNotFound = errors.New("favorite not found")
var ErrFileNotFound = errors.New("file not found")


type Service struct  {
	NextAccountID int64
	accounts []*types.Account 
	payments []*types.Payment
	favorites []*types.Favorite 
}


// RegisterAccount регистрирует  нового пользователя в системе
func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return nil, ErrPhoneRegistered
		}	
	}
	s.NextAccountID++
	account := &types.Account {
		ID : s.NextAccountID,
		Phone: phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)

	return account, nil
}


// Deposit пополняет счёт пользователя
func (s *Service) Deposit(accountID int64, amount types.Money) error {
	if amount <= 0 {
		return ErrAmountMustBePositive
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID  == accountID {
			account = acc
			break
		}
	}

	if account == nil {
		return ErrAccountNotFound
	}

	account.Balance += amount
	return nil
}


// Pay платит определенную сумму денег за категорию
func (s *Service) Pay(accountID int64, amount types.Money, category types.PaymentCategory) (*types.Payment, error) {
	if amount <= 0 {
		return nil, ErrAmountMustBePositive
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID  == accountID {
			account = acc
			break
		}
	}
	if account == nil {
		return nil, ErrAccountNotFound
	}
	if account.Balance < amount {
		return nil, ErrNotEnoughtBalance
	}

	account.Balance -= amount
	paymentID := uuid.New().String()
	payment := &types.Payment {
		ID: paymentID,
		AccountID: accountID,
		Amount: amount,
		Category: category,
		Status: types.PaymentStatusInProgress,
	}
	s.payments = append(s.payments, payment)
	return payment, nil
}


// FindAccountById ищет пользователя по ID
func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}

	if account == nil {
		return nil, ErrAccountNotFound
	}

	return account, nil
}


// FindPaymentByID ищет платёж по ID
func (s *Service) FindPaymentByID(paymentID string) (*types.Payment, error) {
	var payment *types.Payment

	for _, pay := range s.payments {
		if pay.ID == paymentID {
			payment = pay
		}
	}

	if payment == nil {
		return nil, ErrPaymentNotFound
	}

	return payment, nil
}


// Reject отменяет платёж
func (s *Service) Reject(paymentID string) error {
	pay, err := s.FindPaymentByID(paymentID)
	if err != nil {
		return ErrPaymentNotFound
	}

	acc, err := s.FindAccountByID(pay.AccountID)
	if err != nil {
		return ErrAccountNotFound
	}

	pay.Status = types.PaymentStatusFail
	acc.Balance += pay.Amount

	return nil
}


// Repeat повторяет платёж по идентификатору 
func (s *Service) Repeat(paymentID string) (*types.Payment, error) {
	pay, err := s.FindPaymentByID(paymentID)
	if err != nil {
	  return nil, err
	}
  
	payment, err := s.Pay(pay.AccountID, pay.Amount, pay.Category)
	if err != nil {
	  return nil, err
	}
  
	return payment, nil
}

// FindFavoriteByID ищет платёж по ID в Избранное
func (s *Service) FindFavoriteByID(favoriteID string) (*types.Favorite, error) {
	for _, favorite := range s.favorites {
		if favorite.ID == favoriteID {
			return favorite, nil
		}
	}
	return nil, ErrFavoriteNotFound
}

// FavoritePayment добавляет новый платеж в Избранное
func (s *Service) FavoritePayment(paymentID string, name string) (*types.Favorite, error) {
	payment, err := s.FindPaymentByID(paymentID)

	if err != nil {
		return nil, err
	}

	favoriteID := uuid.New().String()
	newFavorite := &types.Favorite{
		ID:        favoriteID,
		AccountID: payment.AccountID,
		Name:      name,
		Amount:    payment.Amount,
		Category:  payment.Category,
	}

	s.favorites = append(s.favorites, newFavorite)
	return newFavorite, nil
}

// PayFromFavorite позволяет совершить платеж из Избранное
func (s *Service) PayFromFavorite(favoriteID string) (*types.Payment, error) {
	favorite, err := s.FindFavoriteByID(favoriteID)
	if err != nil {
		return nil, err
	}

	payment, err := s.Pay(favorite.AccountID, favorite.Amount, favorite.Category)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

// ExportToFile экспортирует все аккаунты в файл, путь к которому указан в переменной path
func (s *Service) ExportToFile(path string) error {
	file, err := os.Create(path)

	if err != nil {
		log.Print(err)
		return ErrFileNotFound
	}

	defer func() {
		err := file.Close(); 
		if err != nil {
			log.Print(err)
		}
	}()

	stringAccount := ""

	for _, acc := range s.accounts {
		ID := strconv.FormatInt(int64(acc.ID), 10)
		phone := string(acc.Phone)
		balance := strconv.FormatInt(int64(acc.Balance), 10)

		stringAccount += ID + ";"
		stringAccount += phone + ";"
		stringAccount += balance + "|"
	}

	_, err = file.Write([]byte(stringAccount))
	if err != nil {
		log.Print(err)
		return ErrFileNotFound
	}

	return nil
}