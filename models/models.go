package models 

import (
  "context"
  "go.mongodb.org/mongo-driver/mongo"
) 

func InsertOne (client *mongo.Client, ctx context.Context, database string, col string, doc interface{}) (*mongo.InsertOneResult, error) {
 
    // select database and collection with Client.Database method
    // and Database.Collection method
    collection := client.Database(database).Collection(col)
     
    // InsertOne accept two argument of type Context
    // and of empty interface  
    result, err := collection.InsertOne(ctx, doc)
    return result, err
}
 
// insertMany is a user defined method, used to insert
// documents into collection returns result of
// InsertMany and error if any.
func InsertMany (client *mongo.Client, ctx context.Context, database string, col string, docs []interface{}) (*mongo.InsertManyResult, error) {
 
    // select database and collection ith Client.Database
    // method and Database.Collection method
    collection := client.Database(database).Collection(col)
     
    // InsertMany accept two argument of type Context
    // and of empty interface  
    result, err := collection.InsertMany(ctx, docs)
    return result, err
}
