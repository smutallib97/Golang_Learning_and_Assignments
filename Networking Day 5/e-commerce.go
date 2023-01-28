package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Product_ID   int    `json:"product_id"`
	Product_Name string `json:"product_name"`
	Price        int    `json:"price"`
}

var products []Product

func main() {
	// Read product information from a JSON file
	data, _ := ioutil.ReadFile("products.json")
	json.Unmarshal(data, &products)
	http.HandleFunc("/products", handleProducts)
	fmt.Println("Server starting on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetProducts(w, r)
	case http.MethodPost:
		handlePostProduct(w, r)
	case http.MethodPut:
		handlePutProduct(w, r)
	case http.MethodDelete:
		handleDeleteProduct(w, r)
	}
}

func handleGetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func handlePostProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&product)
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

func handlePutProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	for i, p := range products {
		if p.Product_ID == product.Product_ID {
			products[i] = product
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	for i, p := range products {
		if p.Product_ID == product.Product_ID {
			products = append(products[:i], products[i+1:]...)
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
