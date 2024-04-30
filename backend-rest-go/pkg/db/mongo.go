package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tonitienda/kadai/backend-rest-go/pkg/tasks"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskDB struct {
	ID          string    `bson:"id"`
	OwnerID     string    `bson:"owner_id"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	Status      string    `bson:"status"`
	DeletedAt   time.Time `bson:"deleted_at"`
	DeletedBy   string    `bson:"deleted_by"`
}

type MongoDB struct {
	client *mongo.Client
}

func taskDbtoTasks(dbTasks []TaskDB) []tasks.Task {
	newTasks := make([]tasks.Task, len(dbTasks))

	for _, t := range dbTasks {
		newTasks = append(newTasks, tasks.Task(t))
	}

	return newTasks

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

func (db *MongoDB) getAllTasks() ([]tasks.Task, error) {
	// Maybe we can initialize this collection as part of the MongoDB struct
	// initialization, and not every time.
	collection := db.client.Database("kadai").Collection("tasks")

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var results []TaskDB

	if err = cur.All(context.Background(), &results); err != nil {
		return taskDbtoTasks(results), err
	}

	return taskDbtoTasks(results), nil
}

func (db *MongoDB) GetTasks(ownerID string) ([]tasks.Task, error) {

	allTasks, _ := db.getAllTasks()

	fmt.Println("All Tasks:", allTasks)

	// Maybe we can initialize this collection as part of the MongoDB struct
	// initialization, and not every time.
	collection := db.client.Database("kadai").Collection("tasks")

	fmt.Println("Looking for tasks that belong to ", ownerID)

	cur, err := collection.Find(context.Background(),
		bson.M{
			"owner_id": ownerID,
			"$or": []interface{}{
				bson.M{"deleted_at": nil},
				bson.M{"deleted_at": time.Time{}},
			},
		})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var results []tasks.Task

	if err = cur.All(context.Background(), &results); err != nil {
		return results, err
	}

	fmt.Println("User active tasks:", results)

	return results, nil
}

func (db *MongoDB) GetTask(taskID string) (tasks.Task, bool) {
	var task TaskDB
	collection := db.client.Database("kadai").Collection("tasks")

	err := collection.FindOne(context.Background(), bson.D{}).Decode(&task)

	if err != nil {
		return tasks.Task(task), false
	}

	return tasks.Task(task), true
}

func (db *MongoDB) AddTask(task tasks.Task) error {
	dbTask := TaskDB(task)

	collection := db.client.Database("kadai").Collection("tasks")

	_, err := collection.InsertOne(context.Background(), dbTask)

	if err != nil {
		return err
	}

	return nil
}

func (db *MongoDB) UpdateTask(task tasks.Task) error {

	existingTask, _ := db.GetTask(task.ID)
	fmt.Println("Updating task", existingTask, "to", task)
	collection := db.client.Database("kadai").Collection("tasks")

	dbTask := TaskDB(task)

	_, err := collection.UpdateOne(context.Background(), bson.M{"id": task.ID}, bson.M{"$set": dbTask})

	if err != nil {
		return err
	}

	existingTask, _ = db.GetTask(task.ID)
	fmt.Println("After updating task", existingTask)

	return nil
}

func (db *MongoDB) DeleteTask(id string) error {
	collection := db.client.Database("kadai").Collection("tasks")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		return err
	}

	return nil
}
