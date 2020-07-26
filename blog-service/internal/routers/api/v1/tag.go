package v1

import "github.com/gin-gonic/gin"

/**
 *@Author tudou
 *@Date 2020/7/26
 **/

type Tag struct{

}

func NewTag() Tag{
	return Tag{}
}


func(t Tag)Get(c *gin.Context){}
func(t Tag)List(c *gin.Context){}
func(t Tag)Create(c *gin.Context){}
func(t Tag)Update(c *gin.Context){}
func(t Tag)Delete(c *gin.Context){}
