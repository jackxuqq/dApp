package logic

import (
	"github.com/jackxu/dApp/dao"
	"github.com/jackxu/dApp/ethereum"
)

type DAppLogic struct {

	//func Mint(uid int64, title string, image string, amount int64)(error, int64);

	//func (d *DAppLogic) Transfer(from int64, to int64, token int64, amount int64) error;

	//dao instance
	ntf      dao.NtfStore
	transID  dao.TransIDStore
	transSeq dao.TransSeqStore

	//ethereum rpc instance
	nodeHandle *ethereum.NodeHandle
}

func NewDAppLogic() (error, *DAppLogic) {

	err,ntf := dao.NewNtfMysql()
	if err != nil {
		return err, nil
	}

	err,transID := dao.NewTransIDMysql()
	if err != nil {
		return err, nil
	}

	err,transSeq := dao.NewTransSeqMysql()
	if err != nil {
		return err, nil
	}

	err, nodeHandle := ethereum.NewNodeHandle()
	if err != nil {
		return err, nil
	}

	return nil, &DAppLogic{
		ntf:        ntf,
		transID:    transID,
		transSeq:   transSeq,
		nodeHandle: nodeHandle,
	}
}
