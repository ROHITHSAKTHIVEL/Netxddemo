package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	AccountID        int                `json:"account_id" bson:"account_id"`
	TransactionCount int                `json:"transaction_count" bson:"transaction_count"`
	BucketStartDate  time.Time          `json:"bucket_start_date" bson:"bucket_start_date"`
	BucketEndDate    time.Time          `json:"bucket_end_date" bson:"bucket_end_date"`
	Transactions     []TransactionItem  `json:"transactions" bson:"transactions"`
}

type TransactionItem struct {
	Date            time.Time `json:"date" bson:"date"`
	Amount          int       `json:"amount" bson:"amount"`
	TransactionCode string    `json:"transaction_code" bson:"transaction_code"`
	Symbol          string    `json:"symbol" bson:"sumbol"`
	Price           string    `json:"price" bson:"price"`
	Total           string    `json:"total" bson:"total"`
}
