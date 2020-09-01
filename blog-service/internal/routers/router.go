package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-project-example/blog-service/docs"
	"go-project-example/blog-service/global"
	"go-project-example/blog-service/internal/middleware"
	"go-project-example/blog-service/internal/routers/api"
	v1 "go-project-example/blog-service/internal/routers/api/v1"
	"net/http"
)

/**
 *@Author tudou
 *@Date 2020/7/26
 **/

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())

	upload := api.NewUpload()
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.POST("/upload/file", upload.UploadFile)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()
	apiV1 := r.Group("/api/v1")
	{
		// 创建标签
		apiV1.POST("/tags", tag.Create)
		// 删除指定标签
		apiV1.DELETE("/tags/:id", tag.Delete)
		// 更新指定标签
		apiV1.PUT("/tags/:id", tag.Update)
		// 获取标签列表
		apiV1.GET("/tags", tag.List)

		// 创建文章
		apiV1.POST("/articles", article.Create)
		// 删除指定文章
		apiV1.DELETE("/articles/:id", article.Delete)
		// 更新指定文章
		apiV1.PUT("/articles/:id", article.Update)
		// 获取指定文章
		apiV1.GET("/articles/:id", article.Get)
		// 获取文章列表
		apiV1.GET("/articles", article.List)
	}

	return r
}
