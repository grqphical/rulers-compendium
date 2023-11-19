package database

import (
	"encoding/json"
	"log"
	"os"
)

type Database struct {
	Leaders       []Leader
	Civilizations []Civilization
	Districts     []District
	Improvements  []Improvement
}

type Agenda struct {
	Name        string `json:"name"`
	Description string `json:"text"`
}

type Ability struct {
	Name        string `json:"name"`
	Description string `json:"text"`
}

type Leader struct {
	Name          string  `json:"name"`
	Civilization  string  `json:"civ"`
	LeaderAbility Ability `json:"ability"`
	LeaderAgenda  Agenda  `json:"agenda"`
	UniqueUnit    string  `json:"unit"`
}

type Civilization struct {
	Name           string   `json:"name"`
	Leaders        []string `json:"leaders"`
	Ability        Ability  `json:"ability"`
	Unit           string   `json:"unit"`
	Infrastructure string   `json:"infrastructure"`
}

type District struct {
	Name          string   `json:"name"`
	Buildings     []string `json:"buildings"`
	Description   string   `json:"description"`
	PlunderYields string   `json:"plunder_yields"`
	ExclusiveTo   string   `json:"exclusive_to"`
}

type Improvement struct {
	Name        string
	Technology  string
	Placement   []string
	Resources   []string
	Description string
	Plunder     string
}

func ReadDatabase() Database {
	db := Database{}

	data, err := os.ReadFile("data/leaders.json")
	if err != nil {
		log.Fatal("Failed to read leaders file")
	}

	db.Leaders = make([]Leader, 0)
	json.Unmarshal(data, &db.Leaders)

	data, err = os.ReadFile("data/civs.json")
	if err != nil {
		log.Fatal("Failed to read civilizations file")
	}

	db.Civilizations = make([]Civilization, 0)
	json.Unmarshal(data, &db.Civilizations)

	data, err = os.ReadFile("data/districts.json")
	if err != nil {
		log.Fatal("Failed to read districts file")
	}

	db.Districts = make([]District, 0)
	json.Unmarshal(data, &db.Districts)

	data, err = os.ReadFile("data/improvements.json")
	if err != nil {
		log.Fatal("Failed to read improvements file")
	}

	db.Improvements = make([]Improvement, 0)
	json.Unmarshal(data, &db.Improvements)

	return db
}
