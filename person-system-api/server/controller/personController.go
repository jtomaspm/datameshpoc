package controller

import (
	"log"

	"datamesh.poc/person-system-api/dal/context"
	"datamesh.poc/person-system-api/dal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PersonController struct {
	router *gin.Engine
	dbCtx  *context.DbContext
}

func New(router *gin.Engine) *PersonController {
	c := &PersonController{
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

func (c *PersonController) Setup() {
	c.router.POST("/api/person", c.CreatePerson)
	c.router.GET("/api/person/:id", c.GetPerson)
	c.router.GET("/api/person", c.GetPersons)
}

func (c *PersonController) CreatePerson(ctx *gin.Context) {
	var person model.Person
	err := ctx.BindJSON(&person)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Println(person)
	id, err := c.dbCtx.CreatePerson(person)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"id": id})
}

func (c *PersonController) GetPerson(ctx *gin.Context) {
	idu, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	person, err := c.dbCtx.GetPerson(idu)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, person)
}

func (c *PersonController) GetPersons(ctx *gin.Context) {
	res, err := c.dbCtx.GetPersons()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}
