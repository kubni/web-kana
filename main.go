 package main

import (
	"fmt"
	"log"
        "net/http"
	"web_kana_v1/routes"
        "web_kana_v1/dbLogic"
        
        "github.com/gorilla/mux"
)




// TODO:
/*
  1.Input to stay focused,
  2. Ranking table in place of username input field after the username has been saved

  3. Some kind of an update system
    // How to prevent someone with the same username from changing the other's score?
    // Maybe by using their unique ID as a pseudo-password?
*/





func main () {

  
  client, ctx, cancel := dbLogic.InitializeDatabaseConnection() 
  defer dbLogic.Close(client, ctx, cancel)
  dbLogic.Ping(client, ctx) 

  r := mux.NewRouter()

  // For serving the static files 
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))) 
  
  // Initalize the routes
  routes.InitRoutes(r, client, ctx)

  // Start the web server
  fmt.Println("Starting server at port 8000...")
  log.Fatal(http.ListenAndServe(":8000", r))
}
