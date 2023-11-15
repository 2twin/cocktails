package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/2twin/cocktails/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateIngredient(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	ingredient, err := cfg.DB.CreateIngredient(r.Context(), database.CreateIngredientParams{
		ID: uuid.New(),
		Name: params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create ingredient: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseIngredientToIngredient(ingredient))
}

func (cfg *apiConfig) handleGetIngredients(w http.ResponseWriter, r *http.Request) {
	ingredients, err := cfg.DB.GetIngredients(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get ingredients: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseIngredientsToIngredients(ingredients))
}