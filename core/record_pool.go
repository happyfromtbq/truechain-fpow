// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core


import (

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// AddLocals enqueues a batch of transactions into the pool if they are valid,
// marking the senders as a local ones in the mean time, ensuring they go around
// the local pricing constraints.
func (pool *TxPool) AddRemoteRecords(records []*types.PbftRecord) []error {
	pool.muRecord.Lock()
	defer pool.muRecord.Unlock()

	for _, record := range records {

		// check whether exits the fruit has the same record
		pre := pool.records[record.Hash()]
		if pre == nil {
			pool.records[record.Hash()] = types.CopyRecord(record)
		}
	}

	return nil
}


// Pending retrieves one currently record.
// The returned record is a copy and can be freely modified by calling code.
func (pool *TxPool) PendingRecords() (*types.PbftRecord, error) {
	pool.muRecord.Lock()
	defer pool.muRecord.Unlock()
	//TODO: get the first record in the pool

	return nil, nil
}


// SubscribeNewRecordsEvent registers a subscription of NewRecordEvent and
// starts sending event to the given channel.
func (pool *TxPool) SubscribeNewRecordEvent(ch chan<- NewRecordEvent) event.Subscription {
	return pool.scope.Track(pool.recordFeed.Subscribe(ch))
}
