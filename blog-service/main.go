package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"go-project-example/blog-service/global"
	"go-project-example/blog-service/internal/model"
	"go-project-example/blog-service/internal/routers"
	"go-project-example/blog-service/pkg/logger"
	"go-project-example/blog-service/pkg/setting"

	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"strings"
	"time"
)

/**
 *@Author tudou
 *@Date 2020/7/26
 **/
func init(){
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}


// @title 博客系统
// @version 1.0
// @description Chapter2
func main(){
	gin.SetMode(global.ServerSetting.RunMode)
	router:=routers.NewRouter()
	s:=&http.Server{
		Addr:              ":"+global.ServerSetting.HttpPort,
		Handler:           router,
		ReadTimeout:       global.ServerSetting.ReadTimeout,
		WriteTimeout:      global.ServerSetting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	_ = s.ListenAndServe()
}

var (
	port      string
	runMode   string
	config    string
	isVersion bool
)

func setupFlag() error {
	flag.StringVar(&port, "port", "", "start port")
	flag.StringVar(&runMode, "mode", "", "start pattern")
	flag.StringVar(&config, "config", "configs/", "setting path")
	flag.BoolVar(&isVersion, "version", false, "information")
	flag.Parse()

	return nil
}

func setupSetting()error{
	s,err:=setting.NewSetting(strings.Split(config,",")...)
	if err!=nil{
		return err
	}
	err=s.ReadSection("Server",&global.ServerSetting)
	if err!=nil{
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	return nil
}

func setupDBEngine()(err error){
	global.DBEngine,err =model.NewDBEngine(global.DatabaseSetting)
	return
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}