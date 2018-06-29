package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type PbftRecordHeader struct {
	Number   *big.Int
	Hash     common.Hash
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


func CopyRecord(r *PbftRecord) *PbftRecord {
	// TODO: copy all record fields
	header := NewPbftRecordHeader(r.header.Number, r.header.Hash, r.header.Time)
	record := &PbftRecord{
		header: header,
	}

	if len(r.transactions) != 0 {
		record.transactions = make(Transactions, len(r.transactions))
		copy(record.transactions, r.transactions)
	}

	return record
}

func NewPbftRecord(header *PbftRecordHeader, transactions Transactions,sig []*string) *PbftRecord {
	return &PbftRecord{
		header:       header,
		transactions: transactions,
		sig:          sig,
	}
}

func NewPbftRecordHeader(Number *big.Int,Hash common.Hash,Time *big.Int) *PbftRecordHeader {
	return &PbftRecordHeader{
		Number:   Number,
		Hash:     Hash,
		Time:     Time,
	}
}