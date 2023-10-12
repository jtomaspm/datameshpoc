package controller

import (
	"log"

	"datamesh.poc/client-system-api/dal/context"
	"datamesh.poc/client-system-api/dal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ClientController struct {
	router *gin.Engine
	dbCtx  *context.DbContext
}

func New(router *gin.Engine) *ClientController {
	c := &ClientController{
		router: router,
		dbCtx:  context.New(),
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
	var client model.Client
	err := ctx.BindJSON(&client)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Println(client)
	id, err := c.dbCtx.CreateClient(client)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"id": id})
}

func (c *ClientController) GetClient(ctx *gin.Context) {
	idu, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	client, err := c.dbCtx.GetClient(idu)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"client": client})
}

func (c *ClientController) GetClients(ctx *gin.Context) {

}
