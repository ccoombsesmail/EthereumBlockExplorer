package main

import (
	"context"
	"log"
	"time"
	"net/http"
	"BlockExplorer/backend/routes"
	"BlockExplorer/backend/config"
	"BlockExplorer/backend/websockets"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func main() {

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://user1:" + config.GetKey() + "@cluster0.4mnma.mongodb.net/<blockHistoryDB>?retryWrites=true&w=majority"))
	if err != nil {
			log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongoClient.Connect(ctx)
	if err != nil {
			log.Fatal(err)
	}
	defer mongoClient.Disconnect(ctx)

	db := mongoClient.Database("blockHistoryDB")
	blocksCollection := db.Collection("blocks")
	transactionsCollection := db.Collection("transactions")


	http.HandleFunc("/ws", websockets.ServeWs)
	routes.SetupBlockRoutes(mongoClient)
	routes.SetupTransactionRoutes(mongoClient)
	go websockets.SubToBlockHeader(blocksCollection, transactionsCollection)
	log.Fatal(http.ListenAndServe(":5000", nil))

}





