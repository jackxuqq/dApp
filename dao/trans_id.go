package dao

import (
	"fmt"

	"github.com/jackxuqq/dApp/model"
	"gorm.io/gorm"
)

type TransIDStore interface {
	Generate() (error, int64)
}

type TransIDMysql struct {
	db *gorm.DB
}

func NewTransIDMysql() (error, TransIDStore) {
	db, err := gorm.Open("mysql", "root:passwd@tcp(127.0.0.1:3306)/nft?charset=utf8")
	if err != nil {
		fmt.Printf("init mysql fail[%v]\n", err)
		return err, nil
	}
	db.AutoMigrate(&model.TransID{})
	ret := &TransIDMysql{
		db: db,
	}
	return nil, ret
}

func (t *TransIDMysl) Generate() (error, int64) {
	t := model.TransID{}
	result := t.db.Create(t)
	if result.Error != nil {
		return result.Error, 0
	}
	return nil, t.ID
}
