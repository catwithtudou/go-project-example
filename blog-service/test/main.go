package main

import (
	"context"
	"fmt"
)

/**
 *@Author tudou
 *@Date 2020/9/1
 **/
//
//func main(){
//	re:=make(map[string]string)
//	re["1"]="2"
//	fun(re)
//	fmt.Println(re["1"])
//}

func fun(m map[string]string) {
	m["1"] = "2"
}

func f(ctx context.Context) {
	context.WithValue(ctx, "foo", -6)
}

func main() {
	ctx := context.TODO()
	f(ctx)
	fmt.Println(ctx.Value("foo"))
}
