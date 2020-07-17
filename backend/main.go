package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"net/http"
	"encoding/json"
	"BlockExplorer/backend/routes"
	"BlockExplorer/backend/typehelper"
	"BlockExplorer/backend/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/websocket"
)

var Client *ethclient.Client
// var RPCClient *rpc.Client

var upgrader = websocket.Upgrader{
  ReadBufferSize:  1024,
  WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool { return true },
}
var broadcast = make(chan typehelper.BlockData)


func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
	}
	reader(ws)
}


func reader(conn *websocket.Conn) {
  for {
		block := <-broadcast
		b, err := json.Marshal(block)
		if err != nil {
			fmt.Println(err)
		}
		if conn != nil {
			if err := conn.WriteMessage(websocket.TextMessage, b); err != nil {
				log.Println(err)
				return
			}
		}
  }
}

func subToBlockHeader(blocksCollection *mongo.Collection, transactionsCollection *mongo.Collection) {
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/7fee20f62b264b7ab41d4fb69dde3c76")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)

	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
			case err := <-sub.Err():
				log.Fatal(err)
			case header := <-headers:			
				block, err := client.BlockByHash(context.Background(), header.Hash())
				if err != nil {
					fmt.Println(err)
					continue
				}
			
			blockData, transactions := typehelper.StructureBlockData(block)
			
			select {
				case broadcast <- *blockData:
					fmt.Println("received")
					addBlockToDb(blockData, blocksCollection, transactionsCollection, transactions)	
				default:
					addBlockToDb(blockData, blocksCollection, transactionsCollection, transactions)	
			}
		}
	}
}


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


	http.HandleFunc("/ws", serveWs)
	routes.SetupBlockRoutes(mongoClient)
	routes.SetupTransactionRoutes(mongoClient)
	go subToBlockHeader(blocksCollection, transactionsCollection)
	log.Fatal(http.ListenAndServe(":5000", nil))

}



func addBlockToDb(b *typehelper.BlockData, c *mongo.Collection, t *mongo.Collection, txs []interface{}) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	blockDataResult, err2 := c.InsertOne(ctx, *b)
	if err2 != nil {
		log.Fatal(err2)
	}
	if len(txs) != 0 {
		_, err3 := t.InsertMany(ctx, txs)
		if err3 != nil {
			fmt.Println("Zero transactions")
		}
	}
	spew.Dump(blockDataResult)
	}


