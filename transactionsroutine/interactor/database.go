package interactor

import (
	"github.com/yugendra/TransactionsRoutine/entities"
)

/*DatabaseInteractor is the contract between main logic and database layer
 *Any database can be used in transactionsroutine only if it is implementing this interface
 */
type DatabaseInteractor interface {
	Close() error

	CreateAccount(account *entities.Account) error
	GetAccount(accountID uint) (*entities.Account, error)
	CreateTransaction(transaction *entities.Transaction) error
}
