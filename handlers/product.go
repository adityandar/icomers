package handlers

import (
	"encoding/json"
	"fmt"
	"icomers/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var products = []models.Product{
	{ID: 1, Name: "Laptop", Description: "High-performance laptop", Price: 27000500.99, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Name: "Smartphone", Description: "Latest release smartphone", Price: 200000.83, CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

// func to get all products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// func to get product by ID
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range products {
		if fmt.Sprintf("%d", item.ID) == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// func to create new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = len(products) + 1
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	products = append(products, product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedProducts models.Product
	_ = json.NewDecoder(r.Body).Decode(&updatedProducts)

	for i, item := range products {
		if fmt.Sprintf("%d", item.ID) == params["id"] {
			products[i].Name = updatedProducts.Name
			products[i].Description = updatedProducts.Description
			products[i].Price = updatedProducts.Price
			products[i].UpdatedAt = time.Now()

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return

		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range products {
		if fmt.Sprintf("%d", item.ID) == params["id"] {
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}
