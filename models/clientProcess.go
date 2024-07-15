package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	Username string
	Email    string
	Password string
}

func NewClient(u, e, p string) Client {
	return Client{Username: u, Email: e, Password: p}
}

func ProcessSignup(client Client) error {
	collection, err := SetCollection("mongodb.uri", "mongodb.database", "mongodb.clientsAccountCollection")
	if err != nil {
		return err
	}

	var existingUsernameClient Client
	err = collection.FindOne(context.Background(), bson.M{"username": client.Username}).Decode(&existingUsernameClient)
	if existingUsernameClient.Username != "" {
		return errors.New("username has already existed.\n使用者名稱已存在")
	} else if err != mongo.ErrNoDocuments {
		return err
	}

	var existingEmailClient Client
	err = collection.FindOne(context.Background(), bson.M{"email": client.Email}).Decode(&existingEmailClient)
	if existingEmailClient.Email != "" {
		return errors.New("email has already existedn.\n電子郵件已存在")
	} else if err != mongo.ErrNoDocuments {
		return err
	}

	_, err = collection.InsertOne(context.Background(), client)
	if err != nil {
		return err
	}

	return nil
}
