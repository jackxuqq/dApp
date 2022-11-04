package logic

import (
	"strconv"
)

// Transfer 转账
// from 发起转账地址
// to 接收方地址
// token NFT 唯一ID
// amount 转账金额
func (d *DAppLogic) Transfer(from string, to string, token int64, amount int64) error {
	//step1: generate transID in db
	err, tID := d.transID.Generate()
	if err != nil {
		return err
	}

	//step2: rpc ethereum transfer func
	err = d.nodeHandle.Transfer(from, to, token, amount, strconv.FormatInt(tID, 10))
	if err != nil {
		return err
	}

	//step3: create trans seq in db
	err = d.transSeq.Create(tID, from, to, token, amount)
	if err != nil {
		return err
	}
	return nil
}
