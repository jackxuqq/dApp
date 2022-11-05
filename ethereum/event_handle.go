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
	Operator common.Address
	From     common.Address
	To       common.Address
	Token    *big.Int
	Amt      *big.Int
	Ext      string
}

type EventHandle struct {
	eth             *ethclient.Client
	mintSigHash     common.Hash
	transferSigHash common.Hash
}

func NewEventHandle() (error, *EventHandle) {
	eth, err := ethclient.Dial(nodeEventAddr)
	if err != nil {
		return err, nil
	}
	mintSigHash := crypto.Keccak256Hash([]byte("MintNFT(uint)"))
	transferSigHash := crypto.Keccak256Hash([]byte("TransferNFT(address,address,address,uint,uint,string)"))
	log.Printf("mintSigHash[%s]\n", mintSigHash.Hex())
	log.Printf("transferSigHash[%s]\n", transferSigHash.Hex())
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
	//case e.mintSigHash.Hex():
	case "0x2cc161d398589ff5526cd6b429338ae827d66c48171d3dbaeac9a7bf822935bd":
		log.Printf("got a mint event\n")
		var ev MintEvent
		err := contractAbi.UnpackIntoInterface(&ev, "MintNFT", event.Data)
		if err != nil {
			log.Printf("Unpack MintNFT event fail:%v\n", err)
		}
		m <- ev
	//case e.transferSigHash.Hex():
	case "0x6a2f2e877bc0a320cc6f53954a92731fe438598b0579b395c58188785ff1d460":
		log.Printf("got a transfer event\n")
		var ev TransferEvent
		err := contractAbi.UnpackIntoInterface(&ev, "TransferNFT", event.Data)
		if err != nil {
			log.Printf("Unpack MintNFT event fail:%v\n", err)
		}
		t <- ev
	default:
		log.Printf("not use event:%+v\n", *event)
	}
}
