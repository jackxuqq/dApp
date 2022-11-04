package ethereum

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type NodeHandle struct {
	eth *ethclient.Client
}

func NewNodeHandle() (error, *NodeHandle) {
	eth, err := ethclient.Dial(nodeAddr)
	if err != nil {
		return err, nil
	}
	ret := &NodeHandle{
		eth: eth,
	}
	return nil, ret
}

func (n *NodeHandle) Mint(addr string, token int64, amount int64, attr map[string]string) error {
	contract, err := NewEthereum(common.HexToAddress(dAppContractAddr), n.eth)
	if err != nil {
		return err
	}
	strAttr := "{"
	for k, v := range attr {
		strAttr += fmt.Sprintf("\"%s\":\"%s\"", k, v)
	}
	strAttr += "}"

	transOpt, _ := bind.NewTransactorWithChainID(strings.NewReader(privateKey), passWord, Int64ToBig(chanID))
	_, err = contract.Mint(transOpt, common.HexToAddress(addr), Int64ToBig(token), Int64ToBig(amount), strAttr)
	if err != nil {
		return err
	}
	return nil
}

func (n *NodeHandle) Transfer(from string, to string, token int64, amount int64, ext string) error {
	contract, err := NewEthereum(common.HexToAddress(dAppContractAddr), n.eth)
	if err != nil {
		return err
	}
	transOpt, _ := bind.NewTransactorWithChainID(strings.NewReader(privateKey), passWord, Int64ToBig(chanID))
	_, err = contract.Transfer(transOpt, common.HexToAddress(from), common.HexToAddress(to), Int64ToBig(token), Int64ToBig(amount), ext)
	if err != nil {
		return err
	}
	return nil
}
