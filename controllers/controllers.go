/* TODO:
4) Separate endpoints (/game, /game/finish, /game/scoreboard...)
5) CSS (divs, paddings, margins, layout of the page...)
*/

package controllers

import (
	"context"
	"fmt"
	"net/http"
	"web_kana_v1/kana/kana_logic"
	"web_kana_v1/kana/tables"
	"web_kana_v1/models"
	"web_kana_v1/templates"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TemplateData struct {
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

// Main page (selection) controller
func (gc *GameController) Selection(w http.ResponseWriter, r *http.Request) {
	// gc.data vs nil - no difference
	if err := templates.TmpMain.Execute(w, gc.data); err != nil {
		panic(err)
	}
}

// Game page (playground) controller
func (gc *GameController) Playground(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Parse the GET form:
		if err := r.ParseForm(); err != nil {
			fmt.Printf("ParseForm() error: %v", err)
		}

		// TODO: Move this into a separate function
		chosenAlphabet := r.FormValue("chosen-alphabet")
		if chosenAlphabet == "Hiragana" {
			gc.chosenAlphabetTable = tables.Hiragana_table
			gc.data.PageTitle = "ひらがな"
		} else {
			gc.chosenAlphabetTable = tables.Katakana_table
			gc.data.PageTitle = "カタカナ"
		}

		gc.data.Character = kana_logic.Play_all_gamemode(gc.chosenAlphabetTable)

		if err := templates.TmpGame.Execute(w, gc.data); err != nil {
			panic(err)
		}
	} else { // POST

		if err := r.ParseForm(); err != nil {
			fmt.Printf("ParseForm() error: %v", err)
		}

		isPlayAgainTrue := r.FormValue("isPlayAgainTrue")
		if isPlayAgainTrue == "true" {
			// TODO: Find a prettier way. The problem: We want to reset everything EXCEPT the PageTitle.
			// If we just comment that, it will still get the default value for string which is "" since we are creating a new
			// TemplateData "object"

			// gc.data = TemplateData{
			// 	//PageTitle:             "", // We don't want to reset the page title
			// 	Character:             "",
			// 	ResultMessage:         "",
			// 	CorrectAnswer:         "",
			// 	IsFinished:            "false",
			// 	CurrentPlayer:         "",
			// 	CurrentPlayerStringID: "",
			// 	CurrentPlayerRank:     0,
			// 	CurrentPlayerScore:    0,
			// 	IsUsernameValid:       "false",
			// 	DisplayScoreboard:     "false",
			// 	Scoreboard:            []models.DocumentSchema{},
			// 	CurrentPage:           0,
			// 	NumOfPages:            1,
			// 	MessageForUser:        "",
			// 	IsPlayAgainTrue:       "true", // We set the IsPlayAgainTrue to true here
			// }

			// TODO: Find a better way. This way is awful, but we can keep the PageTitle
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

		// TODO: Is this okay? r.FormValue("finish-value") doesn't have the TRUE value if we don't click on the Finish button,
		// so if we check with that, we will only pass the condition once.
		// Therefore, this is needed because IsFinished will always be true once set because its in a global struct variable
		if gc.data.IsFinished == "true" {
			// Parse the username and check if it already exists in the database:
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

			// TODO: Write this better
			/*
			   First time that we arrive here after the Play Again button,
			   we don't want the program to compare answer and gc.data.Character because
			   those values got reset to "" so it will be true always and increment the score.

			   We do the answer checking only if IsPlayAgain is false, and if it isn't we set it in this if's else.
			*/
			if gc.data.IsPlayAgainTrue == "false" {
				// Parse the answer
				answer := r.FormValue("answer")

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

			// TODO: Remember why you put this here and why it doesn't work if its above the answer = ... line
			gc.data.Character = kana_logic.Play_all_gamemode(gc.chosenAlphabetTable)

		}
		if err := templates.TmpGame.Execute(w, gc.data); err != nil {
			panic(err)
		}
	}
}
