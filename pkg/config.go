package pkg

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	sdkcommon "github.com/harmony-one/go-sdk/pkg/common"
	"github.com/harmony-one/harmony/numeric"
)

type ChainConfig struct {
	Endpoint          string
	ChainID           sdkcommon.ChainID
	ShardID           uint32
	IbcHandlerAddress common.Address
	IbcHandlerABI     string
	GasLimit          uint64
	GasPrice          numeric.Dec
	PrivKey           *ecdsa.PrivateKey
	KeyStorePath      string
}
