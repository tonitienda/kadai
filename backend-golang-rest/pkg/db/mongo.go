package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tonitienda/kadai/backend-golang-rest/pkg/tasks"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

func (db *MongoDB) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://database:27017"))
	if err != nil {
		return err
	}

	db.client = client
	return nil
}

func (db *MongoDB) GetTasks(ownerID string) ([]tasks.Task, error) {
	// Maybe we can initialize this collection as part of the MongoDB struct
	// initialization, and not every time.
	collection := db.client.Database("kadai").Collection("tasks")

	cur, err := collection.Find(context.Background(), bson.D{{"deletedAt", nil}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var results []tasks.Task

	if err = cur.All(context.Background(), &results); err != nil {
		return results, err
	}

	return results, nil
}

func (db *MongoDB) GetTask(taskID string) (tasks.Task, bool) {
	var task tasks.Task
	collection := db.client.Database("kadai").Collection("tasks")

	err := collection.FindOne(context.Background(), bson.D{}).Decode(&task)

	if err != nil {
		return task, false
	}

	return task, true
}

func (db *MongoDB) AddTask(task tasks.Task) error {
	collection := db.client.Database("kadai").Collection("tasks")

	_, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		return err
	}

	return nil
}

func (db *MongoDB) UpdateTask(task tasks.Task) error {

	existingTask, _ := db.GetTask(task.ID)
	fmt.Println("Updating task", existingTask, "to", task)
	collection := db.client.Database("kadai").Collection("tasks")

	_, err := collection.UpdateOne(context.Background(), bson.D{{"id", task.ID}}, bson.D{{"$set", task}})

	if err != nil {
		return err
	}

	existingTask, _ = db.GetTask(task.ID)
	fmt.Println("After updating task", existingTask)

	return nil
}

func (db *MongoDB) DeleteTask(id string) error {
	collection := db.client.Database("kadai").Collection("tasks")

	_, err := collection.DeleteOne(context.TODO(), bson.D{{"id", id}})
	if err != nil {
		return err
	}

	return nil
}
