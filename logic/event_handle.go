package logic

import (
	"strconv"

	"github.com/jackxu/dApp/ethereum"
	"github.com/jackxu/dApp/model"
)

func (d *DAppLogic) HandleEvent() {
	m := make(chan ethereum.MintEvent, 10)
	t := make(chan ethereum.TransferEvent, 10)
	d.eventHandle.Do(m, t)

	go func(m chan ethereum.MintEvent) {
		for {
			e := <-m
			_ = d.ntf.UpdateStatus(e.Token.Int64(), model.NSBuildDone)
		}
	}(m)

	go func(t chan ethereum.TransferEvent) {
		for {
			e := <-t
			transID, _ := strconv.ParseInt(e.Ext, 10, 64)
			_ = d.transSeq.UpdateStatus(transID, model.TSTransferDone)
		}
	}(t)
}
