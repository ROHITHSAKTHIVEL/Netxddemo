package transactions

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"transaction.go/config"
	"transaction.go/models"
)

func TransactionContext() *mongo.Collection {
	client, err := config.ConnectDataBase()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return config.GetCollection(client, "sample_analytics", "transactions")
}

func FindTransaction() ([]*models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"transaction_count": bson.D{{Key: "$gt", Value: 85}, {Key: "$lt", Value: 90}}} // Add your filter condition here

	limit := options.Find().SetSort(bson.D{{Key: "transaction_count", Value: 1}}).SetLimit(10)
	result, err := TransactionContext().Find(ctx, filter, limit)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		defer result.Close(ctx) // Close the result cursor when done

		var val []*models.Transaction
		for result.Next(ctx) {
			val2 := &models.Transaction{}
			err := result.Decode(val2)
			if err != nil {
				return nil, err
			}
			val = append(val, val2)
		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		if len(val) == 0 {
			return []*models.Transaction{}, nil
		}
		return val, nil
	}
}

//{"transaction_count",bson.D{{"$gt":85}}}

func FetchAggregate() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	matchstage := bson.D{{Key: "$match", Value: bson.D{{Key: "transaction_count", Value: 100}}}}

	groupsatge := bson.D{
		{
			Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$account_id"},
				{Key: "total_count", Value: bson.D{{Key: "$sum", Value: "$transaction_count"}}},
			}}}
	result, err := TransactionContext().Aggregate(ctx, mongo.Pipeline{matchstage, groupsatge})
	if err != nil {
		fmt.Println(err.Error())

	} else {

		var showswithinfo []bson.M
		if err = result.All(ctx, &showswithinfo); err != nil {
			panic(err)
		}
		formatted_data, err := json.MarshalIndent(showswithinfo, "", " ")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(formatted_data))
		}
	}
}

func Update(initialValue int, newValue int) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.D{{Key: "account_id", Value: initialValue}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "account_id", Value: newValue}}}}
	result, err := TransactionContext().UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return result, nil

}
