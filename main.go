package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/2twin/cocktails/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't find .env file")
	}

	const filepathRoot = "."
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	apiCfg := apiConfig{
		DB: dbQueries,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		},
	))

	apiRouter := chi.NewRouter()

	apiRouter.Get("/ingredients", apiCfg.handleGetIngredients)
	apiRouter.Post("/ingredients/create", apiCfg.handlerCreateIngredient)

	apiRouter.Get("/equipments", apiCfg.handleGetEquipments)
	apiRouter.Post("/equipments/create", apiCfg.handlerCreateEquipment)

	apiRouter.Get("/cocktails", apiCfg.handleGetCocktails)
	apiRouter.Post("/cocktails/create", apiCfg.handlerCreateCocktail)
	apiRouter.Get("/cocktails/{cocktailID}", apiCfg.handleGetCocktailById)
	apiRouter.Post("/cocktails/{cocktailID}/add_ingredient", apiCfg.handleAddIngredientToCocktail)
	apiRouter.Post("/cocktails/{cocktailID}/add_equipment", apiCfg.handleAddEquipmentToCocktail)

	router.Mount("/v1", apiRouter)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
