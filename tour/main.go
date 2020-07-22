package main

import (
	"go-project-example/tour/cmd"
	"log"
)

/**
 *@Author tudou
 *@Date 2020/7/22
 **/


func main(){
	err:=cmd.Execute()
	if err!=nil{
		log.Fatalf("cmd.Execute err: %v",err)
	}
}