package model

type Ntf struct {
	Token  int64  `gorm:"column:token;primaryKey;type:int;auto_increament;"`
	Title  string `gorm:"column:title;type:varchar(30);"`
	Image  string `gorm:"column:image;type:varchar(100);"`
	Amount string `gorm:"column:amount;type:int;"`
}
