package entities

import "time"

//OperationsType enum for 4 operations
type OperationsType uint

const (
	//NormalPurchase ...
	NormalPurchase OperationsType = iota + 1

	//PurchaseWithInstallments ...
	PurchaseWithInstallments

	//Withdrawal ...
	Withdrawal

	//CreditVoucher ...
	CreditVoucher
)

//String ...
func (o OperationsType) String() string {
	return [...]string{"None", "Normal Purchase", "Purchase With Installments", "Withdrawal", "Credit Voucher"}[o]
}

// GetOperationsType ...
func GetOperationsType(operationsType string) OperationsType {
	return map[string]OperationsType{
		"Normal Purchase":            NormalPurchase,
		"Purchase With Installments": PurchaseWithInstallments,
		"Withdrawal":                 Withdrawal,
		"Credit Voucher":             CreditVoucher,
	}[operationsType]
}

//Account entity
type Account struct {
	AccountID      uint
	DocumentNumber uint
}

//Transaction entity
type Transaction struct {
	TransactionID uint
	AccountID     uint
	OperationType OperationsType
	Amount        float64
	EventDate     time.Time
}
