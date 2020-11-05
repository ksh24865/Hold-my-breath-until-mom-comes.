package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, data gin.H, templateName string) {

	c.HTML(http.StatusOK, templateName, data)

}

func arender(c *gin.Context, data gin.H, templateName string) {
	for {
		c.HTML(http.StatusOK, templateName, data)
		time.Sleep(time.Second * 20)
		println("sleep")
	}

}
