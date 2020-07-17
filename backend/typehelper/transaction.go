package typehelper

import (
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"fmt"

)

func StructureTransactionData(tx *types.Transaction) *TransactionData {
	signer := types.NewEIP155Signer(tx.ChainId())
	sender, err := signer.Sender(tx)
	if err != nil {
		fmt.Printf("ERROOOR")
	}
	var to string
	if tx.To() == nil {
		to = ""
	} else {
		to = tx.To().String()
	}
	transaction := &TransactionData {
				From: sender.String(),
				To: to,
				Value: strconv.FormatUint(tx.Value().Uint64(), 10)}
	return transaction
}

	
type TransactionData struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	From string  `bson:"from"`
	To string    `bson:"to"`
	Value string  `bson:"value"`

}