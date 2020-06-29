package simulation

import (
	"github.com/Tao-Network/tao2/crypto"
	"math/big"
	"os"
)

var (
	RpcEndpoint = "http://127.0.0.1:4545/"
	MinApply    = big.NewInt(0).Mul(big.NewInt(1000), big.NewInt(100000000000000000)) // 100 TAO
	Cap         = big.NewInt(0).Mul(big.NewInt(10000000000000), big.NewInt(10000000000000))
	Fee         = big.NewInt(100)

	MainKey, _ = crypto.HexToECDSA(os.Getenv("MAIN_ADDRESS_KEY"))
	MainAddr   = crypto.PubkeyToAddress(MainKey.PublicKey) //0x17F2beD710ba50Ed27aEa52fc4bD7Bda5ED4a037

	AirdropKey, _  = crypto.HexToECDSA(os.Getenv("RELAYER_COINBASE_KEY"))
	AirdropAddr    = crypto.PubkeyToAddress(AirdropKey.PublicKey) // 0x0D3ab14BBaD3D99F4203bd7a11aCB94882050E7e
	AirDropAmount  = big.NewInt(10000000000)
	TransferAmount = big.NewInt(100000)

	ReceiverKey, _ = crypto.HexToECDSA(os.Getenv("RELAYER_OWNER_KEY"))
	ReceiverAddr   = crypto.PubkeyToAddress(ReceiverKey.PublicKey) //0x703c4b2bD70c169f5717101CaeE543299Fc946C7
)
