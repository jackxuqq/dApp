package logic

import (
	"github.com/jackxu/dApp/dao"
	"github.com/jackxu/dApp/ethereum"
)

type DAppLogic struct {

	//func Mint(uid int64, title string, image string, amount int64)(error, int64);

	//func (d *DAppLogic) Transfer(from int64, to int64, token int64, amount int64) error;

	//dao instance
	ntf      NtfStore
	transID  TransIDStore
	transSeq TransSeqStore

	//ethereum rpc instance
	nodeHandle *NodeHandle
}

func NewDAppLogic() (error, *DAppLogic) {

	ntf, err := dao.NewNtfMysql()
	if err != nil {
		return err, nil
	}

	transID, err := dao.NewTransIDMysql()
	if err != nil {
		return err, nil
	}

	transSeq, err := dao.NewTransSeqMysql()
	if err != nil {
		return err, nil
	}

	nodeHandle, err := ethereum.NewNodeHandle()
	if err != nil {
		return err, nil
	}

	return &DAppLogic{
		ntf:        ntf,
		transID:    transID,
		transSeq:   transSeq,
		nodeHandle: nodeHandle,
	}
}
