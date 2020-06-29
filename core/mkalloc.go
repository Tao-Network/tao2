// Copyright 2017 The go-ethereum Authors
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

// +build none

/*

   The mkalloc tool creates the genesis allocation constants in genesis_alloc.go
   It outputs a const declaration that contains an RLP-encoded list of (address, balance) tuples.

       go run mkalloc.go genesis.json

*/
package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/Tao-Network/tao2/common"
	"github.com/Tao-Network/tao2/core"
	"github.com/Tao-Network/tao2/rlp"
)

type allocItem struct {
	Addr    *big.Int                    `json:"code"       rlp:"nil"`
	Balance *big.Int                    `json:"code"       rlp:"nil"`
	Code    []byte                      `json:"code"       rlp:"nil"`
	Storage map[common.Hash]common.Hash `json:"storage"       rlp:"nil"`
}

type allocList []allocItem

func (a allocList) Len() int           { return len(a) }
func (a allocList) Less(i, j int) bool { return a[i].Addr.Cmp(a[j].Addr) < 0 }
func (a allocList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func makelist(g *core.Genesis) []core.GenesisAccount {
	a := make([]core.GenesisAccount, len(g.Alloc))
	for addr, account := range g.Alloc {
		if account.Nonce != 0 {
			panic(fmt.Sprintf("can't encode account %x", addr))
		}
		bigAddr := new(big.Int).SetBytes(addr.Bytes())
		aItem := allocItem{Addr: bigAddr, Balance: account.Balance}
		if len(account.Code) > 0 {
			aItem.Code = account.Code
		}
		if len(account.Storage) > 0 {
			aItem.Storage = map[common.Hash]common.Hash{}
			for k, v := range account.Storage {
				aItem.Storage[k] = v
			}
		}
		for k, _ := range aItem.Storage {
			fmt.Println(k.Hex())
		}

		a = append(a, account)
	}
	//sort.Sort(a)
	return a
}

func makealloc(g *core.Genesis) string {
	a := makelist(g)

	data, err := rlp.EncodeToBytes(a)
	if err != nil {
		panic(err)
	}
	return strconv.QuoteToASCII(string(data))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: mkalloc genesis.json")
		os.Exit(1)
	}

	g := new(core.Genesis)
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(file).Decode(g); err != nil {
		panic(err)
	}

	fmt.Println("const allocData =", makealloc(g))
}