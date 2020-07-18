package websockets

import (
	"context"
	"fmt"
	"log"
	"time"
	"net/http"
	"encoding/json"
	"EthereumBlockExplorer/typehelper"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/websocket"
)
var Client *ethclient.Client
var broadcast = make(chan typehelper.BlockData)

var upgrader = websocket.Upgrader{
  ReadBufferSize:  1024,
  WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool { return true },
}


func ServeWs(w http.ResponseWriter, r *http.Request) {
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

func SubToBlockHeader(blocksCollection *mongo.Collection, transactionsCollection *mongo.Collection) {
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/7fee20f62b264b7ab41d4fb69dde3c76")
	if err != nil {
		log.Fatal(err)
	}

	// tx, err4 := client.TransactionReceipt(context.Background(), common.HexToHash("0x12400a58cf1b6a3b48a6ba2630be571f3e98b190ad9eb4a71a7e168fa1dd817e"))
	// if err4 != nil {
	// 	spew.Dump(err4)
	// }
	// spew.Dump(tx.GasUsed)

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


func addBlockToDb(b *typehelper.BlockData, c *mongo.Collection, t *mongo.Collection, txs []interface{}) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	blockDataResult, err2 := c.InsertOne(ctx, *b)
	if err2 != nil {
		log.Fatal(err2)
	}
	if len(txs) != 0 {
		_, err3 := t.InsertMany(ctx, txs)
		if err3 != nil {
			log.Fatal(err3)
		}
	}
	spew.Dump(blockDataResult)
	}

