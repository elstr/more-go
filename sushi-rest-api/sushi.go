package main

import (
	"strconv"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Roll model
type Roll struct {
	ID string `json:"id"`
	ImageNumber string `json:"imageNumber"`
	Name string `json:"name"`
	Ingredients string `json:"ingredients"`
}

// Where we will save our rolls
var rolls []Roll

func getRolls(w http.ResponseWriter, r *http.Request){
	// 0. Set headers for the response
	w.Header().Set("Content-Type", "application/json")
	// 1. Render the roll slice as json
	json.NewEncoder(w).Encode(rolls)
}

func getRoll(w http.ResponseWriter, r *http.Request){
		// 0. Set headers for the response
		w.Header().Set("Content-Type", "application/json")
		// 1. Get the ID of the roll
		params := mux.Vars(r)
		// 2. Find the rol that has the ID given
		for _, roll := range rolls {
			if roll.ID == params["id"] {
				// 3. Render the roll
				json.NewEncoder(w).Encode(roll)
				return
			}
		}
	
}

func createRoll(w http.ResponseWriter, r *http.Request){
		// 0. Set headers for the response
		w.Header().Set("Content-Type", "application/json")
		// 1. Define the new role
		var newRoll Roll
		// 2. Get new roll data from request body and assign it to newRoll
		err := json.NewDecoder(r.Body).Decode(&newRoll)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// 3. Set new ID 
		newRoll.ID = strconv.Itoa(len(rolls) +1) // ver como recuperar el ultimo id
		// 4. Update rolls
		rolls = append(rolls, newRoll)
		// 5. Return created roll
		json.NewEncoder(w).Encode(newRoll)
}

func updateRoll(w http.ResponseWriter, r *http.Request){
	// 0. Set headers for the response
	w.Header().Set("Content-Type", "application/json")
	// 1. Define the new role
	var newRoll Roll
	// 2. Get params
	params := mux.Vars(r)
	// 3. Search the ID and remove it from the array 
	for i, roll := range rolls {
		if roll.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			newRoll.ID = params["id"]
			json.NewDecoder(r.Body).Decode(&newRoll)
			rolls = append(rolls, newRoll)
			json.NewEncoder(w).Encode(rolls)
			return
		}
	}
}

func deleteRoll(w http.ResponseWriter, r *http.Request){
	// 0. Set headers for the response
	w.Header().Set("Content-Type", "application/json")
	// 2. Get params
	params := mux.Vars(r)
	// 3. Search the ID and remove it from the array 
	for i, roll := range rolls {
		if roll.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			break		
		}
	}
	// 2. Return the rolls
	json.NewEncoder(w).Encode(rolls)
}

func main() {  
	rolls = append(rolls, 
		Roll{ID: "1", ImageNumber: "1", Name: "Spicy Tuna Roll", Ingredients: "Tuna, Chili sauce, Nori, Rice"},
		Roll{ID: "2", ImageNumber: "2", Name: "NY", Ingredients: "Salmon, Avocado, Nori, Rice"},
	)

	// 0. Define the router
	router := mux.NewRouter()

	// 1. Define the endpoints
	router.HandleFunc("/sushi",getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("PUT")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	// 2. Initialize server and throw an error if it fails
	log.Fatal(http.ListenAndServe(":8000", router))

}