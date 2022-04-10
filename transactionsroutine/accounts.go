package transactionsroutine

import (
	"errors"
	"fmt"
	"github.com/yugendra/TransactionsRoutine/entities"
	"github.com/yugendra/TransactionsRoutine/transactionsroutine/interactor"
	"log"
)

/*CreateAccount Core logic to create customer account in DB.
 * For now this function is very simple as data set is very small and
 * there is no validations or filtering is required at this stage
 */
func CreateAccount(db interactor.DatabaseInteractor, account *entities.Account) error {
	//No validations required here as document id got validated in handler
	//And in case of duplicate document id DB will throw an error

	err := db.CreateAccount(account)
	if err != nil {
		errMsg := fmt.Sprintf("failed to create account for document number: %d.", account.DocumentNumber)
		log.Printf("%s Error: %s", errMsg, err.Error())
		return errors.New(errMsg)
	}

	return nil
}

/*GetAccount Core logic to fetch customer account in DB.
 * For now this function is very simple as data set is very small and
 * there is no validations or filtering is required at this stage
 */
func GetAccount(db interactor.DatabaseInteractor, accountID int) (*entities.Account, error) {
	//No validations required here as account id got validated in handler

	account, err := db.GetAccount(uint(accountID))
	if err != nil {
		errMsg := fmt.Sprintf("failed to get account : %d.", accountID)
		log.Printf("%s Error: %s", errMsg, err.Error())
		return nil, errors.New(errMsg)
	}

	return account, nil
}
