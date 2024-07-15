package models

import (
	"context"
	"errors"
	"hic/configs"
	"log"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetCollection(uri, database, collection string) (*mongo.Collection, error) {
	mongodbUri, err := configs.GetConfigs(uri)
	if err != nil {
		log.Println(err)
		return nil, errors.New("server error.\n服務器錯誤")
	}

	clientOptions := options.Client().ApplyURI(mongodbUri)

	db, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println("server error => SetCollection function error\n\t" + err.Error())
		return nil, errors.New("server error.\n服務器錯誤")
	}

	err = db.Ping(context.Background(), nil)
	if err != nil {
		log.Println("server error => SetCollection function error\n\t" + err.Error())
		return nil, errors.New("server error.\n服務器錯誤")
	}

	mongodbDatabase, err := configs.GetConfigs(database)
	if err != nil {
		log.Println("server error => SetCollection function error\n\t" + err.Error())
		return nil, errors.New("server error.\n服務器錯誤")
	}

	mongodbCollection, err := configs.GetConfigs(collection)
	if err != nil {
		log.Println("server error => SetCollection function error\n\t" + err.Error())
		return nil, errors.New("server error.\n服務器錯誤")
	}

	return db.Database(mongodbDatabase).Collection(mongodbCollection), nil
}
