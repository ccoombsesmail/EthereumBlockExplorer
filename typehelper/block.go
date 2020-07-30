package typehelper

import (
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)


func StructureBlockData(b *types.Block) (*BlockData, []interface{})  {

		formatedTransactions := make([]TransactionData, 0)
		var transactionsInt []interface{}
		transactions := b.Transactions()
		for i:=0; i < len(transactions); i++ {
			tx := transactions[i]
			formatedTransactions = append(formatedTransactions, *StructureTransactionData(tx))
			transactionsInt = append(transactionsInt, *StructureTransactionData(tx))
		}
		
			// strconv.FormatUint(b.Nonce(), 10)
		block := &BlockData {
			Hash: b.Hash().Hex(),
			Number: b.Number().Uint64(),
			Timestamp: time.Unix(int64(b.Time()), 0),
			Nonce:  strconv.FormatUint(b.Nonce(), 10),             
			Transactions: formatedTransactions }
		return block, transactionsInt
	}


	type BlockData struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Hash string  `bson:"hash"`
	Number uint64  `bson:"number"`
	Timestamp time.Time  `bson:"timestamp"`
	Nonce string  `bson:"nonce"`
	Transactions []TransactionData `bson:"transactions"`
}