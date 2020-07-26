package main

import (
	"go-project-example/blog-service/internal/routers"
	"net/http"
	"time"
)

/**
 *@Author tudou
 *@Date 2020/7/26
 **/


func main(){
	router:=routers.NewRouter()
	s:=&http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	_ = s.ListenAndServe()
}