package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Expanded Product struct with more detailed information
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	Brand       string    `json:"brand"`
	InStock     bool      `json:"in_stock"`
	Rating      float64   `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	ImageURL    string    `json:"image_url"`
}

// Generate mock product data
func Products() []Product {
	return []Product{
		{
			ID:          1,
			Name:        "Smartphone Pro X",
			Description: "High-performance smartphone with advanced camera technology",
			Price:       799.99,
			Category:    "Electronics",
			Brand:       "TechGiant",
			InStock:     true,
			Rating:      4.7,
			CreatedAt:   time.Now(),
			ImageURL:    "/api/placeholder/400/300",
		},
		{
			ID:          2,
			Name:        "Wireless Noise-Canceling Headphones",
			Description: "Premium over-ear headphones with active noise cancellation",
			Price:       349.99,
			Category:    "Audio",
			Brand:       "SoundWave",
			InStock:     true,
			Rating:      4.5,
			CreatedAt:   time.Now(),
			ImageURL:    "/api/placeholder/400/300",
		},
		{
			ID:          3,
			Name:        "Smart Home Hub",
			Description: "Central control system for all your smart home devices",
			Price:       199.99,
			Category:    "Smart Home",
			Brand:       "HomeConnect",
			InStock:     false,
			Rating:      4.2,
			CreatedAt:   time.Now(),
			ImageURL:    "/api/placeholder/400/300",
		},
		{
			ID:          4,
			Name:        "Ergonomic Office Chair",
			Description: "Comfortable chair with lumbar support for long working hours",
			Price:       299.99,
			Category:    "Furniture",
			Brand:       "ComfortWork",
			InStock:     true,
			Rating:      4.8,
			CreatedAt:   time.Now(),
			ImageURL:    "/api/placeholder/400/300",
		},
		{
			ID:          5,
			Name:        "4K Ultra HD Monitor",
			Description: "27-inch display with exceptional color accuracy",
			Price:       499.99,
			Category:    "Electronics",
			Brand:       "ScreenPro",
			InStock:     true,
			Rating:      4.6,
			CreatedAt:   time.Now(),
			ImageURL:    "/api/placeholder/400/300",
		},
		{
			ID:          6,
			Name:        "Portable Bluetooth Speaker",
			Description: "Waterproof speaker with 360-degree sound",
			Price:       129.99,
			Category:    "Audio",
			Brand:       "SoundBlast",
			InStock:     true,
			Rating:      4.4,
			CreatedAt:   time.Now(),
			ImageURL:    "/api/placeholder/400/300",
		},{
			ID:          7,
			Name:        "Smart Fitness Tracker",
			Description: "Advanced health monitoring with GPS and heart rate tracking",
			Price:       149.99,
			Category:    "Wearables",
			Brand:       "FitTech",
			InStock:     true,
			Rating:      4.6,
			CreatedAt:   time.Now(),
			ImageURL:    "/api/placeholder/400/300",
		},
		{
			ID:          8,
			Name:        "Noise-Canceling Earbuds",
			Description: "Wireless earbuds with active noise cancellation and long battery life",
			Price:       249.99,
			Category:    "Audio",
			Brand:       "SoundWave",
			InStock:     true,
			Rating:      4.5,
			CreatedAt:   time.Now(),
			ImageURL:    "/api/placeholder/400/300",
		},
		{
			ID:          9,
			Name:        "Smart Home Security Camera",
			Description: "4K resolution camera with AI-powered motion detection",
			Price:       199.99,
			Category:    "Smart Home",
			Brand:       "SecureView",
			InStock:     false,
			Rating:      4.3,
			CreatedAt:   time.Now(),
			ImageURL:    "/api/placeholder/400/300",
		},
	}
}

// Handlers for each rendering strategy
func SSRProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := Products()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func SSGProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := Products()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func ISRProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := Products()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	router := mux.NewRouter()

	// CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Route handlers
	router.HandleFunc("/api/ssr-products", SSRProductsHandler).Methods("GET")
	router.HandleFunc("/api/ssg-products", SSGProductsHandler).Methods("GET")
	router.HandleFunc("/api/isr-products", ISRProductsHandler).Methods("GET")

	// Start server
	handler := c.Handler(router)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}