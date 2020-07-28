package service

import (
	"context"
	otgorm "github.com/eddycjy/opentracing-gorm"
	"go-project-example/blog-service/global"
	"go-project-example/blog-service/internal/dao"
)

/**
 *@Author tudou
 *@Date 2020/7/28
 **/


type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
