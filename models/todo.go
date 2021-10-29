package models

import (
	"context"
	"log"
	"stephan/todo/database"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Item      string             `bson:"item,omitempty" json:"item"`
	Completed bool               `bson:"completed"`
}

var todoCollection *mongo.Collection = database.OpenCollection(database.Client, "todos")
var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)

func getObjID(id string) primitive.ObjectID {
	objID, _ := primitive.ObjectIDFromHex(id)
	return objID
}

func SetupNewTodo(todo Todo) Todo {
	todo.ID = primitive.NewObjectID()
	todo.Completed = false
	return todo
}

func (t Todo) GetAllTodos() []Todo {
	cursor, err := todoCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	var todos []Todo
	if err = cursor.All(ctx, &todos); err != nil {
		log.Fatal(err)
	}

	return todos
}

func (t Todo) GetByID(id string) Todo {
	var todo Todo
	objID := getObjID(id)

	if err := todoCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&todo); err != nil {
		log.Fatal(err)
	}

	return todo
}

func (t Todo) InsertOne(c *gin.Context) *mongo.InsertOneResult {
	var todo Todo
	if err := c.BindJSON(&todo); err != nil {
		log.Fatal(err)
	}
	readyTodo := SetupNewTodo(todo)
	res, err := todoCollection.InsertOne(ctx, readyTodo)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (t Todo) DeleteOne(id string) (*mongo.DeleteResult, error) {
	objID := getObjID(id)

	res, err := todoCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t Todo) ToggleComplete(id string) error {
	currentTodo := t.GetByID(id)

	_, err := todoCollection.UpdateOne(
		ctx,
		bson.M{"_id": currentTodo.ID},
		bson.M{
			"$set": bson.M{"completed": !currentTodo.Completed},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
