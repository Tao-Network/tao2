// Copyright 2015 The go-ethereum Authors
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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Tao Bootnodes Mainnet
	"enode://7b3717f1318eb10e0e04a2590741fd9521a80d8bf5bb715d6507618d930f9d94a2a8edeff812e8d0e4406fa573d176c004e71e7f142aeeb088756875852c68c8@149.28.69.174:20202",
	"enode://063f4c05478262ddcd9fdd6a8b894a312972a23e0a9aa277fb34f9728e7fec2f1b3e35ce3f845c0f6a1bc22a476aad3ef68a85fe79cc73a71c8bfdb03482887f@8.3.29.246:20202",
	"enode://01de96e04af0d35dcd7aef53450f80356efc0c56faceaa59c29ce38f55f002e3e9252154ca325fbef1c9926e79e90ef416f9b0de5ca1888944f6ec6dcef997be@45.32.227.55:20202",
	"enode://f85a8d102d8a177a661044139cf6368296f8a55bfd1cec347cba0752b9d444ddc2ce03c009dd27dfdf631ef8db048f89b8dd5f30660a9750ba1eac2c569fb8c5@140.82.50.27:20202",
	"enode://c45b5055f551c17bbf777c80a3c74a54a8b6e6d6cc6502bad6a1bd35bdc5773b71338bd9efc9162e48f8e50b913d5b7187599f795443679c8ab49c4dfc378daa@149.248.44.254:20202",
	"enode://c5ca2fb91ae21360b1a94148bf8a5754b91e159f1eb0e7e3590ea7729f7b0f119df81146b5fd43105ddd55f7b58d19fa285ff83e5a450fdf048ceeb6482279e0@45.77.85.124:20202",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	// Tao Bootnodes Testnet
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
}
