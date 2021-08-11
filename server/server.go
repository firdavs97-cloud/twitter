package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"twitter.go/src/api/routers"
	"twitter.go/src/migrate"
)

type Server struct {
	Name   string
	Schema string
	Host   string
	Port   string
	fiber  *fiber.App
}

func BuildServer() *Server {
	app := fiber.New()
	routers.SetupRouter(app)
	return &Server{
		Name:   migrate.ServerConfigs.Server.Host + migrate.ServerConfigs.Server.Port,
		Schema: migrate.ServerConfigs.Server.Schema,
		Host:   migrate.ServerConfigs.Server.Host,
		Port:   migrate.ServerConfigs.Server.Port,
		fiber:  app,
	}
}

func (server *Server) Run() {
	log.Fatal(server.fiber.Listen(fmt.Sprintf("0.0.0.0:%s", server.Port)))
}

func (server *Server) Close() {
	if err := server.fiber.Shutdown(); err != nil {
		log.Println(err.Error())
	}
}
