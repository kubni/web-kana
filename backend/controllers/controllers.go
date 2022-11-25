package controllers

import (
	"context"
	"fmt"
	"net/http"
  "encoding/json"
	"web_kana_v1/kana/kana_logic"
	"web_kana_v1/kana/tables"
	"web_kana_v1/models"
	"web_kana_v1/templates"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	/*  	"github.com/gorilla/schema" */)

// var decoder = schema.NewDecoder()

type AlphabetData struct {
  Name string
}

type TemplateData struct {
	ChosenAlphabet        string
	PageTitle             string
	Character             string
	ResultMessage         string
	CorrectAnswer         string
	IsFinished            string
	CurrentPlayer         string
	CurrentPlayerObjectID primitive.ObjectID
	CurrentPlayerStringID string // For comparison in layout_game.html
	CurrentPlayerRank     int64
	CurrentPlayerScore    int
	IsUsernameValid       string
	DisplayScoreboard     string
	Scoreboard            []models.DocumentSchema
	CurrentPage           int
	NumOfPages            int
	MessageForUser        string
	IsPlayAgainTrue       string
}

type GameController struct {
	data                TemplateData
	model               *models.Model
	chosenAlphabetTable map[string][]string
}

func NewGameController(ctx context.Context, client *mongo.Client) *GameController {
	var gc GameController

	gc.data = TemplateData{
		ChosenAlphabet:        "",
		PageTitle:             "",
		Character:             "",
		ResultMessage:         "",
		CorrectAnswer:         "",
		IsFinished:            "false",
		CurrentPlayer:         "",
		CurrentPlayerStringID: "",
		CurrentPlayerRank:     0,
		CurrentPlayerScore:    0,
		IsUsernameValid:       "false",
		DisplayScoreboard:     "false",
		Scoreboard:            []models.DocumentSchema{},
		CurrentPage:           0,
		NumOfPages:            1,
		MessageForUser:        "",
		IsPlayAgainTrue:       "false",
	}

	gc.model = models.NewModel(client, "testdb", "scoreboard3")

	gc.chosenAlphabetTable = make(map[string][]string)

	return &gc
}

//Main page (selection) controller
func (gc *GameController) Selection(w http.ResponseWriter, r *http.Request) {
	// gc.data vs nil - no difference
  fmt.Println("WE ARE IN SELECTION")	
}





// Game page (playground) controller
func (gc *GameController) Playground(w http.ResponseWriter, r *http.Request) {
  fmt.Println("We are in the Playground!")
  /*
    For other HTTP methods, or when the Content-Type is not
    application/x-www-form-urlencoded, the request Body is not read, and
    r.PostForm is initialized to a non nil, empty value.

  TODO: Does this mean we have to specify content-type application/x-www-form-urlencoded
        instead of application/json in frontend/MainReactForm
*/	



// if err := r.ParseForm(); err != nil {
// 		fmt.Printf("ParseForm() error: %v", err)
// 	}

  // var body interface {}
  // if err := json.NewDecoder(r.Body).Decode()
   

  testReactForm := r.FormValue("chosenAlphabet")
  fmt.Println("Test react form: ", testReactForm)

  return

	chosenAlphabet := r.FormValue("chosen-alphabet")
	if chosenAlphabet == "Hiragana" {
		gc.chosenAlphabetTable = tables.Hiragana_table
		gc.data.ChosenAlphabet = "hiragana"
	}
	if chosenAlphabet == "Katakana" {
		gc.chosenAlphabetTable = tables.Katakana_table
		gc.data.ChosenAlphabet = "katakana"
	}

	if gc.data.ChosenAlphabet == "hiragana" {
		gc.data.PageTitle = "ひらがな"
	} else {
		gc.data.PageTitle = "カタカナ"
	}

	isPlayAgainTrue := r.FormValue("isPlayAgainTrue")
	if isPlayAgainTrue == "true" {
		gc.data.ChosenAlphabet = ""
		gc.data.Character = ""
		gc.data.ResultMessage = ""
		gc.data.CorrectAnswer = ""
		gc.data.IsFinished = "false"
		gc.data.CurrentPlayer = ""
		gc.data.CurrentPlayerStringID = ""
		gc.data.CurrentPlayerRank = 0
		gc.data.CurrentPlayerScore = 0
		gc.data.IsUsernameValid = "false"
		gc.data.DisplayScoreboard = "false"
		gc.data.Scoreboard = []models.DocumentSchema{}
		gc.data.CurrentPage = 0
		gc.data.NumOfPages = 1
		gc.data.MessageForUser = ""
		gc.data.IsPlayAgainTrue = "true" // We set the IsPlayAgainTrue to true here
	}

	// Check if the finish button has been clicked, if it was, we don't check the answer and the game stops.
	if r.FormValue("isFinished") == "true" {
		gc.data.IsFinished = "true"
	}

	if gc.data.IsFinished == "true" {
		if r.FormValue("username") != "" && !gc.model.CheckIfUsernameAlreadyExists(r.FormValue("username")) {
			gc.data.CurrentPlayer = r.FormValue("username")
		} else {
			// Potential problem: We enter here every time we go to a next or previous table page. This isn't a problem for now.
			fmt.Println("The username you entered isn't valid. Please enter another one.")
			gc.data.IsUsernameValid = "false"
		}

		// gc.data.CurrentPlayer will not be empty if the entered username is valid (passes the check above)
		if gc.data.CurrentPlayer != "" {
			gc.data.IsUsernameValid = "true"

			if r.FormValue("isNextPageClicked") == "true" {
				gc.data.CurrentPage++
			} else if r.FormValue("isPreviousPageClicked") == "true" {
				gc.data.CurrentPage--
			} else { // We don't want to insert same user into the db each time we press "Next Page" button
				var document interface{}
				// As per the official documentation, bson.M should be used if the order of the elements in the document doesn't matter
				document = bson.M{
					"ID":       gc.data.CurrentPlayerStringID, //  At this point this is empty, but we populate it in the model with `bson` notation
					"Username": gc.data.CurrentPlayer,
					"Score":    gc.data.CurrentPlayerScore,
					"Rank":     gc.data.CurrentPlayerRank,
				}

				// Add the player to the database
				fmt.Println("Document: ", document)
				fmt.Println("Inserting the user into the db...")

				insertOneResult, err := gc.model.InsertOne(document)
				if err != nil {
					panic(err)
				} else {
					fmt.Println("Insert result: ", insertOneResult)

					// Decode the insertOneResult into a string
					if id, ok := insertOneResult.InsertedID.(primitive.ObjectID); ok {
						gc.data.CurrentPlayerObjectID = id
						gc.data.CurrentPlayerStringID = id.Hex()
					}
				}
			}
			// Get and set the player rank
			gc.data.CurrentPlayerRank = gc.model.GetAndSetPlayerRank(gc.data.CurrentPlayerObjectID, gc.data.CurrentPlayerScore)

			// Update the ranks of other players that are below the current player.
			gc.model.UpdateOtherRanks(gc.data.CurrentPlayerObjectID, gc.data.CurrentPlayerScore)

			gc.data.DisplayScoreboard = "true"
			gc.data.Scoreboard, gc.data.NumOfPages = gc.model.GetScoreboard(gc.data.CurrentPage)
		}
	} else {
		if gc.data.IsPlayAgainTrue == "false" {
			// Parse the answer
			answer := r.FormValue("answer")

      /* 
         If the character that we need to guess is empty, skip the check, 
         or in other words, if its not empty, do the check:
      */
      // This happens the first time we come here only!
      if gc.data.Character != "" { 
        // Check if the answer is correct
        if kana_logic.Check_answer(answer, gc.data.Character) {
          gc.data.ResultMessage = "Correct answer!"
          gc.data.CorrectAnswer = ""
          gc.data.CurrentPlayerScore++
        } else {
          gc.data.CorrectAnswer = tables.Romaji_table[gc.data.Character]
          gc.data.ResultMessage = fmt.Sprintf("Wrong, the right answer was ")

          if gc.data.CurrentPlayerScore > 0 {
            gc.data.CurrentPlayerScore--
          }
        }
      } else {
        gc.data.IsPlayAgainTrue = "false"
      }
    }


    gc.data.Character = kana_logic.Play_all_gamemode(gc.chosenAlphabetTable)

	}
	if err := templates.TmpGame.Execute(w, gc.data); err != nil {
		panic(err)
	}
}
