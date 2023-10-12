package controller

import (
	"github.com/gin-gonic/gin"
)

type ClientController struct {
	router *gin.Engine
}

func New(router *gin.Engine) *ClientController {
	c := &ClientController{
		router: router,
	}
	c.Setup()
	return c
}

func (c *ClientController) Setup() {
	c.router.POST("/api/client", c.CreateClient)
	c.router.GET("/api/client/:id", c.GetClient)
	c.router.GET("/api/client", c.GetClients)
}

func (c *ClientController) CreateClient(ctx *gin.Context) {
}

func (c *ClientController) GetClient(ctx *gin.Context) {
}

func (c *ClientController) GetClients(ctx *gin.Context) {
}
