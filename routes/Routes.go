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


)


func SetupBlockRoutes(c *mongo.Client) {

	filter := bson.D{{"hash", "0x8e2e724849f406b21bfe3d9661b5ac326237d3a362aa9195c49195727214f72a"}}
	blocksCollection := c.Database("blockHistoryDB").Collection("blocks")
	var block typehelper.BlockData
	http.HandleFunc("/api/block", func(w http.ResponseWriter, r *http.Request) {
		err := blocksCollection.FindOne(context.TODO(), filter).Decode(&block)
		// err := blocksCollection.FindId(primitive.ObjectIDFromHex).One(&block)
		if err != nil {
				log.Fatal(err)
		}
		data, _ := json.Marshal(block)
		w.Write(data)
		fmt.Printf("Found a single document: %+v\n", block)

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


func SetupTransactionRoutes(c *mongo.Client) {
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
}

