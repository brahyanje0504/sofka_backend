package routers

import (
	"fmt"
	"github.com/brahyanje0504/sofka_backend/infrastructure/entrypoint/rest/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

type router struct {
	app               *fiber.App
	spacecraftHandler *handlers.SpacecraftHandler
}

func NewRouter(app *fiber.App, spacecraftHandler *handlers.SpacecraftHandler) *router {
	return &router{
		app:               app,
		spacecraftHandler: spacecraftHandler,
	}
}

func (r *router) registry() {
	r.app.Post("spaceships/create", r.spacecraftHandler.CreateSpacecraft)
	r.app.Get("spaceships", r.spacecraftHandler.GetAll)
	r.app.Get("spaceships/countries", r.spacecraftHandler.GetCountries)
	r.app.Get("spaceships/fuels", r.spacecraftHandler.GetFuels)
}

func (r *router) Start() error {

	r.app.Use(cors.New())
	r.app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	r.registry()

	port := os.Getenv("PORT")
	return r.app.Listen(fmt.Sprintf(":%v", port))
}
