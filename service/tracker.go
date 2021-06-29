package service

import "github.com/odit-bit/covid-tracker/model"

// API endpoint handler need to fulfill this signature
type CovidDataService interface {
	FindData() (*model.CovidData, error)
	// StoreData() error
}

// with go composition is done with embedding type (duck typing)
type DefaultCovidDataService struct {
	repo model.CovidDataRepository
}

//FindData() calling the field repo method signature FindData() (CovidDataRepository type)
func (cd *DefaultCovidDataService) FindData() (*model.CovidData, error) {
	// repository type method signature
	return cd.repo.FindData()
}

//initialize function to create service object(struct) with embedding repository type
func New(repo model.CovidDataRepository) *DefaultCovidDataService {
	return &DefaultCovidDataService{repo: repo}
}
