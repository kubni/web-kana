package routes 

import (
  "context"
  "web_kana_v1/controllers"

  "go.mongodb.org/mongo-driver/mongo"
  "github.com/gorilla/mux"
)



 
func InitRoutes(r *mux.Router, client *mongo.Client, ctx context.Context) {
  r.HandleFunc("/", controllers.IndexController).Methods("GET")
  r.HandleFunc("/game", controllers.GameController)
}

