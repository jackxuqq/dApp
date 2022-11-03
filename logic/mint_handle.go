package logic

// Mint 铸造NFT
// uid  用户ID
// title title
// image  预览图
// amount 铸造数量
// return nft token
func (d *DAppLogic) Mint(uid int64, title string, image string, amount int64) (error, int64) {
	//step1: create nft record in db

	//step2: rpc ethereum mint func

}
