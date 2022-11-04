package ethereum

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func Int64ToAddress(v int64) common.Address {
	b := big.Int{}
	b.SetInt64(v)
	return common.BigToAddress(&b)
}

func Int64ToBig(v int64) *big.Int {
	b := big.Int{}
	b.SetInt64(v)
	return &b
}
