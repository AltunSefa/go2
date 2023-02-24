package user

import (
	"github.com/cruxia/go/controllers"
	"github.com/cruxia/go/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes( micro *fiber.App) {

	
	micro.Get("/users/me", middleware.DeserializeUser, controllers.GetMe)
	micro.Put("/users/:userId", middleware.DeserializeUser, controllers.UpdateUser)
}