package transactionsroutine

import (
	"errors"
	"fmt"
	"github.com/yugendra/TransactionsRoutine/entities"
	"github.com/yugendra/TransactionsRoutine/transactionsroutine/interactor"
	"log"
)

/*CreateTransaction Core logic to create transaction for a customer in DB.
 * For now this function is very simple.
 * Only validation the transaction type and amount.
 * There is no other validation of filtering is required at this stage.
 */
func CreateTransaction(db interactor.DatabaseInteractor, transaction *entities.Transaction) error {

	err := validateTransaction(transaction)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = db.CreateTransaction(transaction)
	if err != nil {
		errMsg := fmt.Sprintf("failed to create transaction: %+v", transaction)
		log.Printf("%s Error: %s", errMsg, err.Error())
		return errors.New(errMsg)
	}

	return nil
}
