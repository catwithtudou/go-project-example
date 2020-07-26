package model

/**
 *@Author tudou
 *@Date 2020/7/26
 **/

const (
	STATE_CLOSE = iota
	STATE_OPEN
)


type Model struct{
	ID uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"` //创建者
	ModifiedBy string `json:"modified_by"` //修改者
	CreatedOn  uint32 `json:"created_on"` //创建时间
	ModifiedOn uint32 `json:"modified_on"` //修改时间
	DeletedOn  uint32 `json:"deleted_on"` //删除时间
	IsDel      uint8  `json:"is_del"` //删除状态：0：未删除 1：已删除
}



