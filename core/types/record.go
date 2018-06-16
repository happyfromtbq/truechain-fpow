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
