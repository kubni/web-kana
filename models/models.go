package models

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"
	"web_kana_v1/dbLogic"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client         *mongo.Client
	dbName         string
	collectionName string
	ctx            context.Context
}

// Practically a constructor
func NewModel(client *mongo.Client, dbName string, collectionName string) *Model {
	return &Model{
		client:         client,
		dbName:         dbName,
		collectionName: collectionName,
	}
}

// TODO: Is this needed when we can use default collection.InsertOne and InsertMany?
func (m *Model) InsertOne(doc interface{}) (*mongo.InsertOneResult, error) {
	// We grab the specified collection
	collection := dbLogic.GetCollection(m.client, m.dbName, m.collectionName)

	// TODO: Cancel functions?
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func (m *Model) InsertMany(docs []interface{}) (*mongo.InsertManyResult, error) {
	collection := dbLogic.GetCollection(m.client, m.dbName, m.collectionName)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := collection.InsertMany(ctx, docs)
	return result, err
}

// TODO: A better way than a global value?
type DocumentSchema struct {
	// TODO: How does this bson annotation actually works?
	ID       string `bson:"_id, omitempty"`
	Username string // Has to have the same name as the corresponding field in the database
	Score    int
	Rank     int
}


func (m *Model) GetAndSetPlayerRank(currentPlayerObjectID primitive.ObjectID, currentPlayerScore int) int64 {
	collection := dbLogic.GetCollection(m.client, m.dbName, m.collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/count/
	// Added Key: and Value: in order to get rid of "unkeyed values" warnings

	// We calculate the player rank by counting the number of scores that are greater than his.
	filter := bson.D{{Key: "Score", Value: bson.D{{Key: "$gt", Value: currentPlayerScore}}}}
	position, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		panic(err)
	}
	// We add 1 since CountDocuments is 0-indexed
	position++

	// We need to set the player's rank in the actual database.
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "Rank", Value: position}}}}
	_, err = collection.UpdateByID(ctx, currentPlayerObjectID, update)
	if err != nil {
		panic(err)
	}

	return position
}

// We must update the ranks of the players that are now lower in rank compared to the currently added player
// We can do that by comparing their scores.
func (m *Model) UpdateOtherRanks(currentPlayerObjectID primitive.ObjectID, currentPlayerScore int) {
	collection := dbLogic.GetCollection(m.client, m.dbName, m.collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  
	// TODO: bson.D or M?
	filter := bson.M{"Score": bson.D{{Key: "$lte", Value: currentPlayerScore}}}
	update := bson.M{"$inc": bson.M{"Rank": 1}}
  _, err := collection.UpdateMany(ctx, filter, update)
  if err != nil {
    panic(err)
  } 
  // Fix the current player's rank:
  update = bson.M{"$inc": bson.M{"Rank": -1}}
  _, err = collection.UpdateByID(ctx, currentPlayerObjectID, update)
  if err != nil {
    panic(err)
  } 
}

// Pagination logic
func (m *Model) CalculateNumberOfPages(playersPerPage int) int {
	numberOfPlayers := dbLogic.CountAllDocuments(dbLogic.GetCollection(m.client, m.dbName, m.collectionName))

	numOfPages := math.Ceil(float64(numberOfPlayers) / float64(playersPerPage))

	return int(numOfPages)
}

func (m *Model) GetScoreboard(currentPage int) ([]DocumentSchema, int) {
	collection := dbLogic.GetCollection(m.client, m.dbName, m.collectionName)
	// Sort by score
	opts := options.Find().SetSort(bson.D{{Key: "Score", Value: -1}})

	// collection.Find will return a Cursor, which is basically a pointer to the set of documents
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
		// fmt.Println("result.ID: ", result.ID) // Program correctly maps the IDs to ID field of our result
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

	// fmt.Println("Scoreboard: ", scoreboard)
	// In case no error occured
	return scoreboard, numOfPages
}



// Index for username 
func (m *Model) CheckIfUsernameAlreadyExists(providedUsername string) bool {
	collection := dbLogic.GetCollection(m.client, m.dbName, m.collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

  filter := bson.M{"Username": bson.M{"$eq": providedUsername}}
  result := collection.FindOne(ctx, filter)
  fmt.Println("Println result.ERr() usernametest: ", result.Err().Error())
  
  if result.Err() == mongo.ErrNoDocuments {
    return false 
  } else {
    return true
  }
}
