package main

import (
	"github.com/baby-vibe/dataservice/mysql"
	"github.com/baby-vibe/rest/handler"
	"github.com/baby-vibe/usecase/manageSignal"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysql.Setup()
	sr := mysql.NewSignalRepo()
	msuc := manageSignal.NewManageSignalUsecase(sr)
	h := handler.NewGinHandler(msuc)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.LoadHTMLGlob("view/*")
	r.Static("/vibe", "./vibes")
	r.GET("/", h.ShowIndexPage)

	r.Run(":8080")
}
