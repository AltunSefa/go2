package authRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/cruxia/go/controllers"
	"github.com/cruxia/go/middleware"
	
)

func SetupRoutes( micro *fiber.App) {

	
	micro.Route("/auth", func(router fiber.Router) {
		router.Post("/register", controllers.SignUpUser)
		router.Post("/login", controllers.SignInUser)
		router.Get("/logout", middleware.DeserializeUser, controllers.LogoutUser)
	})
}