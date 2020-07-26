package model

/**
 *@Author tudou
 *@Date 2020/7/26
 **/


type ArticleTag struct{
	*Model
	TagID uint32  `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func ( a ArticleTag) TableName()string{
	return "blog_article_tag"
}

