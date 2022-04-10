package transactionsroutine

import (
	"errors"
	"fmt"
	"github.com/yugendra/TransactionsRoutine/entities"
)

/*validateTransaction validates the amount for transaction type.
 * amount should be negative for types: Normal Purchase, Purchase With Installments and Withdrawal
 * amount should be positive for type: Credit Voucher
 */
func validateTransaction(transaction *entities.Transaction) error {
	if (transaction.OperationType == entities.NormalPurchase ||
		transaction.OperationType == entities.PurchaseWithInstallments ||
		transaction.OperationType == entities.Withdrawal) && transaction.Amount >= 0 {
		errMsg := fmt.Sprintf("Invalid amount %f for operation type %s",
			transaction.Amount, transaction.OperationType.String())
		return errors.New(errMsg)
	}

	if transaction.OperationType == entities.CreditVoucher && transaction.Amount <= 0 {
		errMsg := fmt.Sprintf("Invalid amount %f for operation type %s",
			transaction.Amount, transaction.OperationType.String())
		return errors.New(errMsg)
	}

	return nil
}
