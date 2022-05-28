package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Recipes struct {
	Recipes []Recipe `json:"recipes"`
}

type Recipe struct {
	Title       string   `json:"title"`
	Time        string   `json:"time"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
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

func Handler(w http.ResponseWriter, r *http.Request) {
	recipes := getJson()

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(recipes.Recipes))
	fmt.Println(i)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(recipes.Recipes[i])
}
