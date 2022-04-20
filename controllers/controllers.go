package controllers 

import (
  "fmt"
  "net/http"
  "web_kana_v1/templates"
  "web_kana_v1/kana/tables"
  "web_kana_v1/kana/kana_logic"
)


type TemplateData struct {
  PageTitle string 
  Character string
  Result    string
  CorrectAnswer string
}


func IndexController (w http.ResponseWriter, r *http.Request) {
  templates.TmpMain.Execute(w, nil)  
}

var chosenAlphabetTable map[string][]string // TODO: Bad practice?
var data = TemplateData {
        PageTitle: "", 
        Character: "",
        Result: "",
        CorrectAnswer: "",
 }

func GameController (w http.ResponseWriter, r *http.Request) {
  
  if r.Method == "GET" {
    if err := r.ParseForm(); err != nil {
      fmt.Printf("ParseForm() error: %v", err)
    }


    fmt.Println("r.Form: \n", r.Form)




    chosenAlphabet := r.FormValue("chosen-alphabet")
    if chosenAlphabet == "Hiragana" {
      chosenAlphabetTable = tables.Hiragana_table
      data.PageTitle = "ひらがな"
    } else {
      chosenAlphabetTable = tables.Katakana_table
      data.PageTitle = "カタカナ"
    }

    data.Character = kana_logic.Play_all_gamemode(chosenAlphabetTable) 
    
    templates.TmpGame.Execute(w, data)
  } else {
    if err := r.ParseForm(); err != nil {
      fmt.Printf("ParseForm() error: %v", err)
    }
   
    answer := r.FormValue("answer")

    if kana_logic.Check_answer(answer, data.Character) {
      data.Result = "Correct answer!"
      data.CorrectAnswer = ""
    } else {
      data.CorrectAnswer = tables.Romaji_table[data.Character]
      data.Result = fmt.Sprintf("Wrong, the right answer was ") 
    }
    
    data.Character = kana_logic.Play_all_gamemode(chosenAlphabetTable) 
    
    templates.TmpGame.Execute(w, data)  
  }
}

