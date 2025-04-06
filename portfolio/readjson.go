package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type CardsData struct {
	Cards []ProjectCard `json:"cards"`
}

func GetCards() CardsData {
	// Read the JSON file (adjust path if needed)
	data, err := os.ReadFile("cardsData.json")
	if err != nil {
		log.Fatalf("failed to read JSON file: %v", err)
	}

	var cardsData CardsData
	if err := json.Unmarshal(data, &cardsData); err != nil {
		log.Fatalf("failed to parse JSON: %v", err)
	}

	fmt.Printf("Cards: %+v\n", cardsData)
	return cardsData

}
