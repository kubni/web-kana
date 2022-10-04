// TODO: Error checking inside function definitions in models.go or in controllers.go?

/* TODO:
1) Make usernames unique
2) Play again button
*/

/* FIXME:
1) For example, if there are 5 players in the scoreboard, and 2, 3. have the same rank (for example 2), players 4 and 5 will have ranks
   4 and 5 instead of 3 and 4.
      // Design choice:
          *) Ranks: 1, 2, 2, 3, 4
              // I can just do $inc on LTE and then $dec only the current player's rank
          *) Ranks: 1, 2, 3, 4, 5 (even though 2. and 3. player have the same score)

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

// TODO: Make this generic
type OnChangeFunction func(string) bool

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

	FunctionTest OnChangeFunction
}

type GameController struct {
	data                TemplateData
	model               *models.Model
	chosenAlphabetTable map[string][]string
}

func NewGameController(ctx context.Context, client *mongo.Client) *GameController {
	var gc GameController

	gc.data = TemplateData{
		PageTitle:     "",
		Character:     "",
		ResultMessage: "",
		CorrectAnswer: "",
		IsFinished:    "false",
		CurrentPlayer: "",
		// CurrentPlayerObjectID: nil,
		CurrentPlayerStringID: "",
		CurrentPlayerRank:     0,
		CurrentPlayerScore:    0,
		IsUsernameValid:       "false",
		DisplayScoreboard:     "false",
		Scoreboard:            []models.DocumentSchema{},
		CurrentPage:           0,
		NumOfPages:            1,
		MessageForUser:        "",
	}

	gc.model = models.NewModel(client, "testdb", "scoreboard3")

	gc.chosenAlphabetTable = make(map[string][]string)

	return &gc
}

// Main page (selection) controller
func (gc *GameController) Selection(w http.ResponseWriter, r *http.Request) {
	// TODO: Why doesn't this reset the data after we go back to main page with the button?
	/*
		  gc.data = TemplateData{
			PageTitle:     "",
			Character:     "",
			ResultMessage: "",
			CorrectAnswer: "",
			IsFinished:    "false",
			CurrentPlayer: "",
			CurrentPlayerStringID: "",
			CurrentPlayerRank:     0,
			CurrentPlayerScore:    0,
			DisplayScoreboard:     "false",
			Scoreboard:            []models.DocumentSchema{},
			CurrentPage:           0,
			NumOfPages:            1,
			MessageForUser:        "",
		}

		  }
	*/

	if err := templates.TmpMain.Execute(w, nil); err != nil {
		panic(err)
	}
}

// Game page (playground) controller
func (gc *GameController) Playground(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
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

	} else {

		if err := r.ParseForm(); err != nil {
			fmt.Printf("ParseForm() error: %v", err)
		}

		// TODO: Get rid of the "double"  ifs for same thing.
		// Javascript frontend validation?

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
			}

			// gc.data.CurrentPlayer will not be empty if the entered username is valid (passes the check above)
			if gc.data.CurrentPlayer != "" {
				// TODO: Is this field even needed?
				// Maybe i can use it in template if checks. !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
				gc.data.IsUsernameValid = "true"

				if r.FormValue("isNextPageClicked") == "true" {
					fmt.Println("CURRENTPAGE IN CONTROLLER: ", gc.data.CurrentPage)
					gc.data.CurrentPage++
				} else if r.FormValue("isPreviousPageClicked") == "true" {
					gc.data.CurrentPage--
				} else { // We don't want to insert same user into the db each time we press "Next Page" button

					var document interface{}
					// As per the official documentation, bson.M should be used if the order of the elements in the document doesn't matter
					document = bson.M{
						// TODO: Figure out how does this actually have the value when it doesn't show up in mongosh when I check the document
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
						fmt.Printf("InsertOne() error: %v", err)
					} else {
						fmt.Println("Insert result: ", insertOneResult)

						// Decode the insertOneResult into a string
						if id, ok := insertOneResult.InsertedID.(primitive.ObjectID); ok {
							gc.data.CurrentPlayerObjectID = id
							gc.data.CurrentPlayerStringID = id.Hex()
						}
					}
				}
				// TODO: We have to get the rank here, because we are setting the gc.data.CurrentPlayerID after we insert the player into the db

				// Get and set the player rank
				gc.data.CurrentPlayerRank = gc.model.GetAndSetPlayerRank(gc.data.CurrentPlayerObjectID, gc.data.CurrentPlayerScore)

				// Update the ranks of other players that are below the current player.
				// FIXME: This doesn't work as it should.
				gc.model.UpdateOtherRanks(gc.data.CurrentPlayerObjectID, gc.data.CurrentPlayerScore)

				// TODO: Change this to bool
				gc.data.DisplayScoreboard = "true"

				fmt.Println("Zovemo GetScoreboard sa parametrom gc.data.currentpage: ", gc.data.CurrentPage)
				gc.data.Scoreboard, gc.data.NumOfPages = gc.model.GetScoreboard(gc.data.CurrentPage)

				fmt.Println("Gc.data.Scoreboard: ", gc.data.Scoreboard)

				fmt.Println("gc.Data.numofpages: ", gc.data.NumOfPages)
			}
		} else {
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

			gc.data.Character = kana_logic.Play_all_gamemode(gc.chosenAlphabetTable)

		}
		if err := templates.TmpGame.Execute(w, gc.data); err != nil {
			panic(err)
		}

	}
}

// TODO: Weird things happen if the page goes to the finished screen and the user goes back by using browser back-arrow

/*
   FIXME:
   This successfully prevents the insertion of a username that already exists and returns the player to the same finish page where
   he can choose another name again.
   Also, if he enters another name that doesn't already exists in the db, he will be properly added into the db.

   Problems:
       1) He won't get highlighted or even shown at all in the scoreboard.
       2) Also, if Next Page button is clicked, we enter that if with the <p> in layout_game.html

      3) There is a problem with pagination. UpdateOtherRanks breaks it probably, since those that get their ranks updated past the page
         on which they were, don't get moved to the next page.
            // Why is this broken when we are making sure to call GetScoreboard AFTER we call UpdateOtherRanks



   // Something probably happens with the forms, also with the TemplateData values when the reset happens.

   // gc.data.DisplayScoreboard = "true" and the surrounding code should probably be put somewhere else.


*/
