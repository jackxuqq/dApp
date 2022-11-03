package model

type TransStatus int

const (
	TSTransfering = iota + 1
	TSTransferDone
	TSTransferFail
)

type TransSeq struct {
	ID         int64  `gorm:"column:id;primaryKey;type:int;"`
	From       int64  `gorm:"column:from;type:int"`
	To         int64  `gorm:"column:to;type:int"`
	Amount     int64 `gorm:"column:amount;type:int"`
	OccurtTs   int64  `gorm:"column:occurt_ts;type:int"`
	CompleteTs int64  `gorm:"column:complete_ts;type:int"`
	Status     int    `gorm:"column:status;type:int;"`
}
