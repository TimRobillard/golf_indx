package store

import (
	"time"

	"github.com/lib/pq"
	"github.com/xeonx/timeago"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

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
	Date   string
}

type RoundStore interface {
	CreateRound(u *UIUser, c Course, f, b [9]int, d string) (*Round, error)
	GetRoundsByUserId(userId int) ([20]*CalcRound, error)
}

func (pg PostgresStore) CreateRound(u *UIUser, c Course, f, b [9]int, d string) (*Round, error) {
	query := `
	INSERT INTO round (
		course_id,
		user_id,
		front,
		back,	
		date
	) VALUES (
		$1,$2,$3,$4,$5
	);`

	round := &Round{
		User:   u,
		Course: c,
		Front:  f,
		Back:   b,
		Date:   d,
	}

	if err := pg.db.QueryRow(query, c.Id, u.Id, pq.Array(f), pq.Array(b)).Scan(round.Id); err != nil {
		return nil, err
	}

	return round, nil
}

func (pg PostgresStore) GetRoundsByUserId(userId int) ([20]*CalcRound, error) {
	var rounds [20]*CalcRound

	query := `
	SELECT 
		c.name,
		r.front,
		r.back,
		r.date
	FROM round r
	LEFT JOIN course c
	ON c.id = r.course_id
	WHERE r.user_id = $1
	ORDER BY r.date DESC
	LIMIT 20;`

	caser := cases.Title(language.Und)

	rows, err := pg.db.Query(query, userId)
	if err != nil {
		return rounds, err
	}

	i := 0
	for rows.Next() {
		round := &CalcRound{}
		var front, back []int32
		var date time.Time
		if err := rows.Scan(&round.Course, (*pq.Int32Array)(&front), (*pq.Int32Array)(&back), &date); err != nil {
			return rounds, err
		}
		round.Course = caser.String(round.Course)
		round.Score = getScoreString(front, back)
		round.TimeAgo = timeago.English.FormatRelativeDuration(time.Now().Sub(date))
		rounds[i] = round
		i += 1
	}

	return rounds, nil
}

func getScoreString(front, back []int32) string {
	var tot int32 = 0

	for _, s := range front {
		tot += s
	}
	for _, s := range back {
		tot += s
	}

	return String(tot)
}

func String(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
