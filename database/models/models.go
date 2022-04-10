package models

import (
	"time"
)

//Account model for accounts db table
type Account struct {
	AccountID      uint `gorm:"column:account_id;primaryKey;autoIncrement"`
	DocumentNumber uint `gorm:"column:document_number;type:int;unique;not null;"`
}

//Transaction model for transactions db table
type Transaction struct {
	TransactionID   uint      `gorm:"column:transaction_id;primaryKey;autoIncrement"`
	AccountID       uint      `gorm:"column:account_id;type:int;not null"`
	OperationTypeID uint      `gorm:"column:operation_type_id;type:int;not null"`
	Amount          float64   `gorm:"column:amount;type:int;not null"`
	EventDate       time.Time `gorm:"column:event_date;type:TIMESTAMP;DEFAULT:NOW()"`
	Account         Account   `gorm:"references:AccountID"`
}
