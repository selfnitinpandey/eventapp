package models

import (
	"context"
	"eventapp/config"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Event struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Location   string             `bson:"location" json:"location"`
	Created_At time.Time          `bson:"created_at" json:"created_at"`
}

func GetEvent(client *mongo.Client, ctx context.Context) ([]Event, error) {
	var events []Event

	appConfig := config.LoadConfig()
	collection := client.Database(appConfig.DATABASENAME).Collection("events")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var event Event
		err = cursor.Decode(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil

}

func InsertEvent(client *mongo.Client, ctx context.Context, event Event) error {
	appConfig := config.LoadConfig()
	collection := client.Database(appConfig.DATABASENAME).Collection("events")
	_, err := collection.InsertOne(ctx, event)
	return err
}

func UpdateEvent(client *mongo.Client, ctx context.Context, eventID string, updatedEvent Event) error {
	appConfig := config.LoadConfig()
	objectID, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		return err
	}
	collection := client.Database(appConfig.DATABASENAME).Collection("events")
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"name":       updatedEvent.Name,
			"location":   updatedEvent.Location,
			"created_at": updatedEvent.Created_At,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update event: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no event found with the given ID")
	}

	return nil
}

func DeleteEvent(client *mongo.Client, ctx context.Context, delId string) (error, int64) {
	appConfig := config.LoadConfig()
	objectID, err := primitive.ObjectIDFromHex(delId)
	if err != nil {
		log.Fatalf("Invalid ObjectID format: %v", err)
		return err, 0
	}

	collection := client.Database(appConfig.DATABASENAME).Collection("events")

	deleteResult, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		log.Fatalf("Failed to delete event: %v", err)
		return err, 0
	}

	return nil, deleteResult.DeletedCount

}
