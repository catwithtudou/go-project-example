package main

import (
	"flag"
	"go-project-example/blog-service/global"
	"go-project-example/blog-service/internal/model"
	"go-project-example/blog-service/internal/routers"
	"go-project-example/blog-service/pkg/setting"
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
	err:=setupFlag()
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main(){
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

func setDBEngine()(err error){
	global.DBEngine,err =model.NewDBEngine(global.DatabaseSetting)
	return
}