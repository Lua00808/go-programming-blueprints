package main

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/go-programming-blueprints/meander"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// meander.APIKey = "Todo"
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	return json.NewEncoder(w).Encode(data)
}
