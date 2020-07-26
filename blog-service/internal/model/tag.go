package model

/**
 *@Author tudou
 *@Date 2020/7/26
 **/



type Tag struct{
	*Model
	Name string `json:"name"`
	State string `json:"state"`
}


func (a *Tag) TableName() string{
	return "blog_tag"
}

