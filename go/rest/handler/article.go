package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *GinHandler) signalPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	if signal, err := h.msuc.GetSignalByID(1); err == nil {
		// Call the render function with the title, article and the name of the
		// template
		println("check ---- signal")
		render(c, gin.H{
			"title":   "signal Page",
			"payload": signal}, "signal.html")
	} else {
		// If the article is not found, abort with an error
		println("check ---- NO signal")
		c.AbortWithError(http.StatusNotFound, err)
	}
}

func (h *GinHandler) ShowIndexPage(c *gin.Context) {
	// Check if the article ID is valid
	if signal, err := h.msuc.GetSignalByID(1); err == nil {
		// Call the render function with the title, article and the name of the
		// template
		println("check ---- signal")
		render(c, gin.H{
			"title":   "Home Page",
			"payload": signal}, "index.html")
	} else {
		// If the article is not found, abort with an error
		println("check ---- NO signal")
		c.AbortWithError(http.StatusNotFound, err)
	}
}
