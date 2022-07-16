package models 

import (
  "fmt"
  "context"
  "time"
  "log"
  "math"
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

func (m *Model) GetPlayerRank(playerID string) (int64, error) {
  collection := dbLogic.GetCollection(m.client, m.dbName, m.collectionName)
  ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)

  // To get the player rank, we count the number of players that exist before him.
  // https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/count/
  // Added Key: and Value: in order to get rid of "unkeyed values" warnings
  filter := bson.D{{Key: "ID", Value: bson.D{{Key: "$lt", Value: playerID}}}}
  position, err := collection.CountDocuments(ctx, filter)

  return position, err
}




// TODO: A better way than a global value?
type DocumentSchema struct{
  ID            string  `bson:"_id, omitempty"` // TODO: Explore the potential of bson notation
  Username      string  // Has to have the same name as the corresponding field in the database 
  Score         int     
}

// Pagination logic
func (m *Model) CalculateNumberOfPages(playersPerPage int) int {
  numberOfPlayers := dbLogic.CountDocuments(dbLogic.GetCollection(m.client, m.dbName, m.collectionName)) // TODO: Check if this works 
  
  numOfPages := math.Ceil(float64(numberOfPlayers)  / float64(playersPerPage))
   
  return int(numOfPages)
}


func (m *Model) GetScoreboard(currentPage int) ([]DocumentSchema, int)  {

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

  // Pagination logic 
  playersPerPage := 10
  numOfPages := m.CalculateNumberOfPages(playersPerPage)

  // Iterate through the results and add them into the previously declared slice
  i := 1
  j := 1
  for cursor.Next(context.Background()) {
    
    // We need to decode 10 players starting from the one that is at 10*currentPage so we skip the ones before it.
    // TODO: There has to be a better way of doing this 
    if j < playersPerPage*currentPage {
      j++
      continue 
    } 

    // We store only the desired number of players per page into our scoreboard 
    if i > playersPerPage {
      break 
    } 

    result := DocumentSchema{} 

    // Decode bson into our chosen Golang data structure
    err := cursor.Decode(&result)
    if err != nil {
      log.Fatal(err) 
    }
    //fmt.Println("result.ID: ", result.ID) // Program correctly maps the IDs to ID field of our result
    scoreboard = append(scoreboard, result)

    i++
  }

  if err := cursor.Err(); err != nil {
    log.Fatal(err)
  }

  if currentPage > numOfPages {
    // TODO: Instead of a println we should remove the next page button because there are no more next pages
    fmt.Println("There are no more pages....") 
    
  }

  //fmt.Println("Scoreboard: ", scoreboard)
  // In case no error occured
  return scoreboard, numOfPages
}
