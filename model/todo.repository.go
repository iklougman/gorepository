package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepo struct {
	mongoCollection *mongo.Collection
	ctx             context.Context
}

func NewTodoRepo(client *mongo.Client, ctx context.Context, dbName string) *TodoRepo {
	collection := client.Database(dbName).Collection("todo")
	return &TodoRepo{
		mongoCollection: collection,
		ctx:             ctx,
	}
}

func (r *TodoRepo) FindAll() ([]TodoItem, error) {
	cur, err := r.mongoCollection.Find(context.Background(), bson.D{})
	// fmt.Println(r.ctx, r.mongoCollection.Name(), err)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var items []TodoItem
	if err = cur.All(context.Background(), &items); err != nil {
		log.Fatal(err)
	}
	fmt.Println(items)
	return items, nil
}

func (r *TodoRepo) Insert(items []TodoItem) (*mongo.InsertManyResult, error) {
	entries := make([]interface{}, len(items))
	for i, v := range items {
		entries[i] = v
	}
	result, err := r.mongoCollection.InsertMany(context.Background(), entries)
	if err != nil {
		return nil, err
	}

	return result, nil
}
