package models 

import (
  "context"
  "web_kana_v1/dbLogic"

  "go.mongodb.org/mongo-driver/mongo"
) 

type Model struct {
  client *mongo.Client
  dbName string 
  collectionName string
  ctx context.Context
}

func NewModel (ctx context.Context, client *mongo.Client, dbName string, collectionName string) *Model {
  return &Model {
    client: client,
    dbName: dbName,
    collectionName: collectionName,
    ctx: ctx,
  } 
}


func (m *Model) InsertOne(doc interface{}) (*mongo.InsertOneResult, error) {
 
    collection := dbLogic.GetCollection(m.client, m.dbName, m.collectionName) 
     
    
    result, err := collection.InsertOne(m.ctx, doc)
    return result, err
}
 
/*
  insertMany is a user defined method, used to insert
  documents into collection returns result of
  InsertMany and error if any.
*/
func (m *Model) InsertMany(docs []interface{}) (*mongo.InsertManyResult, error) {
 
    // select database and collection ith Client.Database
    // method and Database.Collection method
    collection := dbLogic.GetCollection(m.client, m.dbName, m.collectionName)
     
    // InsertMany accept two argument of type Context
    // and of empty interface  
    result, err := collection.InsertMany(m.ctx, docs)
    return result, err
}
