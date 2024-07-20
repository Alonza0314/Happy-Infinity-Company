package models

import (
	"context"
	"errors"
	"time"

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
		return errors.New("username has already existed\n使用者名稱已存在")
	} else if err != mongo.ErrNoDocuments {
		return err
	}

	var existingEmailClient Client
	err = collection.FindOne(context.Background(), bson.M{"email": client.Email}).Decode(&existingEmailClient)
	if existingEmailClient.Email != "" {
		return errors.New("email has already existedn\n電子郵件已存在")
	} else if err != mongo.ErrNoDocuments {
		return err
	}

	_, err = collection.InsertOne(context.Background(), client)
	if err != nil {
		return err
	}

	return nil
}

func ProcessSignin(client Client) error {
	collection, err := SetCollection("mongodb.uri", "mongodb.database", "mongodb.clientsAccountCollection")
	if err != nil {
		return err
	}

	var existingUsernameClient Client
	err = collection.FindOne(context.Background(), bson.M{"username": client.Username}).Decode(&existingUsernameClient)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("username doesn't exist\n\n使用者名稱不存在")
		}
		return err
	}

	if client.Password != existingUsernameClient.Password {
		return errors.New("wrong password\n\n密碼錯誤")
	}

	return nil
}

func ProcessPwfind(client Client) error {
	collection, err := SetCollection("mongodb.uri", "mongodb.database", "mongodb.clientsAccountCollection")
	if err != nil {
		return err
	}

	var existingUsernameClient Client
	err = collection.FindOne(context.Background(), bson.M{"username": client.Username}).Decode(&existingUsernameClient)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("username doesn't exist.\n\n使用者名稱不存在")
		}
		return err
	}

	if client.Email != existingUsernameClient.Email {
		return errors.New("email error.\n\n電子郵件錯誤")
	}

	return nil
}

func ProcessPwreset(client Client) error {
	collection, err := SetCollection("mongodb.uri", "mongodb.database", "mongodb.clientsAccountCollection")
	if err != nil {
		return err
	}

	filter := map[string]interface{}{
		"username": client.Username,
	}

	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"password": client.Password,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err = collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.New("server error\npassword change fail\n\n服務器錯誤，密碼更改失敗")
	}

	return nil
}
