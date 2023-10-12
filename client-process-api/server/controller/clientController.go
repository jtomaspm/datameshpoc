package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"datamesh.poc/client-process-api/logger"
	"datamesh.poc/client-process-api/logger/message"
	"datamesh.poc/client-process-api/server/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var client model.Client
	err := ctx.BindJSON(&client)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	body, err := json.Marshal(client.ToPersonBase())
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := http.Post("http://person-system-api:8080/api/person", "application/json", bytes.NewBuffer(body))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		ctx.JSON(400, gin.H{"error": "Person not created"})
		return
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var m map[string]string
	err = json.Unmarshal(body, &m)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	id, ok := m["id"]
	if !ok {
		ctx.JSON(500, gin.H{"error": "Person not created"})
		return
	}
	newClient := client.ToClientBase(uuid.MustParse(id))
	body, err = json.Marshal(newClient)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err = http.Post("http://client-system-api:8080/api/client", "application/json", bytes.NewBuffer(body))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		ctx.JSON(400, gin.H{"error": "Client not created"})
		return
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	clientId, ok := m["id"]
	if !ok {
		ctx.JSON(500, gin.H{"error": "Client not created"})
		return
	}
	l := logger.New()
	save, err := json.Marshal(map[string]string{
		"clientId": clientId,
		"personId": id,
	})
	if err != nil {
		log.Println(err)
	} else {
		l.Log(message.Info("Client created", string(save)))
	}
	ctx.JSON(200, gin.H{"id": id})
}

func (c *ClientController) GetClient(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := http.Get("http://client-system-api:8080/api/client/" + id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		ctx.JSON(400, gin.H{"error": "Client not found"})
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var client model.ClientBase
	err = json.Unmarshal(body, &client)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	resp, err = http.Get("http://person-system-api:8080/api/person/" + client.PersonId.String())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		ctx.JSON(400, gin.H{"error": "Person not found"})
		return
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var person model.PersonBase
	err = json.Unmarshal(body, &person)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, model.NewClient(&person, &client))
}

func (c *ClientController) GetClients(ctx *gin.Context) {
	resp, err := http.Get("http://client-system-api:8080/api/client")
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		ctx.JSON(400, gin.H{"error": "Clients api error"})
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var clients []model.ClientBase
	err = json.Unmarshal(body, &clients)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	res := []model.Client{}
	for _, client := range clients {
		resp, err = http.Get("http://person-system-api:8080/api/person/" + client.PersonId.String())
		if err != nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			log.Println(err)
			continue
		}
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			continue
		}
		var person model.PersonBase
		err = json.Unmarshal(body, &person)
		if err != nil {
			log.Println(err)
			continue
		}
		res = append(res, *model.NewClient(&person, &client))
	}
	ctx.JSON(200, res)
}
