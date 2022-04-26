package controllers 

import (
  "fmt"
  "net/http"
  "web_kana_v1/templates"
  "web_kana_v1/kana/tables"
  "web_kana_v1/kana/kana_logic"
  "web_kana_v1/models"

  "go.mongodb.org/mongo-driver/bson"
  
  // Test 
  "go.mongodb.org/mongo-driver/mongo"
  "context"
)


type TemplateData struct {
  PageTitle     string 
  Character     string
  ResultMessage string
  CorrectAnswer string
  IsFinished    string
  TotalScore    int
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
    PageTitle: "", 
    Character: "",
    ResultMessage: "",
    CorrectAnswer: "",
    IsFinished: "false",
    TotalScore: 0,
  }

  gc.model = models.NewModel(ctx, client, "testdb", "user")
 
  gc.chosenAlphabetTable = make(map[string][]string)
  
  return &gc
}

// Main page (selection) controller 

// TODO: Templates field in the controller ? 
func (gc *GameController) Selection(w http.ResponseWriter, r *http.Request) {
  templates.TmpMain.Execute(w, nil)  
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
    
    templates.TmpGame.Execute(w, gc.data)

  } else {
    if err := r.ParseForm(); err != nil {
      fmt.Printf("ParseForm() error: %v", err)
    }
  
    // Check if the finish button has been clicked, if it was, we don't check the answer and the game stops. 
    if r.FormValue("finish-value") == "true" {
      gc.data.IsFinished = "true"
    }

    // TODO: Is this okay? r.FormValue("finish-value") doesn't have a TRUE value if we don't click on the Finish button, so if we check with that, we will only pass the condition once. 
    // Therefore, this is needed because IsFinished will always be true once set because its in a global struct variable 
    if gc.data.IsFinished == "true" {
      // Parse the username 
      // TODO: The first time we pass this condition (precisely when the Finish button is clicked) we will have username = "", but this could actually work, we only need to check if its "" before adding it to the db.
      // There must be a better way.
      username := r.FormValue("username")
     
      // TODO: Check if it exists in the db, if it exists an update should be done?
      // How to prevent someone else from updating other's score? We would need authentication.


      // If the username isn't empty (which only happens the first time, read the comments up and fix it)
      if username != "" {
        var document interface{}

        // As per the official documentation, bson.M should be used if the order of teh elements in the document doesn't matter
        document = bson.M {
          "Username": username,
          "Score": gc.data.TotalScore,
        }

        // Add the player to the database
        fmt.Println("document: ", document)
        fmt.Println("Inserting the user into the db...!") // TODO: Not working
        insertOneResult, err := gc.model.InsertOne(document)
        if err != nil {
          fmt.Printf("InsertOne() error: %v", err)
        } else {
          fmt.Println("Insert result: ", insertOneResult)
        }
      }
      fmt.Println("Username: ", username)
    } else {

      // Parse the answer 
      answer := r.FormValue("answer")
      fmt.Println("Finish value: ", gc.data.IsFinished)
      // Check if the answer is correct 
      if kana_logic.Check_answer(answer, gc.data.Character) {
        gc.data.ResultMessage = "Correct answer!"
        gc.data.CorrectAnswer = ""
        gc.data.TotalScore++
      } else {
        gc.data.CorrectAnswer = tables.Romaji_table[gc.data.Character]
        gc.data.ResultMessage = fmt.Sprintf("Wrong, the right answer was ") 
        gc.data.TotalScore--
      }

      gc.data.Character = kana_logic.Play_all_gamemode(gc.chosenAlphabetTable) 
    } 
    templates.TmpGame.Execute(w, gc.data)  
  }
}

// TODO: Weird things happen if the page goes to the finished screen and the user goes back by using browser back-arrow 

