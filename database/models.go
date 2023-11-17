package database

import (
	"encoding/json"
	"os"
)

type Agenda struct {
	Name        string `json:"name"`
	Description string `json:"text"`
}

type Ability struct {
	Name        string `json:"name"`
	Description string `json:"text"`
}

type Leader struct {
	Id            int
	Civilization  string  `json:"name"`
	LeaderAbility Ability `json:"ability"`
	LeaderAgenda  Agenda  `json:"agenda"`
}

type Civlization struct {
	Id      int
	Leaders []string `json:"leaders"`
}

func ReadLeaders() []Leader {
	data, err := os.ReadFile("leaders.json")
	if err != nil {
		panic(err)
	}

	var leaders []Leader = make([]Leader, 0)
	json.Unmarshal(data, &leaders)

	return leaders

}
