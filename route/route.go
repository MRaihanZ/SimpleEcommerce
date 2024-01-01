package route

import (
	"SimpleEcommerce/database/product"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors" // Import the CORS package
)

func Routes() {
	// Create a new router
	router := httprouter.New()

	// Define your routes
	subdir := make(map[string]string)
	subdir["getProducts"] = "/api/products"
	subdir["getCarts"] = "/api/carts"
	subdir["getCart"] = "/api/carts/:id"
	subdir["postCart"] = "/api/cart"
	subdir["patchCart"] = "/api/cart/:id/:quantity"
	subdir["deleteCart"] = "/api/cart/:id"

	router.GET(subdir["getProducts"], product.GetProducts)
	router.GET(subdir["getCarts"], product.GetCarts)
	router.GET(subdir["getCart"], product.GetCart)
	router.POST(subdir["postCart"], product.PostCart)
	router.PATCH(subdir["patchCart"], product.PatchCart)
	router.DELETE(subdir["deleteCart"], product.DeleteCart)

	// Create a new CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowedMethods:   []string{"OPTIONS", "GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "X-CSRF-Token", "client-platform", "application/json"},
		AllowCredentials: true,
	})

	// Use the CORS middleware with the router
	handler := c.Handler(router)

	// Create a new HTTP server
	server := http.Server{
		Handler: handler,
		Addr:    "localhost:8080",
	}

	// Print server information
	fmt.Println("Server URL:", server.Addr)
	for _, value := range subdir {
		fmt.Printf("Sub Directory: %s\n", value)
	}

	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
