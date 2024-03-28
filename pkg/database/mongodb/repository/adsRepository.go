package repository

import (
	"context"
	"time"
	"vk-test/pkg/database/mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var adsCollection *mongo.Collection = OpenCollection(Client, "ads")

// Function that creates a project in the database and updates the project count
func CreateProject(project models.Ad) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	return nil
}

// Function that updates a project in the database
func UpdateProject(projectId int, updatedProject models.Ad) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"id": projectId}
	oldOne, err := GetProjectByID(projectId)
	if err != nil {
		return err
	}
	updatedProject.Id = oldOne.Id
	updatedProject.ID = oldOne.ID
	update := bson.M{"$set": updatedProject}

	_, err = adsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// Function that retrieves a project from the database by its custom ID
func GetProjectByID(projectId int) (foundProject models.Ad, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"id": projectId}
	err = adsCollection.FindOne(ctx, filter).Decode(&foundProject)
	if err != nil {
		return
	}

	return
}

// Function that returns all projects from the database
func GetAllProjects() ([]models.Ad, error) {
	// Define the context for the operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define options for the find operation
	findOptions := options.Find()

	// Perform the find operation to retrieve all project documents
	cursor, err := adsCollection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate over the results and construct project objects
	var projects []models.Ad
	for cursor.Next(ctx) {
		var project models.Ad
		if err := cursor.Decode(&project); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
