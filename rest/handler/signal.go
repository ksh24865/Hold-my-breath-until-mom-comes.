package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
func (h *GinHandler) ShowGraphPopup(c *gin.Context) {
	// Check if the article ID is valid
	render(c, gin.H{
		"title": "Home Page"}, "graph.html")
}
func (h *GinHandler) RunTest1(c *gin.Context) {
	if h.Humid == 0 {
		h.Humid = 1
		RunPythonScript("py/test1.py")
	} else {
		h.Humid = 0
		RunPythonScript("py/test2.py")
	}
}
func (h *GinHandler) RunTest2(c *gin.Context) {
	RunPythonScript("py/test2.py")
	h.ShowIndexPage(c)
}
func (h *GinHandler) Runhumid(c *gin.Context) {
	if h.Humid == 0 {
		h.Humid = 1
		RunPythonScript("py/humid_on.py")
	} else {
		h.Humid = 0
		RunPythonScript("py/humid_off.py")
	}
}
func (h *GinHandler) Runblanket(c *gin.Context) {
	if h.Humid == 0 {
		h.Blanket = 1
		RunPythonScript("py/blanket_on.py")
	} else {
		h.Blanket = 0
		RunPythonScript("py/blanket_off.py")
	}
}
func (h *GinHandler) Runlamp(c *gin.Context) {
	if h.Humid == 0 {
		h.Lamp = 1
		RunPythonScript("py/lamp_on.py")
	} else {
		h.Lamp = 0
		RunPythonScript("py/lamp_off.py")
	}
}
