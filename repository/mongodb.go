package repository

import (
	"context"
	"time"

	"github.com/odit-bit/covid-tracker/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepo struct {
	client  *mongo.Client
	dbName  string
	timeout time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepo(mongoURL, name string, mongoTimeout int) (model.CovidDataRepository, error) {
	c, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, err
	}
	repo := &mongoRepo{client: c, dbName: name, timeout: time.Duration(mongoTimeout) * time.Second}

	return repo, nil
}

func (r *mongoRepo) FindData() (*model.CovidData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	covidData := &model.CovidData{}
	collection := r.client.Database(r.dbName).Collection("kasus")
	filter := bson.M{"name": "rw011"}
	err := collection.FindOne(ctx, filter).Decode(&covidData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}
	return covidData, nil
}
