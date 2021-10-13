package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

type Task struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Text      string             `bson:"text"`
	Completed bool               `bson:"completed"`
}

func main() {
	credential := options.Credential{
		Username: "admin",
		Password: "123456",
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/").SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	str := "Tarefa 1"

	task := &Task{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Text:      str,
		Completed: false,
	}

	collection = client.Database("local").Collection("tasks")

	// persistindo informações
	// createTask(task)

	// recuperando informações
	// getAll()
	// cur, err := collection.Find(ctx, bson.D{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cur.Close(ctx)
	// for cur.Next(ctx) {
	// 	var result bson.D
	// 	err := cur.Decode(&result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(result)
	// }
	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// recuperando uma informação
	// filter := bson.D{"text", "Tarefa"}

	// filter = append(filter, bson.E{Key: "name", Value: bson.M{"$text": bson.M{"$search": f.Name}}})

	query := bson.M{
		"text": "Tarefa 6",
	}

	//ctx, collection = context.WithTimeout(context.Background(), 5*time.Second)
	//defer collection()

	var result struct {
		Value Task
	}
	err = collection.FindOne(ctx, query).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	fmt.Println("Fim", task)

}

func createTask(task *Task) error {
	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func getAll() ([]*Task, error) {
	// passing bson.D{{}} matches all documents in the collection
	filter := bson.D{}
	return filterTasks(filter)
}

func filterTasks(filter interface{}) ([]*Task, error) {
	// A slice of tasks for storing the decoded documents
	var tasks []*Task

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return tasks, err
	}

	for cur.Next(ctx) {
		var t Task
		err := cur.Decode(&t)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, &t)
	}

	if err := cur.Err(); err != nil {
		return tasks, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil
}

func printTasks(tasks []*Task) {
	for i, v := range tasks {
		if v.Completed {
			fmt.Printf("%d: %s\n", i+1, v.Text)
		} else {
			fmt.Printf("%d: %s\n", i+1, v.Text)
		}
	}
}
