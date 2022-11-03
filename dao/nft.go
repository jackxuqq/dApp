package dao

import (
	"fmt"
	"strconv"

	"github.com/jackxuqq/dApp/model"
	"gorm.io/gorm"
)

type NtfStore interface {

	// Create 持久化ntf对象到中心化存储, 返回token
	Create(attributes map[string]string) (error, int64)

	// UpdateStaus 更新ntf对象的铸造状态
	UpdateStatus(token int64, status model.NtfStatus) error
}

type NtfMysql struct {
	db *gorm.DB
}

func NewNtfMysql() (error, NtfStore) {
	db, err := gorm.Open("mysql", "root:passwd@tcp(127.0.0.1:3306)/nft?charset=utf8")
	if err != nil {
		fmt.Printf("init mysql fail[%v]\n", err)
		return err, nil
	}
	db.AutoMigrate(&model.Ntf{})
	ret := &NtfMysql{
		db: db,
	}
	return nil, ret
}

func (n *NtfMysql) Create(attributes map[string]string) (error, int64) {
	m := model.Ntf{}
	m.Title = attributes[model.AttributeTitle]
	m.Image = attributes[model.AttributeImage]
	m.Amount, _ = strconv.Atoi(attributes[model.AttributeAmount])
	m.Status = model.NSBulding
	result := n.db.Create(m)
	if result.Error != nil {
		return result.Error, 0
	}
	return nil, m.Token
}

func (n *NtfMysql) UpdateStatus(token int64, status NtfStatus) error {
	result := n.db.Where("token=?", token).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
