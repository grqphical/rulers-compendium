package database

import (
	"encoding/json"
	"os"
)

type Database struct {
	Leaders       []Leader
	Civilizations []Civlization
	Districts     []District
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
}

type Civlization struct {
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
}

func ReadDatabase() Database {
	db := Database{}

	data, err := os.ReadFile("data/leaders.json")
	if err != nil {
		panic(err)
	}

	db.Leaders = make([]Leader, 0)
	json.Unmarshal(data, &db.Leaders)

	data, err = os.ReadFile("data/civs.json")
	if err != nil {
		panic(err)
	}

	db.Civilizations = make([]Civlization, 0)
	json.Unmarshal(data, &db.Civilizations)

	data, err = os.ReadFile("data/districts.json")
	if err != nil {
		panic(err)
	}

	db.Districts = make([]District, 0)
	json.Unmarshal(data, &db.Districts)

	return db
}
