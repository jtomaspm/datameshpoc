package server

import (
	"datamesh.poc/card-system-api/server/controller"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Host    string
	Port    string
	Context string
}

type Server struct {
	config *Config
	router *gin.Engine
}

func New(config *Config) *Server {
	s := &Server{
		config: config,
		router: gin.Default(),
	}
	s.Setup()
	return s
}

func (s *Server) Config() Config {
	return Config{
		Host:    s.config.Host,
		Port:    s.config.Port,
		Context: s.config.Context,
	}
}

func (s *Server) Setup() {
	controller.New(s.router)
}

func (s *Server) Run() {
	s.router.Run(s.config.Host + ":" + s.config.Port)
}
