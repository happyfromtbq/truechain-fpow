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
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

type PoolInfo struct {
	Hash   common.Hash
	Number *big.Int
}

func (pool *TxPool) ProcessFruits() {

}

// AddLocals enqueues a batch of transactions into the pool if they are valid,
// marking the senders as a local ones in the mean time, ensuring they go around
// the local pricing constraints.
func (pool *TxPool) AddRemoteFruits(fruits []*types.Block) []error {
	pool.muFruit.Lock()
	defer pool.muFruit.Unlock()

	for _, fruit := range fruits {
		// check the fruit

		// check whether exits the fruit has the same record
		pre := pool.fruits[fruit.Hash()]
		if pre == nil {
			pool.fruits[fruit.Hash()] = types.CopyFruit(fruit)
		} else {
			// contain the fruit who has the smaller hash
			if fruit.Hash().Big().Cmp(pre.Hash().Big()) > 0 {
				pool.fruits[fruit.Hash()] = types.CopyFruit(fruit)
			}
		}
	}

	return nil
}

// Pending retrieves all currently processable allFruits, sorted by record number.
// The returned fruit set is a copy and can be freely modified by calling code.
func (pool *TxPool) PendingFruits() (map[common.Hash]types.Block, error) {
	pool.muFruit.Lock()
	defer pool.muFruit.Unlock()

	pending := make(map[common.Hash]types.Block)
	for addr, list := range pool.fruits {
		pending[addr] = *types.CopyFruit(list)
	}

	return pending, nil
}

// SubscribeNewFruitsEvent registers a subscription of NewFruitEvent and
// starts sending event to the given channel.
func (pool *TxPool) SubscribeNewFruitsEvent(ch chan<- NewFruitEvent) event.Subscription {
	return pool.scope.Track(pool.fruitFeed.Subscribe(ch))
}
