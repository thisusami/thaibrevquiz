package controller

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/thisusami/thaibrevquiz/models"
	"github.com/thisusami/thaibrevquiz/services"
	"github.com/thisusami/thaibrevquiz/utils"
)

type Controller struct {
	fiber   *fiber.App
	service *services.Service
}

func (c *Controller) RegisterRoutes() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var user models.User
		log.Printf("{action:INBOUND,route:register,message:%v}", string(ctx.Body()))
		if err := json.Unmarshal(ctx.Body(), &user); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    fiber.StatusBadRequest,
				"error":   "Bad Request",
				"message": err.Error(),
			})
		}

		_, err := c.service.RegisterService(&user)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    fiber.StatusInternalServerError,
				"error":   "Internal Server Error",
				"message": err.Error(),
			})
		}
		log.Printf("{action:OUTBOUND,route:register,message:}")
		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"code":    fiber.StatusCreated,
			"message": "Success",
		})
	}
}
func (c *Controller) LoginRoutes(JWTSecret []byte) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var user models.User
		log.Printf("{action:INBOUND,route:login,message:%v}", string(ctx.Body()))
		if err := json.Unmarshal(ctx.Body(), &user); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    fiber.StatusBadRequest,
				"error":   "Bad Request",
				"message": err.Error(),
			})
		}
		result, err := c.service.LoginService(&user)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    fiber.StatusInternalServerError,
				"error":   "Internal Server Error",
				"message": err.Error(),
			})
		}
		if result == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"error":   "Unauthorized",
				"message": "Invalid credentials",
			})
		}
		token, err := utils.GenerateToken(result.Username, JWTSecret)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    fiber.StatusInternalServerError,
				"error":   "Internal Server Error",
				"message": err.Error(),
			})
		}
		log.Printf("{action:OUTBOUND,route:login,message:%v}", string(result.Username))
		return ctx.JSON(
			fiber.Map{
				"message":  "Success",
				"username": result.Username,
				"token":    token,
			},
		)
	}
}
func (c *Controller) GetRoutes() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Printf("{action:INBOUND,route:get,message:%v}", string(ctx.Body()))
		userID, err := utils.ExtractUserID(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"error":   "Unauthorized",
				"message": err.Error(),
			})
		}
		log.Printf("{action:OUTBOUND,route:get,message:%v}", string(userID))
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":     fiber.StatusOK,
			"message":  "Success",
			"username": userID,
		})
	}
}
func NewController(fiber *fiber.App, service *services.Service) *Controller {
	return &Controller{fiber: fiber, service: service}
}
