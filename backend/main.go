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
	"web_kana_v1/routes"
	"github.com/gorilla/mux"
)

type spaHandler struct {
  staticPath string 
  indexPath string
}


func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // get the absolute path to prevent directory traversal
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


 
  // We need something like "/*" so we can serve index.html (which React dynamically changes?) 
  // on all endpoints

  /* Actually, from Mux documentation we can see that using PathPrefix("/") actually makes it so 
     that Handler(spa) gets called on /*    :
    
    "Note that the path provided to PathPrefix() represents a "wildcard":
     calling PathPrefix("/static/").Handler(...) means that the handler will 
     be passed any request that matches "/static/*". 
     This makes it easy to serve static files with mux:"
  */

   // FIXME: Problem:                                                                   
   /* 
      Need to figure out how to make Playground controller respond to /game endpoint,
      since r.HandleFunc("/game", controllers.NewGameController(ctx, client).Playground)
      in routes.go isn't getting called when React locally changes endpoint to /game
    */

  
	// For serving the static files (OLD WAY)
/*   r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../frontend/build/static")))) */

  r.PathPrefix("/").Handler(spa) 

	// Initalize the routes
	routes.InitRoutes(r, client, ctx)

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
