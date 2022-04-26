 package main

import (
	"fmt"
	"log"
        "net/http"
	"web_kana_v1/routes"
        "web_kana_v1/dbLogic"
        
        "github.com/gorilla/mux"
)


func main () {

  
  // TODO:  not needed?
  dbClient, ctx := dbLogic.InitializeDatabaseConnection() 

  r := mux.NewRouter()

  // For serving the static files 
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))) 
  
  // Initalize the routes
  routes.InitRoutes(r)

  // Start the web server
  fmt.Println("Starting server at port 8000...")
  log.Fatal(http.ListenAndServe(":8000", r))
}
