package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/cruxia/go/initializers"
	routerRouter "github.com/cruxia/go/router"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	
	initializers.ConnectDB(&config)
}

func main() {
	app := fiber.New()
	micro := fiber.New()
	routerRouter.SetupRoutes(app, micro)
	
}
