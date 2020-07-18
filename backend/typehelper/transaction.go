package typehelper

import (
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"fmt"
	"math/big"

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
	price := new(big.Int).SetUint64(tx.Gas())
	var txfee big.Int
	txfee.Mul(price, tx.GasPrice())

	transaction := &TransactionData {
		Hash : tx.Hash().String(),
		From: sender.String(),
		To: to,
		Value: strconv.FormatUint(tx.Value().Uint64(), 10),
		Size: tx.Size().String(),
		GasLimit: tx.Gas(),
		TxFee: txfee.String()}
	return transaction
}

	
type TransactionData struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Hash string   `bson:"hash"`
	From string   `bson:"from"`
	To string     `bson:"to"`
	Value string  `bson:"value"`
	Size string  `bson:"size"`
	GasLimit uint64  `bson:"gaslimit"`
	TxFee string `bson:"txfee"`



}