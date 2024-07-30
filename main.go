package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/vec-search/lib"
	"github.com/vec-search/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	router := chi.NewRouter()

	lib.IntiPolygon()
	lib.SetUpMiddleware(router)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.PlainText(w, r, "Hello World")
	})

	routes.GatewayRoutes(router)

	fmt.Println("Hello World")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, router)
}
