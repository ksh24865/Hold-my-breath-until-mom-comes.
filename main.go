package main

import (
	"github.com/baby-vibe/dataservice/mysql"
	"github.com/baby-vibe/rest/handler"
	"github.com/baby-vibe/usecase/manageSignal"
	"github.com/baby-vibe/usecase/manageWeight"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	println("-1")
	mysql.Setup()
	println("-2")
	sr := mysql.NewSignalRepo()
	println("0")
	wr := mysql.NewWeightRepo()
	println("1")
	msuc := manageSignal.NewManageSignalUsecase(sr)
	println("2")
	mwuc := manageWeight.NewManageWeightUsecase(wr)
	println("3")
	h := handler.NewGinHandler(msuc, mwuc)
	println("4")

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
