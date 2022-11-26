package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"web_kana_v1/controllers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		/*
		   If we failed to get the absolute path respond with a 400 bad request
		   and stop
		*/
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepend the path with the path to the static directory
	// For example we can get something like:
	// my_project/build/ + static/js/main.js
	// my_project/build/ + src/index.js
	path = filepath.Join(h.staticPath, path)
	fmt.Println("path after Join: ", path)

	// Check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// File does not exist, serve index.html
		fmt.Println("ServeHTTP: File doesn't exist, serving index.html...")
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		/* If we got an error (that wasn't that the file doesn't exist) stating the
		   file, return a 500 internal server error and stop
		*/
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func InitRoutes(r *mux.Router, client *mongo.Client, ctx context.Context) {
	r.HandleFunc("/game", controllers.NewGameController(ctx, client).Playground)

	spa := spaHandler{staticPath: "../frontend/build", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)
}
