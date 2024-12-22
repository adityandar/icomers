package handlers

import (
	"encoding/json"
	"icomers/database"
	"icomers/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var products = []models.Product{
	{ID: 1, Name: "Laptop", Description: "High-performance laptop", Price: 27000500.99, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Name: "Smartphone", Description: "Latest release smartphone", Price: 200000.83, CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

// func to get all products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var allProducts []models.Product
	if err := database.DB.Find(&allProducts).Error; err != nil {
		http.Error(w, "Erorr when fetching all products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allProducts)
}

// func to get product by ID
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product models.Product

	if err := database.DB.First(&product, params["id"]).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error when fetching products", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// func to create new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&product).Error; err != nil {
		http.Error(w, "Error when creating product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product models.Product

	if err := database.DB.First(&product, params["id"]).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error when fetching products", http.StatusInternalServerError)
		}
		return
	}

	var updatedProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}

	product.Name = updatedProduct.Name
	product.Description = updatedProduct.Description
	product.Price = updatedProduct.Price
	product.UpdatedAt = updatedProduct.UpdatedAt

	if err := database.DB.Save(&product).Error; err != nil {
		http.Error(w, "Error while updating product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if err := database.DB.Delete(&models.Product{}, params["id"]).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
