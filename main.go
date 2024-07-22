package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"github.com/vec-search/controller"
)

func main() {
	router := chi.NewRouter()
	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]interface{}{"hello": "world"})
	})

	controller.EndOfDayRoutes(router)

	fmt.Println("Hello World")

	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
