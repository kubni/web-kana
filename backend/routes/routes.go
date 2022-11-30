package routes

import (
	"context"
	"web_kana_v1/controllers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(r *mux.Router, client *mongo.Client, ctx context.Context) {
  r.HandleFunc("/game/checkAnswer", controllers.CheckAnswer).Methods("POST")
  r.HandleFunc("/game/generateHiraganaCharacter", controllers.GenerateHiraganaCharacter).Methods("GET")
  r.HandleFunc("/game/generateKatakanaCharacter", controllers.GenerateKatakanaCharacter).Methods("GET")
  r.HandleFunc("/game/insertUserIntoDatabase", controllers.NewGameController(ctx, client).HandleUserData).Methods("POST")
  r.HandleFunc("/game/calculatePlayerRank", controllers.NewGameController(ctx, client).CalculatePlayerRank).Methods("POST")
	r.HandleFunc("/game", controllers.NewGameController(ctx, client).Playground)
}
