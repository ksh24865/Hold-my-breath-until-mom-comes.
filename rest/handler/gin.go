package handler

import (
	"log"

	"github.com/baby-vibe/usecase"
)

func catchPanic() {
	if p := recover(); p != nil {
		log.Printf("%+v\n", p)
	}
}

type GinHandler struct {
	msuc     usecase.ManageSignalUsecase
	signalID int
	Humid    int
	Blanket  int
	Lamp     int
}

func NewGinHandler(msuc usecase.ManageSignalUsecase) *GinHandler {
	RunPythonScript("py/baby_care.py")
	return &GinHandler{
		msuc:    msuc,
		Humid:   0,
		Blanket: 0,
		Lamp:    0,
	}
}

/*
func (h *GinHandler) ShowArticleCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func (h *GinHandler) ShowLoginPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func (h *GinHandler) ShowRegistrationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Register"}, "register.html")
}
*/
