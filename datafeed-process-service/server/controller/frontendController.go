package controller

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

type CardController struct {
	router *gin.Engine
}

func New(router *gin.Engine) *CardController {
	c := &CardController{
		router: router,
	}
	c.Setup()
	return c
}

func (c *CardController) Setup() {
	c.router.GET("/", c.HomePage)
	c.router.POST("/clicked", c.Clicked)
}

func (c *CardController) Clicked(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "clicked"})
}

func (c *CardController) HomePage(ctx *gin.Context) {
	attr := map[string]interface{}{
		"title": "Data Feed Service",
	}
	t, err := template.ParseFiles("./server/resource/index.html")
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Content-Type", "text/html")
	t.Execute(ctx.Writer, attr)
}
