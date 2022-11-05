package dao

import (
	"fmt"

	"github.com/jackxuqq/dApp/model"
	"github.com/jinzhu/gorm"
)

type TransIDStore interface {
	Generate() (error, int64)
}

type TransIDMysql struct {
	db *gorm.DB
}

func NewTransIDMysql() (error, TransIDStore) {
	db, err := gorm.Open("mysql", "root:root@123@tcp(127.0.0.1:3306)/nft?charset=utf8")
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

func (t *TransIDMysql) Generate() (error, int64) {
	tr := model.TransID{}
	result := t.db.Create(&tr)
	if result.Error != nil {
		return result.Error, 0
	}
	return nil, tr.ID
}
