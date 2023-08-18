package main

import (
	"context"
	"db_api/config"
	"db_api/constants"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server      *gin.Engine
	mongoclient *mongo.Client
	ctx         context.Context
)

//	func DisconnectDataBase() {
//		mongoclient.Disconnect(ctx)
//		fmt.Println("Disconnecting the database....")
//	}
func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)

	if err != nil {
		panic(err)
	}
	log.Fatal(server.Run(constants.Port))
}
