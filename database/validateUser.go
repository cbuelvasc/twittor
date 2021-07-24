package database

import (
	"context"
	"time"

	"github.com/cbuelvasc/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ValidateUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConecctiont.Database("twittor")
	col := db.Collection("Users")
	condition := bson.M{"emil": email}
	var result models.User
	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
