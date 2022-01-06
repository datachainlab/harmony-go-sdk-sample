package main

import (
	"fmt"

	"github.com/datachainlab/harmony-go-sdk-sample/pkg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	sdkcommon "github.com/harmony-one/go-sdk/pkg/common"
	"github.com/harmony-one/harmony/numeric"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
)

const (
	endpoint            = "http://localhost:9599"
	shardID             = uint32(0)
	ibcHandlerAddr      = "0x65A4191e03E681a35F0B524fbf3ac9323A17F545"
	gasLimit            = uint64(80000000)
	gasPrice            = 1
	privKey             = "1f84c95ac16e6a50f08d44c7bde7aff8742212fda6e4321fde48bf83bef266dc"
	defaultKeyStorePath = "./keystore"
)

var (
	chainID = sdkcommon.Chain.TestNet
)

func main() {
	key, err := crypto.HexToECDSA(privKey)
	if err != nil {
		panic(err)
	}
	config := pkg.ChainConfig{
		Endpoint:          endpoint,
		ShardID:           shardID,
		ChainID:           chainID,
		IbcHandlerAddress: common.HexToAddress(ibcHandlerAddr),
		PrivKey:           key,
		KeyStorePath:      defaultKeyStorePath,
		GasLimit:          gasLimit,
		GasPrice:          numeric.NewDec(gasPrice),
	}
	chain, err := pkg.NewChain(config)
	if err != nil {
		panic(err)
	}

	/* call */
	hostAddr, err := chain.GetHostAddress()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("IBCHost address: %s\n", hostAddr.Hex())

	/* send tx */
	msg := ibchandler.IBCMsgsMsgCreateClient{
		ClientType:          "mock-client",
		Height:              uint64(1),
		ClientStateBytes:    []byte("invalid"),
		ConsensusStateBytes: []byte("invalid"),
	}
	txHash, err := chain.Tx("createClient", msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("TxHash: %s\n", txHash)
}
