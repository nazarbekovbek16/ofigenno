package main

import (
	"context"
	"nosql/internal/data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	client *mongo.Client
	ctx    context.Context
}

func (app *Application) Insert(item *data.Item) error {
	DB := app.db.client.Database("barket")
	itemCollection := DB.Collection("items")
	_, err := itemCollection.InsertOne(app.db.ctx, bson.D{
		{Key: "description", Value: item.Description},
		{Key: "marketid", Value: item.MarketID},
		{Key: "price", Value: item.Price},
		{Key: "expireddate", Value: item.ExpiredDate},
		{Key: "name", Value: item.Name},
		{Key: "category", Value: item.Category},
	})
	if err != nil {
		return err
	}
	return nil
}

func (app *Application) Get(id string) (*data.Item, error) {
	DB := app.db.client.Database("barket")
	itemCollection := DB.Collection("items")
	var item bson.M
	var result data.Item

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := itemCollection.FindOne(app.db.ctx, bson.M{"_id": objectId}).Decode(&item); err != nil {
		return nil, err
	}

	bsonBytes, _ := bson.Marshal(item)
	bson.Unmarshal(bsonBytes, &result)
	result.ID = id
	return &result, nil
}
func (app *Application) Update(item *data.Item) error {
	DB := app.db.client.Database("barket")
	itemCollection := DB.Collection("items")
	_, err := itemCollection.UpdateOne(context.TODO(), bson.D{{"_id", item.ID}}, bson.D{{"$set", bson.D{
		{"description", item.Description},
		{"marketid", item.MarketID},
		{"price", item.Price},
		{"expireddate", item.ExpiredDate},
		{"name", item.Name},
		{"category", item.Category},
	},
	}})
	if err != nil {
		return err
	}
	return nil
}
func (app *Application) Delete(id string) error {
	DB := app.db.client.Database("barket")
	itemCollection := DB.Collection("items")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	_, err = itemCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (app *Application) GetAllItems() {

	// DB := app.db.client.Database("barket")
	// itemCollection := DB.Collection("items")
	// var item bson.M
	// var result data.Item

	// objectId, err := primitive.ObjectIDFromHex()
	// if err != nil {
	// 	return nil, err
	// }

	// if err := itemCollection.FindOne(app.db.ctx, bson.M{"_id": objectId}).Decode(&item); err != nil {
	// 	return nil, err
	// }

	// bsonBytes, _ := bson.Marshal(item)
	// bson.Unmarshal(bsonBytes, &result)
	// result.ID = id
	// return &result, nil
}
