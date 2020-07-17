package typehelper

import (
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func StructureBlockData(b *types.Block) (*BlockData, []interface{})  {
		var nonce = b.Nonce()
		if nonce >= 184467440737 {
			nonce = 18446744     
		}

		formatedTransactions := make([]TransactionData, 0)
		var transactionsInt []interface{}
		transactions := b.Transactions()
		for i:=0; i < len(transactions); i++ {
			tx := transactions[i]
			formatedTransactions = append(formatedTransactions, *StructureTransactionData(tx))
			transactionsInt = append(transactionsInt, *StructureTransactionData(tx))
		}
		
			
		block := &BlockData {
			Hash: b.Hash().Hex(),
			Number: b.Number().Uint64(),
			Timestamp: b.Time(),
			Nonce: nonce,             
			Transactions: formatedTransactions }
		return block, transactionsInt
	}


	type BlockData struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Hash string  `bson:"hash"`
	Number uint64  `bson:"number,omitempty"`
	Timestamp uint64  `bson:"timestamp,omitempty"`
	Nonce uint64  `bson:"nonce,omitempty"`
	Transactions []TransactionData `bson:"transactions,omitempty"`
}