package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/2twin/cocktails/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateCocktail(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name         string   `json:"name"`
		Instructions []string `json:"instructions"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	cocktail, err := cfg.DB.CreateCocktail(r.Context(), database.CreateCocktailParams{
		ID:           uuid.New(),
		Name:         params.Name,
		Instructions: params.Instructions,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create equipment: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, cfg.databaseCocktailToCocktail(cocktail))
}

func (cfg *apiConfig) handleGetCocktails(w http.ResponseWriter, r *http.Request) {
	cocktails, err := cfg.DB.GetCocktails(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get cocktails: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, cfg.databaseCocktailsToCocktails(cocktails))
}

func (cfg *apiConfig) handleGetCocktailById(w http.ResponseWriter, r *http.Request) {
	cocktailID := chi.URLParam(r, "cocktailID")
	cocktailUuid, err := uuid.Parse(cocktailID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't parse cockatil's id: %v", err))
	}

	cocktail, err := cfg.DB.GetCocktailById(r.Context(), cocktailUuid)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't find cocktail: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, cfg.databaseCocktailToCocktail(cocktail))
}

func (cfg *apiConfig) handleAddIngredientToCocktail(w http.ResponseWriter, r *http.Request) {
	cocktailID := chi.URLParam(r, "cocktailID")
	cocktailUuid, err := uuid.Parse(cocktailID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't parse cockatil's id: %v", err))
	}

	type parameters struct {
		IngredientID uuid.UUID `json:"ingredient_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	_, err = cfg.DB.AddIngredientToCocktail(r.Context(), database.AddIngredientToCocktailParams{
		CocktailID:   cocktailUuid,
		IngredientID: params.IngredientID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't add ingredient to cockail: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}

func (cfg *apiConfig) handleAddEquipmentToCocktail(w http.ResponseWriter, r *http.Request) {
	cocktailID := chi.URLParam(r, "cocktailID")
	cocktailUuid, err := uuid.Parse(cocktailID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't parse cockatil's id: %v", err))
	}

	type parameters struct {
		EquipmentID uuid.UUID `json:"equipment_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	_, err = cfg.DB.AddEquipmentToCocktail(r.Context(), database.AddEquipmentToCocktailParams{
		CocktailID:  cocktailUuid,
		EquipmentID: params.EquipmentID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't add equipment to cockail: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
