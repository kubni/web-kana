package controllers 

import (
  "fmt"
  "net/http"
  // "time"
  "web_kana_v1/templates"
  "web_kana_v1/kana/tables"
  "web_kana_v1/kana/kana_logic"
  "web_kana_v1/models"

  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
  // Test 
  "go.mongodb.org/mongo-driver/mongo"

  "context"
)


type TemplateData struct {
  PageTitle             string 
  Character             string
  ResultMessage         string
  CorrectAnswer         string
  IsFinished            string
  CurrentPlayer         string
  CurrentPlayerID       string 
  CurrentRank           int64 
  TotalScore            int
  DisplayScoreboard     string
  Scoreboard            []models.DocumentSchema
  CurrentPage           int
  NumOfPages            int
  MessageForUser        string
}


type GameController struct {
  data TemplateData
  model *models.Model
  chosenAlphabetTable map[string][]string
}

// TODO: Do we need to return a *? 
func NewGameController(ctx context.Context, client *mongo.Client)  *GameController {
  var gc GameController

  gc.data = TemplateData {
    PageTitle:          "", 
    Character:          "",
    ResultMessage:      "",
    CorrectAnswer:      "",
    IsFinished:         "false",
    CurrentPlayer:      "",
    CurrentPlayerID:    "",
    CurrentRank:        0, // TODO: Implement this again 
    TotalScore:         0,
    DisplayScoreboard:  "false",
    Scoreboard:         []models.DocumentSchema{}, 
    CurrentPage:        0, 
    NumOfPages:         1,
    MessageForUser:     "",
  }

  gc.model = models.NewModel(client, "testdb", "scoreboard2")
 
  gc.chosenAlphabetTable = make(map[string][]string)
  
  return &gc
}

// Main page (selection) controller 
func (gc *GameController) Selection(w http.ResponseWriter, r *http.Request) {

  // TODO: Why doesn't this reset the data?
  /*
    gc.data = TemplateData {
    PageTitle: "", 
    Character: "",
    ResultMessage: "",
    CorrectAnswer: "",
    IsFinished: "false",
    CurrentPlayer: "",
    TotalScore: 0,
    DisplayScoreboard: "false",
    Scoreboard: []models.DocumentSchema{}, 
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
   
    // Check if the finish button has been clicked, if it was, we don't check the answer and the game stops. 
    if r.FormValue("isFinished") == "true" {
      gc.data.IsFinished = "true"
    }
 

    // TODO: Is this okay? r.FormValue("finish-value") doesn't have a TRUE value if we don't click on the Finish button,
    // so if we check with that, we will only pass the condition once. 
    // Therefore, this is needed because IsFinished will always be true once set because its in a global struct variable 
    if gc.data.IsFinished == "true" {
      // Parse the username 
      if r.FormValue("username") != "" {
        gc.data.CurrentPlayer = r.FormValue("username")
      }

      // If the username isn't empty (which only happens the first time, read the comments above and fix it)
      if gc.data.CurrentPlayer != "" {

        if r.FormValue("isNextPageClicked") == "true" {
          gc.data.CurrentPage++

        } else if r.FormValue("isPreviousPageClicked") == "true" {
          gc.data.CurrentPage-- 

        } else { // We don't want to insert same user into the db each time we press "Next Page" button

          // Get the player rank 
          rank, err := gc.model.GetPlayerRank(gc.data.CurrentPlayerID)
          if err != nil {
            fmt.Printf("GetPlayerRank() error: %v", err)
          } else {
            gc.data.CurrentRank = rank
            fmt.Println("gc.data.CurrentRank = ", gc.data.CurrentRank) // FIXME: Always 0
          }
          
          var document interface{}
          // As per the official documentation, bson.M should be used if the order of the elements in the document doesn't matter
          document = bson.M {
            "ID":       gc.data.CurrentPlayerID, // At this point this is empty, but we populate it in the model with `bson ...` annotation  
            "Username": gc.data.CurrentPlayer,                  
            "Score":    gc.data.TotalScore,
            "Rank":     gc.data.CurrentRank,
          }

          // Add the player to the database
          fmt.Println("document: ", document)
          fmt.Println("Inserting the user into the db...") 

          insertOneResult, err := gc.model.InsertOne(document)
          if err != nil {
            fmt.Printf("InsertOne() error: %v", err)
          } else {
            fmt.Println("Insert result: ", insertOneResult)

            // Decode the insertOneResult into a string
            if id, ok := insertOneResult.InsertedID.(primitive.ObjectID); ok {
              gc.data.CurrentPlayerID = id.Hex()
            }

          
            
          }
        } 
        
        // testRank := gc.model.Find()


        // TODO: Change this to bool
        gc.data.DisplayScoreboard = "true"
        
        gc.data.Scoreboard, gc.data.NumOfPages = gc.model.GetScoreboard(gc.data.CurrentPage)
      }
    } else {

      // Parse the answer 
      answer := r.FormValue("answer")

      // Check if the answer is correct 
      if kana_logic.Check_answer(answer, gc.data.Character) {
        gc.data.ResultMessage = "Correct answer!"
        gc.data.CorrectAnswer = ""
        gc.data.TotalScore++
      } else {
        gc.data.CorrectAnswer = tables.Romaji_table[gc.data.Character]
        gc.data.ResultMessage = fmt.Sprintf("Wrong, the right answer was ") 
        
        if gc.data.TotalScore > 0 {
          gc.data.TotalScore--
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

