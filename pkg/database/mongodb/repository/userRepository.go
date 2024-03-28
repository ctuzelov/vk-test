package repository

import (
	"context"
	"time"
	"vk-test/pkg/database/mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = OpenCollection(Client, "user")

// Function that updates user's tokens in the database
func UpdateTokens(signedToken string, signedRefreshToken string, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"token":         signedToken,
			"refresh_token": signedRefreshToken,
		},
	}

	filter := bson.M{"user_id": userId}

	_, err := userCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return err
}

// Function that creates a user in the database
func CreateUser(user models.User) (err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	_, err = userCollection.InsertOne(ctx, user)

	return
}

// Function that returns the count of users with a given email from the database
func CountUsersByEmail(email string) (count int64, err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	count, err = userCollection.CountDocuments(ctx, bson.M{"email": email})
	return
}

// Function that returns a user object from the database using their email
func GetUserByEmail(email string) (foundUser models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	err = userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&foundUser)

	return
}

func UpdateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"email": user.Email} // Use the user's unique ID as the filter

	update := bson.M{
		"$set": bson.M{
			"name":          user.Name,
			"number":        user.Number,
			"date_of_birth": user.DateOfBirth,
		},
	}

	_, err := userCollection.UpdateOne(ctx, filter, update)

	return err
}

// Function that returns all users from the database
func GetAllUsers() (users []models.User, err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// Define options for the find operation
	findOptions := options.Find()

	// Perform the find operation to retrieve all user documents
	cursor, err := userCollection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate over the results and construct user objects
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
