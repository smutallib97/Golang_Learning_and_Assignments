package main

import (
	"encoding/json"
	"fmt"
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
}
