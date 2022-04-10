package handlers

import "time"

//TODO: API requests and responses modules are manually written. Use swagger code generated to generate the modules.

//Account request and response module for /account get and post API.
type Account struct {
	AccountID      uint `json:"account_id"`
	DocumentNumber uint `json:"document_number" validate:"required,numeric,gte=1"`
}

//Transaction request and response module for /transaction post API.
type Transaction struct {
	TransactionID uint      `json:"transaction_id"`
	AccountID     uint      `json:"account_id" validate:"required,numeric,gte=1"`
	OperationType string    `json:"operation_type" validate:"required,eq=Normal Purchase|eq=Purchase With Installments|eq=Withdrawal|eq=Credit Voucher"`
	Amount        float64   `json:"amount" validate:"required,numeric"`
	EventDate     time.Time `json:"event_date"`
}
