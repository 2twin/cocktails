package main

import (
	"context"
	"log"
	"time"

	"github.com/2twin/cocktails/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Favorites struct {
	ID        uuid.UUID  `json:"id"`
	UserID    uuid.UUID  `json:"user_id"`
	Cocktails []Cocktail `json:"cocktails"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Cocktail struct {
	ID           uuid.UUID    `json:"id"`
	Name         string       `json:"name"`
	Instructions []string     `json:"instructions"`
	Ingredients  []Ingredient `json:"ingredients"`
	Equipments   []Equipment  `json:"equipments"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type Ingredient struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Equipment struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (cfg *apiConfig) AddIngredientsAndEquipmentsToCocktail(cocktail *Cocktail) *Cocktail {
	dbIngredients, err := cfg.DB.GetCocktailIngredients(context.Background(), cocktail.ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	ingredients := []Ingredient{}
	for _, item := range dbIngredients {
		ingredient, err := cfg.DB.GetIngredientById(context.Background(), item.IngredientID)
		if err != nil {
			log.Fatal(err.Error())
		}
		ingredients = append(ingredients, databaseIngredientToIngredient(ingredient))
	}

	dbEquipments, err := cfg.DB.GetCocktailEquipments(context.Background(), cocktail.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	
	equipments := []Equipment{}
	for _, item := range dbEquipments {
		equipment, err := cfg.DB.GetEquipmentById(context.Background(), item.EquipmentID)
		if err != nil {
			log.Fatal(err.Error())
		}
		equipments = append(equipments, databaseEquipmentToEquipment(equipment))
	}

	cocktail.Ingredients = ingredients
	cocktail.Equipments = equipments
	return cocktail
}

func (cfg *apiConfig) databaseCocktailToCocktail(dbCocktail database.Cocktail) Cocktail {
	cocktail := Cocktail{
		ID:           dbCocktail.ID,
		Instructions: dbCocktail.Instructions,
		CreatedAt:    dbCocktail.CreatedAt,
		UpdatedAt:    dbCocktail.UpdatedAt,
	}

	cocktail = *cfg.AddIngredientsAndEquipmentsToCocktail(&cocktail)
	return cocktail
}

func (cfg *apiConfig) databaseCocktailsToCocktails(dbCocktails []database.Cocktail) []Cocktail {
	cocktails := []Cocktail{}
	for _, cocktail := range dbCocktails {
		cocktails = append(cocktails, cfg.databaseCocktailToCocktail(cocktail))
	}
	return cocktails
}

func databaseIngredientToIngredient(dbIngredient database.Ingredient) Ingredient {
	return Ingredient{
		ID:   dbIngredient.ID,
		Name: dbIngredient.Name,
		CreatedAt: dbIngredient.CreatedAt,
		UpdatedAt: dbIngredient.UpdatedAt,
	}
}

func databaseIngredientsToIngredients(dbIngredients []database.Ingredient) []Ingredient {
	ingredients := []Ingredient{}
	for _, ingredient := range dbIngredients {
		ingredients = append(ingredients, databaseIngredientToIngredient(ingredient))
	}
	return ingredients
}

func databaseEquipmentToEquipment(dbEquipment database.Equipment) Equipment {
	return Equipment{
		ID:   dbEquipment.ID,
		Name: dbEquipment.Name,
		CreatedAt: dbEquipment.CreatedAt,
		UpdatedAt: dbEquipment.UpdatedAt,
	}
}

func databaseEquipmentsToEquipments(dbEquipments []database.Equipment) []Equipment {
	equipments := []Equipment{}
	for _, equipment := range dbEquipments {
		equipments = append(equipments, databaseEquipmentToEquipment(equipment))
	}
	return equipments
}
