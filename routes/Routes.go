package routes

import (
	"context"
	"fmt"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/davecgh/go-spew/spew"
	"encoding/json"
	"EthereumBlockExplorer/typehelper"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"strconv"


)


func SetupBlockRoutes(c *mongo.Client) {

	filter := bson.D{{"hash", "0xbf2f5e0820ea24163176ddff953ee950ef940ff7aad5258427d25b2131b294fa"}}
	blocksCollection := c.Database("blockHistoryDB").Collection("blocks")
	var block typehelper.BlockData
	http.HandleFunc("/api/block", func(w http.ResponseWriter, r *http.Request) {
		err := blocksCollection.FindOne(context.TODO(), filter).Decode(&block)
		if err != nil {
				spew.Dump(blocksCollection)
				spew.Dump(err)
		}
		data, _ := json.Marshal(block)
		w.Write(data)
		fmt.Printf("Found a single document: %+v\n", block)

	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) { 
		spew.Dump(c)
	})


	http.HandleFunc("/api/blocks", func(w http.ResponseWriter, r *http.Request) {
		options := options.Find()
		options.SetSort(bson.D{{"_id", -1}})
		options.SetLimit(100)
		cursor, err := blocksCollection.Find(context.Background(), bson.D{}, options)
		if err != nil {
				spew.Dump(err)
		}
		blocks := make([]typehelper.BlockData, 0)
		for cursor.Next(context.Background()) {
			var blockData typehelper.BlockData 
			err = cursor.Decode(&blockData)
			if err != nil {
				log.Fatal(err)
			}
			blocks = append(blocks, blockData)
		}

		data, _ := json.Marshal(blocks)
		w.Write(data)
})


	http.HandleFunc("/api/recentblocks", func(w http.ResponseWriter, r *http.Request) {
		options := options.Find()
		options.SetSort(bson.D{{"_id", -1}})
		options.SetLimit(4)

		cursor, err := blocksCollection.Find(context.Background(), bson.D{}, options)
		if err != nil {
			spew.Dump(err)
			spew.Dump("api/recent")


		}
		blocks := make([]typehelper.BlockData, 0)
		for cursor.Next(context.Background()) {
			var blockData typehelper.BlockData 
			err = cursor.Decode(&blockData)
			if err != nil {
				spew.Dump(err)
				spew.Dump("api/recent2")

			}
			blocks = append(blocks, blockData)
		}

		data, _ := json.Marshal(blocks)
		w.Write(data)
	})

}


func SetupTransactionRoutes(c *mongo.Client, ethClient *ethclient.Client) {
	transactionsCollection := c.Database("blockHistoryDB").Collection("transactions")

	http.HandleFunc("/api/transactions", func(w http.ResponseWriter, r *http.Request) {
		searchQuery := ""
		if len(r.URL.Query()["searchQuery"]) != 0 {
			searchQuery = r.URL.Query()["searchQuery"][0]
		}
		spew.Dump(searchQuery)
		

		cursor, err := transactionsCollection.Find(context.Background(), bson.M{
		"$or": []bson.M{
			bson.M{"hash": searchQuery},
			bson.M{"from": searchQuery},
			bson.M{"to": searchQuery}}})
		if err != nil {
				spew.Dump("Could not find any txs")
				w.Write([]byte("Could not find any txs"))
		} else {
			transactions := make([]typehelper.TransactionData, 0)
			for cursor.Next(context.Background()) {
				var TransactionData typehelper.TransactionData 
				err = cursor.Decode(&TransactionData)
				if err != nil {
					log.Fatal(err)
				}
				transactions = append(transactions, TransactionData)
			}

			data, _ := json.Marshal(transactions)
			w.Write(data)
		}
	})

	http.HandleFunc("/api/gasused", func(w http.ResponseWriter, r *http.Request) { 
		txHash := ""
		if len(r.URL.Query()["txhash"]) != 0 {
			txHash = r.URL.Query()["txhash"][0]
		}
		tx, err := ethClient.TransactionReceipt(context.Background(), common.HexToHash(txHash))
		if err != nil {
			spew.Dump(err)
		} else {
			w.Write([]byte(strconv.FormatUint(tx.GasUsed, 10)))
		}
	})

}

