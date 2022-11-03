package dao

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/jackxuqq/dApp/model"
)

type TransSeqStore interface {
	Create(transID int64, from int64, to int64, token int64, amount int64) error
	UpdateStatus(transID int64, status model.TransStatus) error
}

type TransSeqMysql struct {
	db *gorm.DB
}

func NewTransSeqMysql(error, TransSeqStore) {
	db, err := gorm.Open("mysql", "root:passwd@tcp(127.0.0.1:3306)/nft?charset=utf8")
	if err != nil {
		fmt.Printf("init mysql fail[%v]\n", err)
		return err, nil
	}
	db.AutoMigrate(&model.TransSeq{})
	ret := &TransSeqMysql{
		db: db,
	}
	return nil, ret
}

func (t *TransSeqMysql) Create(transID int64, from int64, to int64, token int64, amount int64) error {
	t := model.TransSeq{}
	t.ID = transID
	t.From = from
	t.To = to
	t.Amount = amount
	t.OccurtTs = time.Now().UnixTimestamp()
	t.Status = model.TSTransfering
	result := t.db.Create(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *TransSeqMysql) UpdateStatus(transID int64, status model.TransStatus) error {
	result := n.db.Where("id=?", transID).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
