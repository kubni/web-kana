package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"web_kana_v1/kana/tables"
	"web_kana_v1/kana/kana_logic"
	"github.com/gorilla/mux"
)

type TemplateData struct {
  PageTitle string 
  Character string
  Result    string
  CorrectAnswer string
}

func main () {

  r := mux.NewRouter()

  // For serving the static files 
  r.PathPrefix("/static/stylesheets").Handler(http.StripPrefix("/static/stylesheets", http.FileServer(http.Dir("./static/stylesheets")))) 

  // TODO: https://stackoverflow.com/questions/26211954/how-do-i-pass-arguments-to-my-handler
  tmp_main := template.Must(template.ParseFiles("./templates/layout_main.html"))

  r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    tmp_main.Execute(w, nil)  
  }).Methods("GET")




  var chosen_alphabet_table map[string][]string


  tmp_game := template.Must(template.ParseFiles("./templates/layout_game.html"))
  data := TemplateData {
        PageTitle: "", 
        Character: "",
        Result: "",
        CorrectAnswer: "",
  }

  r.HandleFunc("/game", func (w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
      r.ParseForm()

      chosen_alphabet := r.FormValue("chosen-alphabet")
      if chosen_alphabet == "Hiragana" {
        chosen_alphabet_table = tables.Hiragana_table
        data.PageTitle = "ひらがな"
      } else {
        chosen_alphabet_table = tables.Katakana_table
        data.PageTitle = "カタカナ"
      }
      
      data.Character = kana_logic.Play_all_gamemode(chosen_alphabet_table) 

      tmp_game.Execute(w, data)
    } else {

      if err := r.ParseForm(); err != nil {
        fmt.Printf("ParseForm() error: %v", err)
      }
     
      answer := r.FormValue("answer")

      if kana_logic.Check_answer(answer, data.Character) {
        data.Result = "Correct answer!"
      } else {
        data.CorrectAnswer = tables.Romaji_table[data.Character]
        data.Result = fmt.Sprintf("Wrong, the right answer  was ") 
      }
      
      data.Character = kana_logic.Play_all_gamemode(chosen_alphabet_table) 
      
      tmp_game.Execute(w, data)  
    }
  })

  fmt.Println("Starting server at port 8000...")
  log.Fatal(http.ListenAndServe(":8000", r))
}
