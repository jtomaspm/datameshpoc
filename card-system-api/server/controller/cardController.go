package controller

import (
	"log"
	"net/http"

	"datamesh.poc/card-system-api/dal/context"
	"datamesh.poc/card-system-api/dal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CardController struct {
	router *gin.Engine
	dbCtx  *context.DbContext
}

func New(router *gin.Engine) *CardController {
	c := &CardController{
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

func (c *CardController) Setup() {
	c.router.POST("/api/card", c.CreateCard)
	c.router.GET("/api/card/:id", c.GetCard)
	c.router.GET("/api/card", c.GetCards)
}

func (c *CardController) CreateCard(ctx *gin.Context) {
	var card model.Card
	err := ctx.BindJSON(&card)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Println(card)
	r, err := http.Get("http://client-system-api:8080/api/client/" + card.ClientId.String())
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		log.Println("Client not found")
		ctx.JSON(400, gin.H{"error": "Client not found"})
		return
	}
	id, err := c.dbCtx.CreateCard(card)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"id": id})
}

func (c *CardController) GetCard(ctx *gin.Context) {
	idu, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	card, err := c.dbCtx.GetCard(idu)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, card)
}

func (c *CardController) GetCards(ctx *gin.Context) {
	res, err := c.dbCtx.GetCards()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}
