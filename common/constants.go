package common

import (
	"math/big"
)

const (
	RewardMasterPercent        = 40
	RewardVoterPercent         = 50
	RewardFoundationPercent    = 10
	HexSignMethod              = "e341eaa4"
	HexSetSecret               = "34d38600"
	HexSetOpening              = "e11f5ba2"
	EpocBlockSecret            = 300
	EpocBlockOpening           = 330
	EpocBlockRandomize         = 360
	MaxMasternodes             = 150
	LimitPenaltyEpoch          = 4
	BlocksPerYear              = uint64(6307200)
	LimitThresholdNonceInQueue = 10
	DefaultMinGasPrice         = 250000000
	MergeSignRange             = 15
	RangeReturnSigner          = 150
	MinimunMinerBlockPerEpoch  = 1
	IgnoreSignerCheckBlock     = uint64(1)
	OneYear                    = uint64(365 * 86400)
	LiquidateLendingTradeBlock = uint64(100)
)

var Rewound = uint64(0)

// hardforks
var TIP2019Block = big.NewInt(1)
var TIPSigning = big.NewInt(1)
var TIPRandomize = big.NewInt(1)
var BlackListHFNumber = uint64(1)
var TIPTaoX = big.NewInt(1555200)
var TIPTaoXTestnet = big.NewInt(1555200)
var TIPTaoXLending = big.NewInt(4665600)

var IsTestnet bool = false
var StoreRewardFolder string
var RollbackHash Hash
var BasePrice = big.NewInt(1000000000000000000)                         // 1
var RelayerLockedFund = big.NewInt(20000)                               // 20000 TOMO
var RelayerFee = big.NewInt(1000000000000000)                           // 0.001
var TaoXBaseFee = big.NewInt(10000)                                    // 1 / TaoXBaseFee
var RelayerCancelFee = big.NewInt(100000000000000)                      // 0.0001
var TaoXBaseCancelFee = new(big.Int).Mul(TaoXBaseFee, big.NewInt(10)) // 1/ (TaoXBaseFee *10)
var RelayerLendingFee = big.NewInt(10000000000000000)                   // 0.01
var RelayerLendingCancelFee = big.NewInt(1000000000000000)              // 0.001
var BaseLendingInterest = big.NewInt(100000000)                         // 1e8

var MinGasPrice = big.NewInt(DefaultMinGasPrice)
var RelayerRegistrationSMC = "0x7ed7f65d63362d35e9977fc64c6f6b2812563108"
var RelayerRegistrationSMCTestnet = "0x7ed7f65d63362d35e9977fc64c6f6b2812563108"
var LendingRegistrationSMC = "0xf0387a1243a593c90a518894341135cbfe303ff5"
var LendingRegistrationSMCTestnet = "0xf0387a1243a593c90a518894341135cbfe303ff5"
var TRC21IssuerSMCTestNet = HexToAddress("0x81e4d65a002417745ff361c225ce533a7300e51c")
var TRC21IssuerSMC = HexToAddress("0x81e4d65a002417745ff361c225ce533a7300e51c")
var TaoXListingSMC = HexToAddress("0x006704b9b461c4b4ffc932208f0650013fe0bd3c")
var TaoXListingSMCTestNet = HexToAddress("0x006704b9b461c4b4ffc932208f0650013fe0bd3c")
var TRC21GasPriceBefore = big.NewInt(2500)
var TRC21GasPrice = big.NewInt(250000000)
var RateTopUp = big.NewInt(90) // 90%
var BaseTopUp = big.NewInt(100)
var BaseRecall = big.NewInt(100)
var Blacklist = map[Address]bool{
}
var TIPTRC21Fee = big.NewInt(13523400)
var LimitTimeFinality = uint64(30) // limit in 30 block
