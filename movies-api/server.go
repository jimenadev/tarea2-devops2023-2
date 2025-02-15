package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lnds/lab-ic-ms/api/database"
	"github.com/lnds/lab-ic-ms/api/handlers"
)

func server(bind string, port string) {
	database.ConnectDb(database.GetDSN())

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,OPTIONS",
	}))

	setupRoutes(app)

	app.Listen(bind + ":" + port)
}

func setupRoutes(app *fiber.App) {
	app.Get("/movies", handlers.ListMovies)
	app.Get("/movies/:id", handlers.GetMovie)
	app.Post("/movies", handlers.CreateMovie)
	app.Put("/movies", handlers.UpdateMovie)

	app.Get("/directors", handlers.ListDirectors)
	app.Get("/directors/:id", handlers.GetDirector)
	app.Post("/directors", handlers.CreateDirector)
	app.Put("/directories", handlers.UpdateMovie)
}
