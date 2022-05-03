package dbLogic


import (
    "context"
    "fmt"
    "time"
 
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
    "go.mongodb.org/mongo-driver/bson"
)

var uri = "mongodb://localhost:27017"

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
            
    /* 
      Calling the CancelFunc cancels the child and its children,
      removes the parent's reference to the child, and stops any associated timers.
    */
    defer cancel()
    
    // client.Disconnect will close the database connection
    defer func(){
     
        // client.Disconnect method also has deadline.
        // returns error if any
        if err := client.Disconnect(ctx); err != nil{
            panic(err)
        }
    }()
}
 
func ConnectTo(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
  /*
    We set the deadline for process operations to 10 seconds 
    Background() returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. 
  */
  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  /* 
    mongo.Connect returns mongo.Client which will be used for further operations with the database. 
    ApplyURI parses the given URI and sets options accordingly.
  */
  client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

  return client, ctx, cancel, err
}
 
// This method is used to ping the mongoDB to check the connection status and return errors if there are any.
func Ping(client *mongo.Client, ctx context.Context) error {
 
    // Deadline of the Ping method will be determined by cxt (context of the process)
    // https://www.mongodb.com/docs/manual/core/read-preference/
    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        return err
    }
    fmt.Printf("Connected successfully to %v\n", uri)
    return nil
}

func GetCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
    collection := client.Database(dbName).Collection(collectionName)
    return collection
}

func CountDocuments(collection *mongo.Collection) int64 {
  // Empty filter so it returns total count 
  filter := bson.D{}
  count, err := collection.CountDocuments(context.TODO(), filter)
  if err != nil {
    panic(err)
  }

  return count
}
func InitializeDatabaseConnection() (client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

  // Connect to the database
  client, ctx, cancel, err := ConnectTo(uri)
  if err != nil {
      panic(err)
  }
   
  return client, ctx, cancel
}
