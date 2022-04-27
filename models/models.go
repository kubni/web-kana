package models 

import (
  "fmt"
  "context"
  "time"
  "log"
  "web_kana_v1/dbLogic"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson"
) 

type Model struct {
  client *mongo.Client
  dbName string 
  collectionName string
  ctx context.Context
}

// Practically a constructor 
func NewModel (client *mongo.Client, dbName string, collectionName string) *Model {
  return &Model {
    client: client,
    dbName: dbName,
    collectionName: collectionName,
  } 
}


func (m *Model) InsertOne(doc interface{}) (*mongo.InsertOneResult, error) {

  // We grab the specified collection 
  collection := dbLogic.GetCollection(m.client, m.dbName, m.collectionName) 

  // TODO: What to do with cancel?
  ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
  
  result, err := collection.InsertOne(ctx, doc)
  return result, err
}
 

func (m *Model) InsertMany(docs []interface{}) (*mongo.InsertManyResult, error) {
 
  collection := dbLogic.GetCollection(m.client, m.dbName, m.collectionName)
  ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
 
  result, err := collection.InsertMany(ctx, docs)
  return result, err
}


// TODO: A better way than a global value?
// At least find a better name 
type DocumentSchema struct{
  Username string // Has to have same name as the corresponding field in the database 
  Score int
}

func (m *Model) GetScoreboard() []DocumentSchema {

  collection := dbLogic.GetCollection(m.client, m.dbName, m.collectionName)

  // Sort by score 
  opts := options.Find().SetSort(bson.D{{Key:"Score", Value:-1}})

  // Find will return a Cursor, which is basically a pointer to the set of documents
  cursor, err := collection.Find(context.Background(), bson.D{}, opts)
  if err != nil {
    log.Fatal(err)
  }

  defer cursor.Close(context.Background())

  // Declare the slice of results (documents):
  // TODO: Is slice the best data structure for this?
  var scoreboard []DocumentSchema 

  // Iterate through the results and add them into the previously declared slice
  for cursor.Next(context.Background()) {
    result := DocumentSchema{} 

    // Decode bson into our chosen Golang data structure
    err := cursor.Decode(&result)
    if err != nil {
      log.Fatal(err) 
    }

    scoreboard = append(scoreboard, result)
  }

  if err := cursor.Err(); err != nil {
    log.Fatal(err)
  }

  // In case no error occured
  fmt.Println("Scoreboard: ", scoreboard)
  return scoreboard
}
