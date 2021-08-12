package routers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"time"
	"twitter.go/src/api/controller/v1/tweets"
	"twitter.go/src/api/controller/v1/users"
)

func SetupRouter(app fiber.Router) {
	app.Use(cors.New(cors.ConfigDefault))
	app.Get("/", initPage)
	app.Use(cors.New(cors.ConfigDefault))
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true" || c.Response().StatusCode() != fiber.StatusOK
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))
	app.Post("ping/", pong)
	api := app.Group("/api")
	apiV1 := api.Group("/v1")
	tweets.Router(apiV1.Group("/tweets"))
	users.Router(apiV1.Group("/users"))
}

func initPage(c *fiber.Ctx) error {
	var response = struct {
		BuildNumber string
		BranchName  string
	}{BuildNumber: os.Getenv("buildNumber"), BranchName: os.Getenv("branchName")}
	return c.JSON(response)
}

type Ping struct {
	Message    string `json:"message"`
	PingNumber int    `json:"pingNumber"`
}

func pong(c *fiber.Ctx) error {
	body := new(Ping)
	if err := c.BodyParser(body); err != nil || body.Message != "ping" {
		return c.SendStatus(fiber.StatusBadRequest) // TODO message
	}
	return c.Status(fiber.StatusOK).JSON(fmt.Sprintf("pong to pingNumber: %d", body.PingNumber))
}
