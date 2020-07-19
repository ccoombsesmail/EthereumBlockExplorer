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
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	// "EthereumBlockExplorer/config"
)



func main() {

	// Connect to mongodb 
	var mongoURI string
	if (os.Getenv("MONGODB_URI") == "") {
		// mongoURI = "mongodb+srv://user1:" + config.GetKey() + "@cluster0.4mnma.mongodb.net/<blockHistoryDB>?retryWrites=true&w=majority"
		mongoURI = ""
	} else {
		mongoURI = os.Getenv("MONGODB_URI")
	}
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
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

	// Setup websocket connection to remote infura node and then sub to receive latest blocks
	ethClient, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/7fee20f62b264b7ab41d4fb69dde3c76")
	if err != nil {
		log.Fatal(err)
	}
	go websockets.SubToBlockHeader(blocksCollection, transactionsCollection, ethClient)


	// Setup http server and routes
  port := os.Getenv("PORT")
 	if port == "" {
			port = "5000"
			log.Printf("Defaulting to port %s", port)
    }
	http.HandleFunc("/ws", websockets.ServeWs)
	routes.SetupBlockRoutes(mongoClient)
	routes.SetupTransactionRoutes(mongoClient, ethClient)

	fs := http.FileServer(http.Dir("./build"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	
}




