package pkg

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	sdkcommon "github.com/harmony-one/go-sdk/pkg/common"
	sdkrpc "github.com/harmony-one/go-sdk/pkg/rpc"
	v1 "github.com/harmony-one/go-sdk/pkg/rpc/v1"
	"github.com/harmony-one/go-sdk/pkg/transaction"
	"github.com/harmony-one/harmony/accounts/abi"
	"github.com/harmony-one/harmony/accounts/keystore"
	"github.com/harmony-one/harmony/numeric"
)

const chainId = 2

type Chain struct {
	config        ChainConfig
	client        *Client
	ethclient     *ethclient.Client
	ibcHandlerABI *abi.ABI
	keyStore      *keystore.KeyStore
}

type Client struct {
	messenger *sdkrpc.HTTPMessenger
}

// from yui-ibc-solidity
// IBCMsgsMsgCreateClient is an auto generated low-level Go binding around an user-defined struct.
type IBCMsgsMsgCreateClient struct {
	ClientType          string
	Height              uint64
	ClientStateBytes    []byte
	ConsensusStateBytes []byte
}

func NewHarmonyClient(endpoint string) *Client {
	messenger := sdkrpc.NewHTTPHandler(endpoint)
	return &Client{
		messenger: messenger,
	}
}

func NewChain(chainConfig ChainConfig) (*Chain, error) {
	client := NewHarmonyClient(chainConfig.Endpoint)
	ethclient, err := NewETHClient(chainConfig.Endpoint)
	if err != nil {
		return nil, err
	}

	ihABI, err := abi.JSON(strings.NewReader(chainConfig.IbcHandlerABI))
	if err != nil {
		return nil, err
	}

	keyStore := sdkcommon.KeyStoreForPath(chainConfig.KeyStorePath)
	if !keyStore.HasAddress(crypto.PubkeyToAddress(chainConfig.PrivKey.PublicKey)) {
		_, err = keyStore.ImportECDSA(chainConfig.PrivKey, "")
		if err != nil {
			return nil, err
		}
	}

	return &Chain{
		config:        chainConfig,
		client:        client,
		ethclient:     ethclient,
		ibcHandlerABI: &ihABI,
		keyStore:      keyStore,
	}, nil
}

func (c *Chain) GetHostAddress() (common.Address, error) {
	res, err := c.Call("getHostAddress")
	if err != nil {
		return common.Address{}, err
	}
	s, ok := res.(string)
	if !ok {
		return common.Address{}, errors.New("invalid result")
	}
	return common.HexToAddress(s), nil
}

func (c *Chain) Call(method string, params ...interface{}) (interface{}, error) {
	input, err := c.ibcHandlerABI.Pack(method, params...)
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{
		To:       &c.config.IbcHandlerAddress,
		Data:     input,
		Gas:      c.config.GasLimit,
		GasPrice: c.config.GasPrice.Int,
		Value:    big.NewInt(0),
	}
	// TODO enable to specify block number
	rep, err := c.client.messenger.SendRPC(v1.Method.Call, []interface{}{toCallArg(msg), "latest"})
	if err != nil {
		return nil, err
	}
	result, ok := rep["result"]
	if !ok {
		return nil, errors.New("invalid request")
	}
	return result, nil
}

func (c *Chain) Tx(method string, params ...interface{}) (string, error) {
	input, err := c.ibcHandlerABI.Pack(method, params...)
	if err != nil {
		return "", err
	}
	accounts := c.keyStore.Accounts()
	if len(accounts) == 0 {
		return "", errors.New("empty keystore")
	}
	account := accounts[0]
	if err = c.keyStore.Unlock(account, ""); err != nil {
		return "", err
	}
	controller := transaction.NewController(c.client.messenger, c.keyStore, &account, c.config.ChainID)
	// XXX or pending nonce
	nonce := transaction.GetNextNonce(account.Address.Hex(), c.client.messenger)
	addr := c.config.IbcHandlerAddress.Hex()
	err = controller.ExecuteTransaction(nonce, c.config.GasLimit, &addr, c.config.ShardID, c.config.ShardID, numeric.NewDec(0), c.config.GasPrice, input)
	if err != nil {
		return "", err
	}
	if err = c.keyStore.Lock(account.Address); err != nil {
		panic(err)
	}
	txHash := controller.TransactionHash()
	if txHash == nil {
		return "", errors.New("can't get txhash")
	}
	return *txHash, nil
}

func NewETHClient(endpoint string) (*ethclient.Client, error) {
	conn, err := rpc.DialHTTP(endpoint)
	if err != nil {
		return nil, err
	}
	return ethclient.NewClient(conn), nil
}
