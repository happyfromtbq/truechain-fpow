package types

import (
	"time"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type PbftRecordHeader struct {
	Number   *big.Int
	Hash     common.Hash
	TxHash	common.Hash
	GasLimit *big.Int
	GasUsed  *big.Int
	Time     *big.Int
}

type PbftRecord struct {
	header       *PbftRecordHeader
	transactions Transactions
	sig          []*string
}


func (r *PbftRecord) Hash() common.Hash {
	return r.header.Hash
}

func (r *PbftRecord) Number() *big.Int {
	return r.header.Number
 }


func (r *PbftRecord) Header() *PbftRecordHeader { return r.header }

func (r *PbftRecord) Transactions() Transactions { return r.transactions }


func CopyRecord(r *PbftRecord) *PbftRecord {
	record := &PbftRecord{
		header: &PbftRecordHeader{
			Number: r.header.Number,
			Hash:	r.header.Hash,
			TxHash:	r.header.TxHash,
			GasLimit: r.header.GasLimit,
			GasUsed: r.header.GasUsed,
			Time: 	r.header.Time,
		},
	}

	if len(r.transactions) != 0 {
		record.transactions = make(Transactions, len(r.transactions))
		copy(record.transactions, r.transactions)
	}

	// TODO: copy sigs

	return record
}


func NewRecord(number *big.Int, txs []*Transaction, sig []*string) *PbftRecord {

	r := &PbftRecord{
		header: &PbftRecordHeader {
			Number: number,
			Time: big.NewInt(time.Now().Unix()),
		},
	}

	if len(txs) == 0 {
		r.header.TxHash = EmptyRootHash
	} else {
		r.header.TxHash = DeriveSha(Transactions(txs))
		r.transactions = make(Transactions, len(txs))
		copy(r.transactions, txs)
	}

	r.header.Hash = rlpHash([]interface{}{
		r.header.Number,
		r.header.TxHash,
		r.header.GasLimit,
		r.header.GasUsed,
		r.header.Time,
		r.sig,
	})

	return r
}
