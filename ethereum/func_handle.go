package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

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
	err, auth := n.generateAuth()
	if err != nil {
		return err
	}
	_, err = contract.Mint(auth, common.HexToAddress(addr), Int64ToBig(token), Int64ToBig(amount), strAttr)
	if err != nil {
		log.Printf("Mint fail:%v\n", err)
		return err
	}
	return nil
}

func (n *NodeHandle) Transfer(from string, to string, token int64, amount int64, ext string) error {
	contract, err := NewEthereum(common.HexToAddress(dAppContractAddr), n.eth)
	if err != nil {
		return err
	}

	err, auth := n.generateAuth()
	if err != nil {
		return err
	}
	_, err = contract.Transfer(auth, common.HexToAddress(from), common.HexToAddress(to), Int64ToBig(token), Int64ToBig(amount), ext)
	if err != nil {
		return err
	}
	return nil
}

func (n *NodeHandle) generateAuth() (error, *bind.TransactOpts) {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return err, nil
	}
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("publicKeyECDSA fail"), nil
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	gasPrice, err := n.eth.SuggestGasPrice(context.Background())
	if err != nil {
		return err, nil
	}
	nonce, err := n.eth.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err, nil
	}
	auth := bind.NewKeyedTransactor(pk)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	return nil, auth
}
