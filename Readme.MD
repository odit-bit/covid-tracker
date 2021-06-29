a little about Hexagonal

the fundamental is around business (logic.go) and domain (model.go)

service is create some work using entity , in this scope service is pull or push data from/to repository as an object (entity/struct)

##Core Definition##

1. what we want to build?
   an app who can present and store some data

2. what data ?
   Covid19 data

3. where we get the data ?
   from our database

4. where to present the data ?
   to user interface

##Abstraction##
from the core definition we need to create service who can provide data in the form of covid19Data and presenter can get those data

##implementation

CovidData is an entity form of data use in the service

CovidDataService is a process who provide CovidData model

CovidDataRepository data types that store and retrieve CovidData

implemnetaion of repository must satisfied the CovidDataRepository type so CovidDataRepository type can call function who can retrieve data that repository provided

implementation of service must satisfied the CovidDataService type so those service can call function that returns CovidData from repository

everything should dependent (satisfied) model package
