package dao

import "github.com/jinzhu/gorm"

/**
 *@Author tudou
 *@Date 2020/7/28
 **/


type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}