package api

import (
	"github.com/gin-gonic/gin"
	"go-project-example/blog-service/global"
	"go-project-example/blog-service/internal/service"
	"go-project-example/blog-service/pkg/app"
	"go-project-example/blog-service/pkg/convert"
	"go-project-example/blog-service/pkg/errcode"
	"go-project-example/blog-service/pkg/upload"
)

/**
 *@Author tudou
 *@Date 2020/9/1
 **/

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

type UploadResp struct {
	FileAccessUrl string `json:"file_access_url"`
}

// @Summary 上传文件
// @Produce json
// @Param file body file true "文件"
// @Param type body string true "文件类型“
// @Success 200 {object} UploadResp "成功”
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /upload/file [post]
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
	return
}
