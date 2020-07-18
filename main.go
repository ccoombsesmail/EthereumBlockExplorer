package main

import (
	"context"
	"log"
	"time"
	"net/http"
	"EthereumBlockExplorer/routes"
	"EthereumBlockExplorer/websockets"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"EthereumBlockExplorer/config"
)



func main() {

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://heroku_pfznv2c0:m2c0tf7v719asg1c0andhkha85@ds035177.mlab.com:35177/heroku_pfznv2c0"))
	// mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://user1:" + config.GetKey() + "@cluster0.4mnma.mongodb.net/<blockHistoryDB>?retryWrites=true&w=majority"))
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

  port := os.Getenv("PORT")
 	if port == "" {
			port = "5000"
			log.Printf("Defaulting to port %s", port)
    }
	// http.HandleFunc("/ws", websockets.ServeWs)
	routes.SetupBlockRoutes(mongoClient)
	routes.SetupTransactionRoutes(mongoClient)
	go websockets.SubToBlockHeader(blocksCollection, transactionsCollection)

	fs := http.FileServer(http.Dir("./build"))
	http.Handle("/", fs)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(config.GetKey()))
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
	
}





