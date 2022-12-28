package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	AppleRecipe := `{
	"Thinly Sliced Peeled Apples":"6 Cups",
	"sugar":"3/4 cup",
	"flour":"2 tablespoons",
	"cinamon":"1/4 teaspoon",
	"nutmeg":"1/8 tablesppon",
	"lemon juice":"1 tablespoon"}`

	AppleRecipeMap := map[string]string{}

	json.Unmarshal([]byte(AppleRecipe), &AppleRecipeMap)
	fmt.Println(AppleRecipeMap)

	for Ingredients := range AppleRecipeMap {
		fmt.Println(Ingredients, "=", AppleRecipeMap[Ingredients])
	}

	file, err := os.Create("Applepie_recipie.json")
	defer file.Close()

	if err != nil {
		panic(err)
	}
	for key, value := range AppleRecipeMap {
		_, err = file.WriteString(fmt.Sprintf("%s : %s\n", key, value))
		if err != nil {
			panic(err)
		}
	}
}
