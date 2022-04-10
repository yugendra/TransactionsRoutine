package database

import (
	"github.com/yugendra/TransactionsRoutine/database/models"
	"github.com/yugendra/TransactionsRoutine/entities"
)

//CreateAccount creates the account in DB for document number and returns account id
func (db *database) CreateAccount(account *entities.Account) error {
	accountModel := &models.Account{DocumentNumber: account.DocumentNumber}
	err := db.conn.Create(accountModel).Error
	if err != nil {
		return err
	}

	account.AccountID = accountModel.AccountID
	return nil
}

//GetAccount fetches account info from DB for account id
func (db *database) GetAccount(accountID uint) (*entities.Account, error) {
	accountModel := &models.Account{}
	err := db.conn.Where("account_id = ?", accountID).First(accountModel).Error
	if err != nil {
		return nil, err
	}

	account := &entities.Account{
		AccountID:      accountModel.AccountID,
		DocumentNumber: accountModel.DocumentNumber,
	}

	return account, nil
}

//CreateTransaction creates transaction in DB for account id, transaction type and amount
func (db *database) CreateTransaction(transaction *entities.Transaction) error {
	transactionModel := models.Transaction{
		AccountID:       transaction.AccountID,
		OperationTypeID: uint(transaction.OperationType),
		Amount:          transaction.Amount,
	}

	err := db.conn.Create(&transactionModel).Error
	if err != nil {
		return err
	}

	transaction.TransactionID = transactionModel.TransactionID
	transaction.EventDate = transactionModel.EventDate
	return nil
}
