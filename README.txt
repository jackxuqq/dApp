核心模块:
1)智能合约
	基于Openzeppelin库编写合约，用remid web ide 发布到以太坊获得合约地址
2)后台服务
	--存储: Mysql5.8
	a) address -- 复用用户uid
	b) trans_id -- 交易表自增ID
	c) token -- NFT表自增ID 
	d) trans_status -- 交易流水表枚举字段
	--Web框架: gin

3)后台服务调用智能合约func
	基于go-ethereum库对以太坊node发起RPC调用

4)后台服务监听智能event
	基于go-ehereum库编写一个独立的协程，订阅合约事件=>解析=>更新铸造、转账状态
