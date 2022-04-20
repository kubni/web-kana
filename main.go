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

  // Database URI 
  uri := "mongodb://localhost:27017"

  // Connect to the database
  client, ctx, cancel, err := dbLogic.ConnectTo(uri)
  if err != nil {
      panic(err)
  }
   
  // Release resource when the main function is returned.
  defer dbLogic.Close(client, ctx, cancel)
   
  // Ping the database 
  dbLogic.Ping(client, ctx, uri)


//  var document interface{}
     
 // document = bson.D{
 //   {Key: "testIme",    Value: "Nikola"},
 //   {Key: "testPrez",   Value: "Kuburovic"},
 //   {Key: "science",    Value: 90},
 //   {Key: "computer",   Value: 95},
 // }

  r := mux.NewRouter()

  // For serving the static files 
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))) 
  
  // Initalize the routes
  routes.InitRoutes(r)

  // Start the web server
  fmt.Println("Starting server at port 8000...")
  log.Fatal(http.ListenAndServe(":8000", r))
}
