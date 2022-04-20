package models 

import (
  "context"
  "go.mongodb.org/mongo-driver/mongo"
) 


// TODO: Turn this database initialization into a function and make it callable by the controller
// Or find a better way 

// Database URI 
  uri := "mongodb://localhost:27017"

  // Connect to the database
  client, ctx, cancel, err := dbLogic.ConnectTo(uri)
  if err != nil {
      panic(err)
  }
   
  // Release resource when the main function is returned.
  defer dbLogic.Close(client, ctx, cancel)
   
  // Ping the database 
  dbLogic.Ping(client, ctx, uri)

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
