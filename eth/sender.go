package eth

import (
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

var (
	sendInterval = 1 * time.Second // Time interval to send record

	sendAddrHex  = "970e8128ab834e8eac17ab8e3812f010678cf791"
	sendPrivHex  = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"

	recvAddrHex = "68f2517b6c597ede0ae7c0559cdd4a84fd08c928"
)

type RecordSender struct {
	txpool *core.TxPool
	signer types.Signer

	sendAccout *ecdsa.PrivateKey
	recvAddr   common.Address
}

// newTestTransaction create a new dummy transaction.
func (sender *RecordSender) newSendTransaction(nonce uint64, datasize int) *types.Transaction {

	tx := types.NewTransaction(nonce, sender.recvAddr, big.NewInt(1e+18), 100000, big.NewInt(1e+15), make([]byte, datasize))
	tx, _ = types.SignTx(tx, sender.signer, sender.sendAccout)
	return tx
}

func (sender *RecordSender) send() {
	var nonce uint64
	ticker := time.NewTicker(sendInterval)
	defer ticker.Stop()

	number := big.NewInt(0)

	for {
		select {
		case <-ticker.C:
			//record
			//nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte
			//*types.Transaction
			txs := make([]*types.Transaction, 1)
			for i := range txs {
				nonce++
				txs[i] = sender.newSendTransaction(nonce, 0)
			}

			//header
			number = number.Add(number, common.Big1)

			prd := types.NewRecord(number, txs, nil)

			var records []*types.PbftRecord
			records = append(records, prd)

			sender.txpool.AddRemoteRecords(records)
		}
	}
}

func (sender *RecordSender) Start() {
	go sender.send()
}

func NewSender(txpool *core.TxPool, chainconfig *params.ChainConfig) *RecordSender {
	acc, _ := crypto.HexToECDSA(sendPrivHex)

	// TODO: get key and account address to send

	sendRecord := &RecordSender{
		txpool:     txpool,
		signer:     types.NewEIP155Signer(chainconfig.ChainId),
		sendAccout: acc,
		recvAddr:   common.HexToAddress(recvAddrHex),
	}

	return sendRecord
}
