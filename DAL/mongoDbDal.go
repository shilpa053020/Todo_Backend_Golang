package DAL

import (
	"Todo_Backend_Golang/Model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDbDAl struct {
	Collection *mongo.Collection
}

func NewMongoDbDal(client *mongo.Client, dbName, collectionName string) *MongoDbDAl {
	collection := client.Database(dbName).Collection(collectionName)
	return &MongoDbDAl{Collection: collection}
}
//CREATE A TODO
func (db *MongoDbDAl) Create(ctx context.Context, todo Model.Todo) (*Model.Todo, error) {
	todo.Id = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := db.Collection.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

//GET ALL TODOS
func (db *MongoDbDAl)Get(ctx context.Context)([]Model.Todo,error){
	var todos []Model.Todo
    Cursor,err := db.Collection.Find(ctx,bson.M{})
	defer Cursor.Close(ctx)

	for Cursor.Next(ctx){
		var todo Model.Todo
		err := Cursor.Decode(&todo)
		if(err !=nil){
			 return nil, fmt.Errorf("error decoding todo: %w", err)
		}
		todos = append(todos,todo)
	}
	if(err != nil) {
		return nil,err
	}
	return todos ,nil
}

//UPDATE A TODO
func(db *MongoDbDAl)Update(ctx context.Context , id string , updatedTodo Model.Todo)(string,error){
	objectId,err := primitive.ObjectIDFromHex(id)
	if(err != nil){
		return "",err
	}
	filter := bson.M{"_id":objectId}
	update := bson.M{"$set": updatedTodo}
	result := db.Collection.FindOneAndUpdate(ctx,filter,update)
	if(result.Err()!=nil){
		return "",result.Err()
	}
	return "Updated successfully",nil
}

// DELETE A TODO
func (db *MongoDbDAl)Delete(ctx context.Context,id string)(string,error){
	objectId,err := primitive.ObjectIDFromHex(id)
	if(err != nil){
		return "",err
	}
	filter := bson.M{"_id":objectId}
	result, err := db.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return "", fmt.Errorf("failed to delete document: %w", err)
	}
	if result.DeletedCount == 0 {
		return "", fmt.Errorf("no document found with the given ID")
	}
	return "Deleted successfully",nil
}