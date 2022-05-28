package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func Filter(w http.ResponseWriter, r *http.Request) {
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
