package routes 

import (
  "web_kana_v1/controllers"
  "github.com/gorilla/mux"
)

 
func InitRoutes(r *mux.Router) {
  r.HandleFunc("/", controllers.IndexController).Methods("GET")
  r.HandleFunc("/game", controllers.GameController)
}

