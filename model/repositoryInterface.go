package model

type CovidDataRepository interface {
	FindData() (*CovidData, error)
}
