package test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yugendra/TransactionsRoutine/entities"
	"github.com/yugendra/TransactionsRoutine/test/mocks"
	"github.com/yugendra/TransactionsRoutine/transactionsroutine"
	"testing"
	"time"
)

var (
	testTime, _ = time.Parse("2022-04-09T19:39:06.834491Z", "2022-04-09T19:39:06.834491Z")
	errTest     = errors.New("failed to create entity in database")

	transaction = &entities.Transaction{
		AccountID:     1,
		OperationType: entities.NormalPurchase,
		Amount:        -10.75,
	}
	expectTransaction = &entities.Transaction{
		TransactionID: 1,
		AccountID:     1,
		OperationType: entities.NormalPurchase,
		Amount:        -10.75,
		EventDate:     testTime,
	}
	account = &entities.Account{
		DocumentNumber: 1234567890,
	}
	expectAccount = &entities.Account{
		AccountID:      1,
		DocumentNumber: 1234567890,
	}
)

//TestCreateTransaction ...
func TestCreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	databaseMock := mocks.NewMockDatabaseInteractor(ctrl)

	makeDatabase := func(transaction *entities.Transaction, expectError bool) *mocks.MockDatabaseInteractor {
		if !expectError {
			databaseMock.EXPECT().CreateTransaction(gomock.AssignableToTypeOf(&entities.Transaction{})).
				DoAndReturn(
					func(transaction *entities.Transaction) error {
						transaction.TransactionID = 1
						transaction.EventDate = testTime
						return nil
					}).Times(1)
		} else {
			databaseMock.EXPECT().CreateTransaction(gomock.AssignableToTypeOf(&entities.Transaction{})).
				Return(errTest).Times(1)
		}

		return databaseMock
	}

	tests := []struct {
		name              string
		db                *mocks.MockDatabaseInteractor
		transaction       *entities.Transaction
		expectTransaction *entities.Transaction
		expectError       bool
	}{
		{
			name: "Normal purchase negative amount",
			db: makeDatabase(&entities.Transaction{
				AccountID:     1,
				OperationType: entities.NormalPurchase,
				Amount:        -10.75,
			}, false),
			transaction:       transaction,
			expectTransaction: expectTransaction,
			expectError:       false,
		},
		{
			name: "Normal purchase positive amount",
			db: makeDatabase(&entities.Transaction{
				AccountID:     1,
				OperationType: entities.NormalPurchase,
				Amount:        10.75,
			}, true),
			transaction:       transaction,
			expectTransaction: transaction,
			expectError:       true,
		},
		{
			name: "Purchase with installments negative amount",
			db: makeDatabase(&entities.Transaction{
				AccountID:     1,
				OperationType: entities.PurchaseWithInstallments,
				Amount:        -10.75,
			}, false),
			transaction:       transaction,
			expectTransaction: expectTransaction,
			expectError:       false,
		},
		{
			name:              "Purchase with installments positive amount",
			db:                makeDatabase(transaction, true),
			transaction:       transaction,
			expectTransaction: transaction,
			expectError:       true,
		},
		{
			name:              "Withdrawal negative amount",
			db:                makeDatabase(transaction, false),
			transaction:       transaction,
			expectTransaction: expectTransaction,
			expectError:       false,
		},
		{
			name:              "Withdrawal positive amount",
			db:                makeDatabase(transaction, true),
			transaction:       transaction,
			expectTransaction: transaction,
			expectError:       true,
		},
		{
			name:              "Credit voucher negative amount",
			db:                makeDatabase(transaction, true),
			transaction:       transaction,
			expectTransaction: transaction,
			expectError:       true,
		},
		{
			name:              "Credit voucher positive amount",
			db:                makeDatabase(transaction, false),
			transaction:       transaction,
			expectTransaction: expectTransaction,
			expectError:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := transactionsroutine.CreateTransaction(tt.db, tt.transaction)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectTransaction, tt.transaction)
			}
		})
	}
}

//TestCreateAccount ...
func TestCreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	databaseMock := mocks.NewMockDatabaseInteractor(ctrl)

	makeDatabase := func(account *entities.Account, expectError bool) *mocks.MockDatabaseInteractor {
		if !expectError {
			databaseMock.EXPECT().CreateAccount(gomock.AssignableToTypeOf(&entities.Account{})).
				DoAndReturn(
					func(account *entities.Account) error {
						account.AccountID = 1
						return nil
					}).Times(1)
		} else {
			databaseMock.EXPECT().CreateAccount(gomock.AssignableToTypeOf(&entities.Account{})).
				Return(errTest).Times(1)
		}
		return databaseMock
	}

	tests := []struct {
		name          string
		db            *mocks.MockDatabaseInteractor
		account       *entities.Account
		expectAccount *entities.Account
		expectError   bool
	}{
		{
			name:          "success",
			db:            makeDatabase(account, false),
			account:       account,
			expectAccount: expectAccount,
			expectError:   false,
		},
		{
			name:        "return error",
			db:          makeDatabase(account, true),
			account:     account,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := transactionsroutine.CreateAccount(tt.db, tt.account)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectAccount, tt.account)
			}
		})
	}
}

func TestGetAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	databaseMock := mocks.NewMockDatabaseInteractor(ctrl)

	makeDatabase := func(accountId uint, expectError bool) *mocks.MockDatabaseInteractor {
		if !expectError {
			databaseMock.EXPECT().GetAccount(gomock.AssignableToTypeOf(uint(0))).
				Return(&entities.Account{AccountID: accountId, DocumentNumber: 1234567890}, nil).Times(1)
		} else {
			databaseMock.EXPECT().GetAccount(gomock.AssignableToTypeOf(uint(0))).
				Return(nil, errTest).Times(1)
		}
		return databaseMock
	}

	tests := []struct {
		name          string
		db            *mocks.MockDatabaseInteractor
		accountID     int
		expectAccount *entities.Account
		expectError   bool
	}{
		{
			name:          "success",
			db:            makeDatabase(1, false),
			accountID:     1,
			expectAccount: expectAccount,
			expectError:   false,
		},
		{
			name:        "return error",
			db:          makeDatabase(1, true),
			accountID:   1,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account, err := transactionsroutine.GetAccount(tt.db, tt.accountID)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, account)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectAccount, account)
			}
		})
	}
}
