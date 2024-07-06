package store

import "github.com/lib/pq"

type CalcRound struct {
	Score   string
	Course  string
	TimeAgo string
}

type Round struct {
	Id     int
	User   *UIUser
	Course Course
	Front  [9]int
	Back   [9]int
}

type RoundStore interface {
	CreateRound(u *UIUser, c Course, f, b [9]int) (*Round, error)
	GetRoundsByUserId(userId int) ([]*Round, error)
}

func (pg PostgresStore) CreateRound(u *UIUser, c Course, f, b [9]int) (*Round, error) {
	query := `
	INSERT INT round (
		course_id,
		user_id,
		front,
		back	
	) VALUES (
		$1,$2,$3,$4
	);`

	round := &Round{
		User:   u,
		Course: c,
		Front:  f,
		Back:   b,
	}

	if err := pg.db.QueryRow(query, c.Id, u.Id, pq.Array(f), pq.Array(b)).Scan(round.Id); err != nil {
		return nil, err
	}

	return round, nil
}

func (pg PostgresStore) GetRoundsByUserId(userId int) ([]*Round, error) {
	return nil, nil
}
