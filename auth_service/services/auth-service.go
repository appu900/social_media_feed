package services

import (
	"auth-service/db"
	"auth-service/model"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)
func RegisterUser(username string, email string, password string) (string, error) {
	// ** check if the user already exits
	collection := db.GetCollections(db.Client, "users")
	ctx,cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser model.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&existingUser)
	if err == nil{
		return "", errors.New("User already exists")
	}

	// ** hash the password
    return "hello",nil
}