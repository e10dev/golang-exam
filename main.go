package main

import (
	"fmt"

	"e10dev.example/exam01/service/lib"
	"e10dev.example/exam01/service/webserver"
	"github.com/gin-gonic/gin"
)

func initMain() *gin.Engine {
	lib.Initialization()
	fmt.Println("init done")
	r := webserver.InitRouter()

	return r
}

func main() {
	r := initMain()

	r.Run()
}
