package logic

// Transfer 转账
// from 发起转账uid
// to 接收方uid
// token NFT 唯一ID
// amount 转账金额
func (d *DAppLogic) Transfer(from int64, to int64, token int64, amount int64) error {
	//step1: generate transID in db

	//step2: rpc ethereum transfer func

	//step3: create trans req in db
}
