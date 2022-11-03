package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type NodeHandle struct {
	eth *ethclient.Client
}

func NewNodeHandle() (error, *NodeHandle) {
	eth, err := ethclient.Dial("http://xxx.xx.com:8545")
	if err != nil {
		return err, nil
	}
	ret := &NodeHandle{
		eth: eth,
	}
	return nil,ret
}

func (n *NodeHandle) Mint(uid int64, token int64, amount int64, attr map[string]string) error {
	return nil
	/*
	contract, err := NewDApp1155(common.HexToAddress("replace 2 contract addr"))
	if err != nil {
		return err
	}
	strAttr := "{"
	for k, v := range attr {
		strAttr += fmt.Sprintf("\"%s\":\"%s\"", k, v)
	}
	strAttr += "}"
	_, err = contract.Mint(uid, token, amount, strAttr)
	if err != nil {
		return err
	}
	return nil
	 */
}

func (n *NodeHandle) Transfer(from int64, to int64, token int64, amount int64, ext string) error {
	return nil
	/*
	contract, err := NewDApp1155(common.HexToAddress("replace 2 contract addr"))
	if err != nil {
		return err
	}
	err = contract.Transfer(from, to, token, amount, ext)
	if err != nil {
		return err
	}
	return nil
	 */
}
