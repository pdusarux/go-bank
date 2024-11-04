package service

import (
	"fmt"
	"go-bank/errs"
	"go-bank/logs"
	"go-bank/repository"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

type accountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return accountService{accountRepo: accountRepo}
}

func (s accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error) {
	// Validate the request
	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type must be either saving or checking")
	}

	if request.Amount < 5000 {
		return nil, errs.NewValidationError("amount must be greater than 5000")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format(time.DateTime),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}
	newAccount, err := s.accountRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return &AccountResponse{
		AccountID:   fmt.Sprintf("%d", newAccount.AccountID),
		OpeningDate: newAccount.OpeningDate,
		AccountType: newAccount.AccountType,
		Amount:      newAccount.Amount,
		Status:      strconv.Itoa(newAccount.Status),
	}, nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accountRepo.GetAll(customerID)
	logs.Info("accounts", zap.Any("accounts", accounts))
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	accountResponses := []AccountResponse{}
	for _, account := range accounts {
		accountResponses = append(accountResponses, AccountResponse{
			AccountID:   fmt.Sprintf("%d", account.AccountID),
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      strconv.Itoa(account.Status),
		})
	}
	return accountResponses, nil
}
