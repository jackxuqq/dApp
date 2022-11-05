package model

type TransID struct {
	ID int64 `gorm:"column:id;primaryKey;type:int auto_increment;"`
}

func (t *TransID) TableName() string {
	return "t_trans_id"
}
