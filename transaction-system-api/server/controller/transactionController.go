package controller

import (
	"log"
	"net/http"

	"datamesh.poc/transaction-system-api/dal/context"
	"datamesh.poc/transaction-system-api/dal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransationController struct {
	router    *gin.Engine
	dbCtx     *context.DbContext
	baseRoute string
}

func New(router *gin.Engine, baseRoute string) *TransationController {
	c := &TransationController{
		router:    router,
		dbCtx:     context.New(),
		baseRoute: baseRoute,
	}
	err := c.dbCtx.MakeMigrations()
	if err != nil {
		log.Println(err)
	}
	c.Setup()
	return c
}

func (c *TransationController) Setup() {
	c.router.POST(c.baseRoute, c.CreateTransaction)
	c.router.GET(c.baseRoute+"/:id", c.GetTransaction)
	c.router.GET(c.baseRoute, c.GetTransactions)
}

func (c *TransationController) CreateTransaction(ctx *gin.Context) {
	var transaction model.Transaction
	err := ctx.BindJSON(&transaction)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Println(transaction)
	r, err := http.Get("http://card-system-api:8080/api/card/" + transaction.CardId.String())
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		log.Println("Card not found")
		ctx.JSON(400, gin.H{"error": "Card not found"})
		return
	}
	id, err := c.dbCtx.CreateTransaction(transaction)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"id": id})
}

func (c *TransationController) GetTransaction(ctx *gin.Context) {
	idu, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	transaction, err := c.dbCtx.GetTransaction(idu)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, transaction)
}

func (c *TransationController) GetTransactions(ctx *gin.Context) {
	res, err := c.dbCtx.GetTransactions()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}
