package main

import (
	"fmt"
	"github.com/vec-search/lib/redis"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"github.com/vec-search/controller"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	redis.InitRedis()
	router := chi.NewRouter()
	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]interface{}{"hello": "world"})
	})

	controller.EndOfDayRoutes(router)
	controller.AuthenticationRoutes(router)
	controller.TickerRoutes(router)

	fmt.Println("Hello World")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, router)
}
