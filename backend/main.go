package main

import (
	"fmt"
	"log"
	"net/http"
  "os"
  "path/filepath"
  "encoding/json"
  "time"

	"web_kana_v1/connMonitor"
	"web_kana_v1/dbLogic"
	// "web_kana_v1/routes"


	"github.com/gorilla/mux"
)

type spaHandler struct {
  staticPath string 
  indexPath string
}


func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // get the absolute path to prevent directory traversal
  // fmt.Println("h: ", h)
  // fmt.Println("w: ", w)
  // fmt.Println("r: ", r)
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
        // if we failed to get the absolute path respond with a 400 bad request
        // and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

  // prepend the path with the path to the static directory

  // For example we can get something like:
  // my_project/build/ + static/js/main.js 
  // my_project/build/ + src/index.js
  path = filepath.Join(h.staticPath, path)
  fmt.Println("path after Join: ", path)

  // check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
      // if we got an error (that wasn't that the file doesn't exist) stating the
      // file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  // Otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}


func main() {
	client, ctx, cancel := dbLogic.InitializeDatabaseConnection()
	defer dbLogic.Close(client, ctx, cancel)
	dbLogic.Ping(client, ctx)

	r := mux.NewRouter()

  r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

  spa := spaHandler{staticPath: "../frontend/build", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa) 

	// For serving the static files
	  //r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Initalize the routes
	  //routes.InitRoutes(r, client, ctx)

	// Configure the web server
	var cw connMonitor.ConnectionWatcher
	s := &http.Server{
		Addr:      ":8000",
		Handler:   r,
		ConnState: cw.OnStateChange,
    WriteTimeout: 15 * time.Second,
    ReadTimeout: 15 * time.Second,
	}
	fmt.Println("Starting server at port 8000...")
	log.Fatal(s.ListenAndServe())
}
