package webserver

import (
	"e10dev.example/exam01/service/webserver/api"
	"e10dev.example/exam01/service/webserver/api/account"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	a := r.Group("/api")
	{
		a.POST("/login", api.Login)
	}

	b := r.Group("/api/account")
	{
		b.GET("", account.GetAllAccount)
		b.GET("/:seq", account.GetAccount)
		b.POST("", account.NewAccount)
		b.PUT("/:seq", account.UpdateAccount)
		b.DELETE("/:seq", account.DeleteAccount)
		b.GET("/download", account.DownloadCSV)
	}

	return r
}
