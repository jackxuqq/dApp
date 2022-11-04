package ethereum

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MintEvent struct {
	Token *big.Int
}

type TransferEvent struct {
	From   common.Address
	To     common.Address
	Token  *big.Int
	Amount *big.Int
	Ext    string
}

type EventHandle struct {
	eth             *ethclient.Client
	mintSigHash     common.Hash
	transferSigHash common.Hash
}

func NewEventHandle() (error, *EventHandle) {
	eth, err := ethclient.Dial(nodeAddr)
	if err != nil {
		return err, nil
	}
	mintSigHash := crypto.Keccak256Hash([]byte("MintNFT(uint _token)"))
	transferSigHash := crypto.Keccak256Hash([]byte("TransferNFT(address _operator, address _from, address _to, uint _token, uint _amt, string _ext)"))
	ret := &EventHandle{
		eth:             eth,
		mintSigHash:     mintSigHash,
		transferSigHash: transferSigHash,
	}
	return nil, ret
}

func (e *EventHandle) Do(m chan MintEvent, t chan TransferEvent) {
	go func() {
		//step1 : init
		logs := make(chan types.Log)
		err, sub, contractAbi := e.init(logs)
		if err != nil {
			return
		}

		//step2: consume event
		for {
			select {
			case err := <-sub.Err():
				log.Printf("consume fail :%v\n", err)
			case vLog := <-logs:
				e.consume(contractAbi, &vLog, m, t)
			}
		}
	}()
}

func (e *EventHandle) init(logs chan types.Log) (error, ethereum.Subscription, *abi.ABI) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(dAppContractAddr)},
	}

	sub, err := e.eth.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Printf("subscribe fail :%v\n", err)
		return err, nil, nil
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(EthereumMetaData.ABI)))
	if err != nil {
		log.Printf("abi fail :%v\n", err)
		return err, nil, nil
	}
	return nil, sub, &contractAbi
}

func (e *EventHandle) consume(contractAbi *abi.ABI, event *types.Log, m chan MintEvent, t chan TransferEvent) {
	switch event.Topics[0].Hex() {
	case e.mintSigHash.Hex():
		var ev MintEvent
		err := contractAbi.UnpackIntoInterface(&ev, "MintNFT", event.Data)
		if err != nil {
			log.Printf("Unpack MintNFT event fail:%v\n", err)
		}
		m <- ev
	case e.transferSigHash.Hex():
		var ev TransferEvent
		err := contractAbi.UnpackIntoInterface(&ev, "TransferNFT", event.Data)
		if err != nil {
			log.Printf("Unpack MintNFT event fail:%v\n", err)
		}
		t <- ev
	default:
		log.Printf("not use event:%v\n", *event)
	}
}
