package model

type TransSeq struct {
	ID         int64  `gorm:"column:id;primaryKey;type:int;auto_increament;"`
	From       int64  `gorm:"column:from;type:int"`
	To         int64  `gorm:"column:to;type:int"`
	Amount     string `gorm:"column:amount;type:int"`
	OccurtTs   int64  `gorm:"column:occurt_ts;type:int"`
	CompleteTs int64  `gorm:"column:complete_ts;type:int"`
}
