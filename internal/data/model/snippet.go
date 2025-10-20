package model

type Snippet struct {
	ID    uint   `gorm:"primarykey" json:"ID"`
	Key   string `gorm:"column:key;not null;uniqueIndex:key_value" json:"key"`
	Value string `gorm:"column:value;not null;uniqueIndex:key_value" json:"value"`
}

func (Snippet) TableName() string {
	return "leo"
}
