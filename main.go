package main

import (
	"context"
	"log"
	"time"
	"net/http"
	"EthereumBlockExplorer/routes"
	"EthereumBlockExplorer/config"
	"EthereumBlockExplorer/websockets"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
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

  port := os.Getenv("PORT")
 	if port == "" {
                port = "5000"
                log.Printf("Defaulting to port %s", port)
        }
	http.HandleFunc("/ws", websockets.ServeWs)
	routes.SetupBlockRoutes(mongoClient)
	routes.SetupTransactionRoutes(mongoClient)
	go websockets.SubToBlockHeader(blocksCollection, transactionsCollection)
	fs := http.FileServer(http.Dir("./build"))
	http.Handle("/", fs)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
	
}





