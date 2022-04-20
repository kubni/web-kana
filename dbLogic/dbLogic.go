package dbLogic


import (
    "context"
    "fmt"
    "time"
 
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
            
    /* 
    The WithCancel, WithDeadline, and WithTimeout functions take a Context (the parent)
    and return a derived Context (the child) and a CancelFunc.
    Calling the CancelFunc cancels the child and its children,
    removes the parent's reference to the child, and stops any associated timers.
    */
    defer cancel()
    
    // client.Disconnect will close the database connection
    // TODO: Why is this deferred if we defer the close() function which calls these 2 functions in the main()
    defer func(){
     
        // client.Disconnect method also has deadline.
        // returns error if any,
        if err := client.Disconnect(ctx); err != nil{
            panic(err)
        }
    }()
}
 
func ConnectTo(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
  /*
  We set the deadline for process operations to 30 seconds 
  Background() returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. 
  It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests. 
  */
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  
     
  /* 
  mongo.Connect returns mongo.Client which will be used for further operations with the database. 
  ApplyURI parses the given URI and sets options accordingly.
  */
  client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

  return client, ctx, cancel, err
}
 
// This method is used to ping the mongoDB to check the connection status and return errors if there are any.
func Ping(client *mongo.Client, ctx context.Context, uri string) error {
 
    // Deadline of the Ping method will be determined by cxt (context of the process)
    // https://www.mongodb.com/docs/manual/core/read-preference/
    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        return err
    }
    fmt.Printf("Connected successfully to %v\n", uri)
    return nil
}
