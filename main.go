package main

import (
	"fmt"
	"log"
	"net/http"
	"web_kana_v1/connMonitor"
	"web_kana_v1/dbLogic"
	"web_kana_v1/routes"

	"github.com/gorilla/mux"
)

func main() {
	client, ctx, cancel := dbLogic.InitializeDatabaseConnection()
	defer dbLogic.Close(client, ctx, cancel)
	dbLogic.Ping(client, ctx)

	r := mux.NewRouter()

	// For serving the static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Initalize the routes
	routes.InitRoutes(r, client, ctx)

	// Configure the web server
	var cw connMonitor.ConnectionWatcher
	s := &http.Server{
		Addr:      ":8000",
		Handler:   r,
		ConnState: cw.OnStateChange,
	}
	fmt.Println("Starting server at port 8000...")
	log.Fatal(s.ListenAndServe())
}
