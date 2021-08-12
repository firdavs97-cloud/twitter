package tweets

import (
	"github.com/gofiber/fiber/v2"
	"twitter.go/src/models/models_v1"
	"twitter.go/src/utils/api_response"
)

func Router(c fiber.Router) {
	c.Post("/", newTweet)
	c.Post("/favourite", addFavourite)
}

func newTweet(c *fiber.Ctx) error {
	tweet := models_v1.Tweet{}
	if err := c.BodyParser(&tweet); err != nil {
		return api_response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}
	if errors := api_response.ValidateStruct(&tweet); errors != nil {
		return api_response.NewErrorResponse(c, fiber.StatusBadRequest, errors)
	}
	if _, err := models_v1.GetUser(tweet.Author); err != nil {
		return api_response.NewErrorResponse(c, fiber.StatusNotFound, err)
	}
	if err := tweet.Save(); err != nil {
		return api_response.NewErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return api_response.MessageResponse(c, "Ok")
}

func addFavourite(c *fiber.Ctx) error {
	fav := models_v1.AddFavouriteRequest{}
	if err := c.BodyParser(&fav); err != nil {
		return api_response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}
	if errors := api_response.ValidateStruct(&fav); errors != nil {
		return api_response.NewErrorResponse(c, fiber.StatusBadRequest, errors)
	}
	if _, err := models_v1.GetUser(fav.IdUser); err != nil {
		return api_response.NewErrorResponse(c, fiber.StatusNotFound, err)
	}
	tweet, err := models_v1.GetTweet(fav.IdTweet)
	if err != nil {
		return api_response.NewErrorResponse(c, fiber.StatusNotFound, err)
	}
	if err = tweet.AddFavourite(fav.IdUser); err != nil {
		return api_response.NewErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return api_response.MessageResponse(c, "Ok")
}
