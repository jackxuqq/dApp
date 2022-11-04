package logic

import (
	"strconv"

	"github.com/jackxu/dApp/model"
)

// Mint 铸造NFT
// uid  用户ID
// title 标题
// image  预览图
// amount 铸造数量
// return nft token
func (d *DAppLogic) Mint(uid int64, title string, image string, amount int64) (error, int64) {
	//step1: create nft record in db
	attr := make(map[string]string)
	attr[model.AttributeTitle] = title
	attr[model.AttributeImage] = image
	attr[model.AttributeAmount] = strconv.FormatInt(amount, 10)
	err, token := d.ntf.Create(attr)
	if err != nil {
		return err, 0
	}

	//step2: rpc ethereum mint func
	err = d.nodeHandle.Mint(uid, token, amount, attr)
	if err != nil {
		return err, 0
	}
	return nil, token
}
