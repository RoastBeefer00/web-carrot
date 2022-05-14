package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//go:embed public
var embeddedFiles embed.FS

// func getNames(w http.ResponseWriter, r *http.Request){

// }

type Recipes struct {
	Recipes []Recipe `json:"recipes"`
}

type Recipe struct {
	Title       string   `json:"title"`
	Time        string   `json:"time"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
}

func corsHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			log.Print("preflight detected: ", r.Header)
			w.Header().Add("Connection", "keep-alive")
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
			w.Header().Add("Access-Control-Allow-Headers", "content-type")
			w.Header().Add("Access-Control-Max-Age", "86400")
		} else {
			setupResponse(&w, r)
			h.ServeHTTP(w, r)
		}
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Content-Type", "application/json")
}

func getJson() Recipes {
	jsonFile, err := os.Open("data.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var recipes Recipes

	json.Unmarshal(byteValue, &recipes)

	return recipes
}

func getRandomRecipe(w http.ResponseWriter, r *http.Request) {
	recipes := getJson()

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(recipes.Recipes))
	fmt.Println(i)
	json.NewEncoder(w).Encode(recipes.Recipes[i])
}

func getAllRecipes(w http.ResponseWriter, r *http.Request) {
	recipes := getJson()

	json.NewEncoder(w).Encode(recipes)
}

func main() {
	http.HandleFunc("/api/getrandom", corsHandler(getRandomRecipe))
	http.HandleFunc("/api/getall", corsHandler(getAllRecipes))
	log.Fatal(http.ListenAndServe(":8050", nil))
}
