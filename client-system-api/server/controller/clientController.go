package controller

import (
	"log"
	"net/http"

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
	err := c.dbCtx.MakeMigrations()
	if err != nil {
		log.Println(err)
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
	r, err := http.Get("http://person-system-api:8080/api/person/" + client.PersonId.String())
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		log.Println("Person not found")
		ctx.JSON(400, gin.H{"error": "Person not found"})
		return
	}
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
	ctx.JSON(200, client)
}

func (c *ClientController) GetClients(ctx *gin.Context) {
	res, err := c.dbCtx.GetClients()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}
