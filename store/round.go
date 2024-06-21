package store

type CalcRound struct {
	Score   string
	Course  string
	TimeAgo string
}

type Round struct {
	Id     int
	User   *User
	Course *Course
	Front  [9]int
	Back   [9]int
}
