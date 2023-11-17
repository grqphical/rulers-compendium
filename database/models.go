package database

type Agenda struct {
	Name        string
	Description string
}

type Leader struct {
	Id           int
	Civilization int
	LeaderBonus  string
	LeaderAgenda Agenda
	Religion     string
}

type Civlization struct {
	Id      int
	Leaders []int
}
