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

func (r *PbftRecord) TxHash() common.Hash {return r.header.TxHash}

func (r *PbftRecord) Transactions() Transactions { return r.transactions }

func (r *PbftRecord) CalcHash() common.Hash {
	return rlpHash([]interface{}{
		r.header.Number,
		r.header.TxHash,
		r.header.GasLimit,
		r.header.GasUsed,
		r.header.Time,
		r.sig,
	})
}


func CopyRecord(r *PbftRecord) *PbftRecord {
	header := *r.header
	if header.Time = new(big.Int); r.header.Time != nil {
		header.Time.Set(r.header.Time)
	}
	if header.Number = new(big.Int); r.header.Number != nil {
		header.Number.Set(r.header.Number)
	}
	if header.GasLimit = new(big.Int); r.header.GasLimit != nil {
		header.GasLimit.Set(r.header.GasLimit)
	}
	if header.GasUsed = new(big.Int); r.header.GasUsed != nil {
		header.GasUsed.Set(r.header.GasUsed)
	}

	record := &PbftRecord{
		header: &header,
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

	r.header.Hash = r.CalcHash()

	return r
}
