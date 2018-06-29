package eth
import (
	"time"
	"math/rand"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
)
type SendRecord struct {
	txpool  *core.TxPool
}
func NewRecord(txpool  *core.TxPool) *SendRecord{
	sendRecord:= &SendRecord{
		txpool: txpool,
	}
	go sendRecord.StartRecord(txpool)
	return sendRecord
}
func (self *SendRecord) StartRecord(txpool  *core.TxPool){
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		//header
		var ph *types.PbftRecordHeader
		gap:=rand.Int63()
		ph=types.NewPbftRecordHeader(big.NewInt(rand.Int63()),common.BigToHash(big.NewInt(rand.Int63())),big.NewInt(gap),big.NewInt(gap-100),big.NewInt(time.Now().Unix()))

		//record
		var prd *types.PbftRecord
		var records []*types.PbftRecord
		//nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte
		//*types.Transaction
		ts:=types.NewTransaction(uint64(10000),common.BytesToAddress([]byte("this")),big.NewInt(0),uint64(1000),big.NewInt(0),[]byte("test"))
		var tss types.Transactions
		tss=append(tss,ts)
		var arr [] *string

		var data string  = "test"
		arr=append(arr,&data)
		//header *PbftRecordHeader, transactions Transactions,sig []*string
		prd =types.NewPbftRecord(ph,tss,arr)
		fmt.Print(prd)
		records=append(records,prd)
		txpool.AddRemoteRecords(records)
	}
}