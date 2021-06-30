package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/odit-bit/covid-tracker/api"
	"github.com/odit-bit/covid-tracker/repository"
	"github.com/odit-bit/covid-tracker/service"
)

func main() {
	//server config
	// port := os.Getenv("port")

	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }

	//repo config
	mongoURL := os.Getenv("db_url") //"mongodb://localhost:27017"
	if mongoURL == "" {
		log.Fatal("$db_Url must be set")
	}
	mongoDataBase := "covid"
	mongoTimeout := 10

	//repository instance
	repository, _ := repository.NewMongoRepo(mongoURL, mongoDataBase, mongoTimeout)
	// service instance
	tracker := service.New(repository)

	//instaniate handler route
	kicuy := api.NewTrackerHandlers(tracker)

	// now we can call service method who call repository method to get data

	app := fiber.New()

	app.Use("/admin", basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "ganteng",
		},
	}))

	//CORS
	app.Use("/", cors.New())

	//logger
	app.Use(logger.New(logger.Config{
		Format:     "${header:} [${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
	}))

	/* endpoint */

	app.Get("/", kicuy.GetCovidData())

	app.Post("/admin", func(c *fiber.Ctx) error {
		return c.SendString(string(c.Request().Header.Header()))
	})

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("SERVER FAIL")
	}
}
