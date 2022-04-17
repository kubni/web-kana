package main 


import (
  "fmt"
  "log"
  "net/http"
  "html/template"
  "web_kana_v1/hiragana"

  "github.com/gorilla/mux"
)

type TemplateData struct {
  PageTitle string 
  Character string
  Result    string
}


/* //Kako da prosledim tmp dole u handlefunc da ne bih imao anonimnu funkciju ?

func rootHandler(w http.ResponseWriter, r *http.Request) {
  data := TemplateData {
    PageTitle: "Web Kana", 
    HelloWorldTest: "Hello from a templated world!",
  }

  tmp.Execute(w, data)
}
*/

func main () {


  r := mux.NewRouter()
 
  // TODO: https://stackoverflow.com/questions/26211954/how-do-i-pass-arguments-to-my-handler
  tmp_main := template.Must(template.ParseFiles("./templates/layout_main.html"))

  r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    tmp_main.Execute(w, nil)  
  }).Methods("GET")


  tmp_game := template.Must(template.ParseFiles("./templates/layout_game.html"))
  r.HandleFunc("/game", func (w http.ResponseWriter, r *http.Request) {



    data := TemplateData {
        PageTitle: "Web Kana", 
        Character: hiragana.Play_all_gamemode(),
        Result: "", 
    }
    


    if err := r.ParseForm(); err != nil {
      fmt.Println("ParseForm() error: %v", err)
    }
   
    answer := r.FormValue("answer")

    if hiragana.Check_answer(answer, data.Character) {
        data.Result = "Correct answer!"
    } else {
        data.Result = "Wrong answer!"
    }

    tmp_game.Execute(w, data)
  }).Methods("POST")
 
  





  fmt.Println("Starting server at port 8000...")
  log.Fatal(http.ListenAndServe(":8000", r))
}
