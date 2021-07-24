package database

import (
	"context"
	"time"

	"github.com/cbuelvasc/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConecctiont.Database("twittor")
	col := db.Collection("users")
	u.Password, _ = EncryptPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
