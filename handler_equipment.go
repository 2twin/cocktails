package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/2twin/cocktails/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateEquipment(w http.ResponseWriter, r *http.Request) {
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

	equipment, err := cfg.DB.CreateEquipment(r.Context(), database.CreateEquipmentParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create equipment: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseEquipmentToEquipment(equipment))
}

func (cfg *apiConfig) handleGetEquipments(w http.ResponseWriter, r *http.Request) {
	Equipments, err := cfg.DB.GetEquipments(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get equipments: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseEquipmentsToEquipments(Equipments))
}
