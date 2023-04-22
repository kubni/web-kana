/* TODO:
- Domain driven design -> Separate API functions and game logic in separate packages -
*/

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"web_kana_v1/kana/kana_logic"
	"web_kana_v1/kana/tables"
	"web_kana_v1/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserData struct {
	CurrentPlayerObjectID primitive.ObjectID
	CurrentPlayerStringID string
}

type GameController struct {
	// data                TemplateData
	data                UserData
	model               *models.Model
	chosenAlphabetTable map[string][]string
}

func NewGameController(ctx context.Context, client *mongo.Client) *GameController {
	var gc GameController

	gc.model = models.NewModel(client, "testdb", "scoreboard3")
	gc.chosenAlphabetTable = make(map[string][]string)

	return &gc
}

// These should maybe go into API package
// However, they are still related to the gameplay
func CheckAnswer(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		UserAnswer             string `json:"userAnswer"`
		CorrectAnswerCharacter string `json:"correctAnswerCharacter"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Println(err)
	}

	// Prepare a response
	userAnswer := requestData.UserAnswer
	correctAnswerCharacter := requestData.CorrectAnswerCharacter
	isAnswerCorrect, correctAnswerRomaji := kana_logic.Check_answer(userAnswer, correctAnswerCharacter)

	responseData := struct {
		IsAnswerCorrect     bool   `json:"isAnswerCorrect"`
		CorrectAnswerRomaji string `json:"correctAnswerRomaji"`
	}{
		IsAnswerCorrect:     isAnswerCorrect,
		CorrectAnswerRomaji: correctAnswerRomaji,
	}

	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		log.Println(err)
		// TODO: Error handling (we can't stop the server)
	}
}

func GenerateHiraganaCharacter(w http.ResponseWriter, r *http.Request) {
	generatedCharacter := kana_logic.Play_all_gamemode(tables.Hiragana_table)

	// Encode the response (generatedCharacter) and send it to the frontend
	if err := json.NewEncoder(w).Encode(generatedCharacter); err != nil {
		log.Println(err)
		// TODO: Error handling (we can't stop the server)
	}
}

func GenerateKatakanaCharacter(w http.ResponseWriter, r *http.Request) {
	generatedCharacter := kana_logic.Play_all_gamemode(tables.Katakana_table)

	// Encode the response (generatedCharacter) and send it to the frontend
	if err := json.NewEncoder(w).Encode(generatedCharacter); err != nil {
		log.Println(err)
		// TODO: Error handling (we can't stop the server)
	}
}

// TODO: Decompose this function
func (gc *GameController) HandleUserData(w http.ResponseWriter, r *http.Request) {
	var userData struct {
		Username string `json:"username"`
		Score    int    `json:"score"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		log.Println(err)
	}

	fmt.Println("userData: ", userData)

	// TODO: Rename CheckIfUsernameAlreadyExists to doesUsernameAlreadyExists
	var isUsernameValid bool
	if userData.Username == "" || gc.model.CheckIfUsernameAlreadyExists(userData.Username) {
		fmt.Println("Username already exists: ", userData.Username)
		isUsernameValid = false
	} else {
		isUsernameValid = true
	}

	// Declare a mongoDB document
	document := bson.M{
		"ID":       "", //  At this point this is empty, but we populate it in the model with `bson` notation
		"Username": userData.Username,
		"Score":    userData.Score,
		"Rank":     0,
	}

	// Add the player to the database
	fmt.Println("Document: ", document)
	fmt.Println("Inserting the user into the db...")

	var insertInfo struct {
		IsInserted bool   `json:"isInserted"`
		Error      string `json:"error"`
		StringID   string `json:"stringID"`
	}

	if isUsernameValid {
		insertOneResult, err := gc.model.InsertOne(document)
		if err != nil {
			log.Println(err)
			insertInfo.IsInserted = false
			insertInfo.Error = err.Error()
		} else {
			insertInfo.IsInserted = true
			insertInfo.Error = ""

			fmt.Println("Insert result: ", insertOneResult)

			if id, ok := insertOneResult.InsertedID.(primitive.ObjectID); ok {
				gc.data.CurrentPlayerObjectID = id
				insertInfo.StringID = id.Hex()
			}
		}
	} else {
		insertInfo.IsInserted = false
		insertInfo.Error = "Username already exists"
	}

	// We need to actually send a json response about the insert action status
	if err := json.NewEncoder(w).Encode(insertInfo); err != nil {
		log.Println(err)
	}
}

func (gc *GameController) CalculatePlayerRank(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We are in CalculatePlayerRank")

	var userData struct {
		CurrentPlayerStringID string `json:"currentPlayerStringID"`
		CurrentPlayerScore    int    `json:"currentPlayerScore"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		log.Println(err)
	}

	fmt.Println("userData in CalculatePlayerRank: ", userData)

	// For GetAndSetPlayerRank we need an objectID which has the type primitive.ObjectID.
	// In order to convert our stringID to ObjectID, we use ObjectIDFromHex
	currentPlayerObjectID, err := primitive.ObjectIDFromHex(userData.CurrentPlayerStringID)
	if err != nil {
		log.Println(err)
	}

	currentPlayerRank := gc.model.GetAndSetPlayerRank(currentPlayerObjectID, userData.CurrentPlayerScore)
	fmt.Println("CurrentPlayerRank: ", currentPlayerRank)

	// Now we need to update other player ranks:
	gc.model.UpdateOtherRanks(currentPlayerObjectID, userData.CurrentPlayerScore)

	// TODO: Should we include an error field in the response that encodes err ?            !!!
	httpResponse := struct {
		CurrentPlayerRank int64 `json:"currentPlayerRank"`
	}{
		CurrentPlayerRank: currentPlayerRank,
	}

	if err := json.NewEncoder(w).Encode(httpResponse); err != nil {
		log.Println(err)
	}

}

func (gc *GameController) GetScoreboard(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		CurrentPage int `json:"currentPage"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Println("CurrentPageError: ", err)
	}

	// GetScoreboard returns []models.DocumentSchema which is the following:
	/*
		    type DocumentSchema struct {
			    ID       string `bson:"_id, omitempty"`
			    Username string
			    Score    int
			    Rank     int
		    }
	*/
	scoreboard, numOfPages := gc.model.GetScoreboard(requestData.CurrentPage)

	fmt.Println("Scoreboard: ", scoreboard)

	responseData := struct {
		Scoreboard []models.DocumentSchema `json:"scoreboard"`
		NumOfPages int                     `json:"numOfPages"`
	}{
		Scoreboard: scoreboard,
		NumOfPages: numOfPages,
	}

	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		log.Println(err)
	}
}
