package book

import (
	"github.com/cruxia/go/controllers"
	"github.com/cruxia/go/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes( micro *fiber.App) {

	
	micro.Get("/book", middleware.DeserializeUser, controllers.GetBook)
	micro.Post("/book", middleware.DeserializeUser, controllers.CreateNote)
}