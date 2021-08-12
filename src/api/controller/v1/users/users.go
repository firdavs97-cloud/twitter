package users

import (
	"github.com/gofiber/fiber/v2"
	"twitter.go/src/models/models_v1"
	"twitter.go/src/utils/api_response"
)

func Router(c fiber.Router) {
	c.Post("/", newUser)
	c.Post("/follow", followUser)
}

func newUser(c *fiber.Ctx) error {
	user := models_v1.User{}
	if err := c.BodyParser(&user); err != nil {
		return api_response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}
	if errors := api_response.ValidateStruct(&user); errors != nil {
		return api_response.NewErrorResponse(c, fiber.StatusBadRequest, errors)
	}
	if _, err := models_v1.GetUser(user.ID); err != nil {
		return api_response.NewErrorResponse(c, fiber.StatusConflict, err)
	}
	if err := user.Save(); err != nil {
		return api_response.NewErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return api_response.MessageResponse(c, "Ok")
}

func followUser(c *fiber.Ctx) error {
	follow := models_v1.FollowUserRequest{}
	if err := c.BodyParser(&follow); err != nil {
		return api_response.NewErrorResponse(c, fiber.StatusBadRequest, err)
	}
	if errors := api_response.ValidateStruct(&follow); errors != nil {
		return api_response.NewErrorResponse(c, fiber.StatusBadRequest, errors)
	}
	if _, err := models_v1.GetUser(follow.IdFollowing); err != nil {
		return api_response.NewErrorResponse(c, fiber.StatusConflict, err)
	}
	follower, err := models_v1.GetUser(follow.IdFollower)
	if err != nil {
		return api_response.NewErrorResponse(c, fiber.StatusConflict, err)
	}
	if err = follower.Follow(follow.IdFollowing); err != nil {
		return api_response.NewErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return api_response.MessageResponse(c, "Ok")
}
