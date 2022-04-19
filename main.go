package main

import (
	"fmt"
	"log"
        "net/http"
	"web_kana_v1/routes"

        "github.com/gorilla/mux"
)


func main () {
  r := mux.NewRouter()

  // For serving the static files 
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))) 

  routes.InitRoutes(r)

  fmt.Println("Starting server at port 8000...")
  log.Fatal(http.ListenAndServe(":8000", r))
}
