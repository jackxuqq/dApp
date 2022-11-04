package model

type NtfStatus int

const (
	NSBuilding = iota + 1
	NSBuildDone
	NDBuildFail
)

const (
	AttributeTitle  = "title"
	AttributeImage  = "image"
	AttributeAmount = "amount"
)

type Ntf struct {
	Token  int64  `gorm:"column:id;primaryKey;type:int;autoIncrement;"`
	Title  string `gorm:"column:title;type:varchar(30);"`
	Image  string `gorm:"column:image;type:varchar(100);"`
	Amount int64  `gorm:"column:amount;type:int;"`
	Status int    `gorm:"column:status;type:int;"`
}

func (n *Ntf) TableName() string {
	return "t_ntf"
}
