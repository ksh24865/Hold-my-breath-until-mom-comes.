package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, data gin.H, templateName string) {

	c.HTML(http.StatusOK, templateName, data)

}

