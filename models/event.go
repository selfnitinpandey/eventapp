package models

import (
	"context"
	"eventapp/config"
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

func DeleteEvent(client *mongo.Client, ctx context.Context) {

}
