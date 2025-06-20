package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Product represents the structure of our resource
type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// In-memory slice to store products
var products []Product

// Get all products
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Get single product by ID
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get route parameters
	id := params["id"]

	for _, item := range products {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Product{})
}

// Create a new product
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	// Generate simple numeric ID
	product.ID = strconv.Itoa(len(products) + 1)
	products = append(products, product)

	json.NewEncoder(w).Encode(product)
}

// Update a product
func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var updatedProduct Product
	_ = json.NewDecoder(r.Body).Decode(&updatedProduct)

	for index, item := range products {
		if item.ID == id {
			updatedProduct.ID = id
			products[index] = updatedProduct
			json.NewEncoder(w).Encode(updatedProduct)
			return
		}
	}

	json.NewEncoder(w).Encode(products)
}

// Delete a product
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for index, item := range products {
		if item.ID == id {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(products)
}

func main() {
	r := mux.NewRouter()

	// Seed some initial data
	products = append(products,
		Product{ID: "1", Name: "Chair"},
		Product{ID: "2", Name: "Table"},
		Product{ID: "3", Name: "Mouse"},
		Product{ID: "4", Name: "Monitor"},
		Product{ID: "5", Name: "Cable"},
		Product{ID: "6", Name: "Smartphone"},
		Product{ID: "7", Name: "Keyboard"},
		Product{ID: "8", Name: "Headset"},
		Product{ID: "9", Name: "Bottle"},
		Product{ID: "10", Name: "CPU"},
	)

	// Routes
	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/products", createProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/api/products/{id}", deleteProduct).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
