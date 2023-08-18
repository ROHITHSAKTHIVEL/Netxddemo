package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	transactions "transaction.go/services"
)

var (
	mongoClient *mongo.Client
)

func main() {
	fmt.Println("MongoDB server started...")
	// transactionss, _ := transactions.FindTransaction()
	// for _, transactions := range transactionss {
	// 	fmt.Print("TransactionCount = ", transactions.TransactionCount, transactions.ID, transactions.AccountID, transactions.BucketStartDate, transactions.BucketEndDate)
	// 	fmt.Println("")
	// }
	//transaction, _ :=
	//transactions.FetchAggregate()
	//for _, transactions := range transaction {
	//	fmt.Println(transactions)
	//}

	result, _ := transactions.Update(209363, 454545)
	fmt.Println(result.MatchedCount)
}
