package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Val struct {
	Value string `json:"value,omitempty"`
}

func GetVals(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(r.Body)
}

func index(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var board [9]string
	var tab string = params["value"]
	if len(tab) != 9 {
		var val Val
		val.Value = "error"
		json.NewEncoder(w).Encode(&val)
		return
	} else {
		for i := 0; i < len(tab); i++ {
			board[i] = string(tab[i])
		}
		index, _ := strconv.Atoi(callAi(board))
		board[index] = "X"
		json.NewEncoder(w).Encode(&board)
	}
}

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/nextMove/{value}", index).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(router)))
}
