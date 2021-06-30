package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/odit-bit/covid-tracker/service"
)

type TrackerHandlers struct {
	service service.CovidDataService
}

// func (t TrackerHandlers) FindData() (*model.CovidData, error) {
// 	return t.service.FindData()
// }

func (th *TrackerHandlers) GetCovidData() func(c *fiber.Ctx) error {
	data, err := th.service.FindData()
	if err != nil {
		log.Fatal(err)
	}
	r := func(c *fiber.Ctx) error {
		return c.JSON(data)
	}
	return r
}

func NewTrackerHandlers(service service.CovidDataService) TrackerHandlers {
	return TrackerHandlers{service: service}
}
