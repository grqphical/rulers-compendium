package database

type Agenda struct {
	Name        string
	Description string
}

type Leader struct {
	Id            int
	Civilization  int
	LeaderAbility string
	LeaderAgenda  Agenda
}

type Civlization struct {
	Id      int
	Leaders []int
}
