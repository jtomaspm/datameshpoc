package controller

import (
	"encoding/json"
	"html/template"
	"strconv"

	"datamesh.poc/datafeed-process-service/feeder"
	"github.com/gin-gonic/gin"
)

type CardController struct {
	router       *gin.Engine
	clientFeeder *feeder.ClientFeeder
}

func New(router *gin.Engine) *CardController {
	c := &CardController{
		router:       router,
		clientFeeder: feeder.NewClientFeeder("http://client-process-api:8080/api/client"),
	}
	c.Setup()
	return c
}

func (c *CardController) Setup() {
	c.router.GET("/", c.HomePage)
	c.router.GET("/all", c.AllHtml)
	c.router.GET("/clients", c.ClientsHtml)
	c.router.GET("/cards", c.CardsHtml)
	c.router.GET("/transactions", c.TransactionsHtml)
	c.router.POST("/api/feedClients/:amount", c.FeedClients)
	c.router.GET("/api/clients", c.GetClients)
}

func (c *CardController) GetClients(ctx *gin.Context) {
	res, err := json.Marshal(c.clientFeeder.GetClients())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"clients": string(res)})
}

func (c *CardController) FeedClients(ctx *gin.Context) {
	amount, err := strconv.Atoi(ctx.Param("amount"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	go c.clientFeeder.Feed(amount)
	ctx.JSON(200, gin.H{"message": "Feeding clients"})
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

func (c *CardController) AllHtml(ctx *gin.Context) {
	attr := map[string]interface{}{}
	t, err := template.ParseFiles("./server/resource/all.html")
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Content-Type", "text/html")
	t.Execute(ctx.Writer, attr)
}

func (c *CardController) ClientsHtml(ctx *gin.Context) {
	attr := map[string]interface{}{}
	t, err := template.ParseFiles("./server/resource/clients.html")
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Content-Type", "text/html")
	t.Execute(ctx.Writer, attr)
}

func (c *CardController) CardsHtml(ctx *gin.Context) {
	attr := map[string]interface{}{}
	t, err := template.ParseFiles("./server/resource/cards.html")
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Content-Type", "text/html")
	t.Execute(ctx.Writer, attr)
}

func (c *CardController) TransactionsHtml(ctx *gin.Context) {
	attr := map[string]interface{}{}
	t, err := template.ParseFiles("./server/resource/transactions.html")
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("Content-Type", "text/html")
	t.Execute(ctx.Writer, attr)
}
