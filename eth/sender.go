package eth
import (
	"fmt"
	"time"
	"math/rand"
	"math/big"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"

)


var (
	sendInterval = 1 * time.Second // Time interval to send record
)

var sendAccount, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

type RecordSender struct {
	txpool  *core.TxPool
}


// newTestTransaction create a new dummy transaction.
func newSendTransaction(from *ecdsa.PrivateKey, nonce uint64, datasize int) *types.Transaction {
	tx := types.NewTransaction(nonce, common.Address{}, big.NewInt(0), 100000, big.NewInt(0), make([]byte, datasize))
	tx, _ = types.SignTx(tx, types.HomesteadSigner{}, from)
	return tx
}

func (sender *RecordSender)send() {
	var nonce uint64
	ticker := time.NewTicker(sendInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			//record
			//nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte
			//*types.Transaction
			txs := make([]*types.Transaction, 1)
			for i := range txs {
				nonce++
				txs[i] = newSendTransaction(sendAccount, uint64(nonce), 0)
			}

			//header
			ph := types.NewPbftRecordHeader(big.NewInt(rand.Int63()), common.BigToHash(big.NewInt(rand.Int63())), big.NewInt(time.Now().Unix()))

			prd := types.NewPbftRecord(ph, txs, nil)

			fmt.Print(prd)

			var records []*types.PbftRecord
			records = append(records, prd)

			sender.txpool.AddRemoteRecords(records)
		}
	}
}


func (sender *RecordSender)Start() {
	go sender.send()
}

func NewSender(txpool  *core.TxPool) *RecordSender {
	sendRecord:= &RecordSender{
		txpool: txpool,
	}
	// TODO: get key and account address to send

	return sendRecord
}
