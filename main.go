package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
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

func getFilteredRecipes(w http.ResponseWriter, r *http.Request) {
	// keys, ok := r.URL.Query()["key"]
	key := strings.TrimPrefix(r.URL.Path, "/api/filter/")
	// if !ok || len(keys[0]) < 1 {
	// 	log.Println("Url Param 'key' is missing")
	// 	return
	// }

	// Query()["key"] will return an array of items,
	// we only want the single item.
	//key := keys[0]

	log.Println("Url Param 'key' is: " + key)

	recipes := getJson()
	var ret Recipes
	for _, recipe := range recipes.Recipes {
		if strings.Contains(strings.ToLower(recipe.Title), strings.ToLower(key)) {
			ret.Recipes = append(ret.Recipes, recipe)
		}
	}
	json.NewEncoder(w).Encode(ret)
}

//go:embed public
var embeddedFiles embed.FS

func main() {
	fsys, err := fs.Sub(embeddedFiles, "public")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.FS(fsys)))

	http.HandleFunc("/api/getrandom", corsHandler(getRandomRecipe))
	http.HandleFunc("/api/getall", corsHandler(getAllRecipes))
	http.HandleFunc("/api/filter/", corsHandler(getFilteredRecipes))
	log.Fatal(http.ListenAndServe(":8050", nil))
}
