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

	r.LoadHTMLGlob("view/*.html")
	r.Static("/vibe", "./vibes")

	r.GET("/", h.ShowIndexPage)

	r.GET("/humid", h.Runhumid)
	r.GET("/blanket", h.Runblanket)
	r.GET("/lamp", h.Runlamp)
	r.GET("/graph", h.ShowGraphPopup)
	r.GET("/test1", h.RunTest1)
	r.GET("/test2", h.RunTest2)
	r.Run(":8080")
}
