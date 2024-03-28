package repository

import (
	"context"
	"time"
	"vk-test/pkg/database/mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var adsCollection *mongo.Collection = OpenCollection(Client, "ads")

// Function that creates an ad in the database and updates the ad count
func CreateAd(ad models.Ad) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	ad.ID = primitive.NewObjectID()

	_, err = adsCollection.InsertOne(ctx, ad)
	if err != nil {
		return err
	}

	return nil
}

// Function that retrieves a ad from the database by its custom ID
func GetAdByID(adId int) (foundad models.Ad, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"id": adId}
	err = adsCollection.FindOne(ctx, filter).Decode(&foundad)
	if err != nil {
		return
	}

	return
}

// Function that returns all ads from the database
func GetAllAds() ([]models.Ad, error) {
	// Define the context for the operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define options for the find operation
	findOptions := options.Find()

	// Perform the find operation to retrieve all ad documents
	cursor, err := adsCollection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate over the results and construct ad objects
	var ads []models.Ad
	for cursor.Next(ctx) {
		var ad models.Ad
		if err := cursor.Decode(&ad); err != nil {
			return nil, err
		}
		ads = append(ads, ad)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return ads, nil
}

func GetAdsByPage(page int) (ads []models.Ad, err error) {
	// Define the context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	const pageSize = 5 // Example: Retrieve 20 ads per page

	// Calculate offset (number of documents to skip)
	offset := (page - 1) * pageSize

	// Create options for the Find operation with skip and limit
	options := options.Find().SetSkip(int64(offset)).SetLimit(int64(pageSize))

	// Execute the query
	cursor, err := adsCollection.Find(ctx, bson.D{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode results into the 'ads' slice
	if err = cursor.All(ctx, &ads); err != nil {
		return nil, err
	}

	return ads, nil
}
