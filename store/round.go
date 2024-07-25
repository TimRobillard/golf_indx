package store

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"strconv"
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
	Date   time.Time
}

type RoundStore interface {
	CreateRound(ctx context.Context, u *UIUser, c Course, f, b [9]int, d string) (*Round, error)
	GetRoundsForUser(ctx context.Context, u *UIUser) ([20]*Round, error)
	GetCalcRoundsByUserId(ctx context.Context, userId int) ([20]*CalcRound, error)
}

type Transaction string

const (
	transactionKey Transaction = "TRANSACTION"
)

func (pg PostgresStore) CreateRound(ctx context.Context, u *UIUser, c Course, f, b [9]int, d string) (*Round, error) {
	insertQuery := `
	INSERT INTO round (
		course_id,
		user_id,
		front,
		back,	
		date
	) VALUES (
		$1,$2,$3,$4,$5
	) RETURNING id;`

	date, err := time.Parse(time.DateTime, d)
	if err != nil {
		return nil, err
	}

	round := &Round{
		User:   u,
		Course: c,
		Front:  f,
		Back:   b,
		Date:   date,
	}

	tx, err := pg.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if err := tx.QueryRowContext(ctx, insertQuery, c.Id, u.Id, pq.Array(f), pq.Array(b), date).Scan(&round.Id); err != nil {
		return nil, err
	}

	txCtx := context.WithValue(ctx, transactionKey, tx)
	if rounds, err := pg.GetRoundsForUser(txCtx, u); err != nil {
		return nil, err
	} else {
		if err = pg.SaveHandicap(txCtx, u.Id, rounds, date); err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return round, nil
}

func (pg PostgresStore) GetRoundsForUser(ctx context.Context, u *UIUser) ([20]*Round, error) {
	var rounds [20]*Round

	insertQuery := `
	SELECT 
		c.id as course_id,
		c.slope,
		c.rating,
		c.front as course_back,
		c.back as course_back,
		r.front,
		r.back,
		r.date
	FROM round r
	LEFT JOIN course c
	ON c.id = r.course_id
	WHERE r.user_id = $1
	ORDER BY r.date DESC, r.id DESC
	LIMIT 20;`

	_tx := ctx.Value(transactionKey)
	var tx *sql.Tx
	if _tx != nil {
		tx = _tx.(*sql.Tx)
	}
	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.QueryContext(ctx, insertQuery, u.Id)
	} else {
		rows, err = pg.db.QueryContext(ctx, insertQuery, u.Id)
	}
	if err != nil {
		return rounds, err
	}

	i := 0
	for rows.Next() {
		round := &Round{
			Course: Course{},
		}
		var cf, cb, f, b []int32

		if err := rows.Scan(
			&round.Course.Id,
			&round.Course.Slope,
			&round.Course.Rating,
			(*pq.Int32Array)(&cf),
			(*pq.Int32Array)(&cb),
			(*pq.Int32Array)(&f),
			(*pq.Int32Array)(&b),
			&round.Date,
		); err != nil {
			return rounds, err
		}
		round.User = u
		err = convertNine(cf, &round.Course.Front, err)
		err = convertNine(cb, &round.Course.Back, err)
		err = convertNine(f, &round.Front, err)
		err = convertNine(b, &round.Back, err)
		if err != nil {
			return rounds, err
		}
		rounds[i] = round
		i++
	}

	return rounds, nil
}

func (pg PostgresStore) GetCalcRoundsByUserId(ctx context.Context, userId int) ([20]*CalcRound, error) {
	var rounds [20]*CalcRound

	insertQuery := `
	SELECT 
		c.name,
		r.front,
		r.back,
		r.date
	FROM round r
	LEFT JOIN course c
	ON c.id = r.course_id
	WHERE r.user_id = $1
	ORDER BY r.date DESC, r.id DESC
	LIMIT 20;`

	rows, err := pg.db.QueryContext(ctx, insertQuery, userId)
	if err != nil {
		return rounds, err
	}

	i := 0
	for rows.Next() {
		var courseName string
		var f, b []int32
		var date time.Time
		if err := rows.Scan(&courseName, (*pq.Int32Array)(&f), (*pq.Int32Array)(&b), &date); err != nil {
			return rounds, err
		}

		caser := cases.Title(language.Und)

		if score, err := calcScore(f, b); err != nil {
			return rounds, err
		} else {
			rounds[i] = &CalcRound{
				Course:  caser.String(courseName),
				Score:   score,
				TimeAgo: timeago.English.FormatRelativeDuration(time.Now().Sub(date)),
			}
			i += 1
		}
	}

	return rounds, nil
}

func calcScore(f, b []int32) (string, error) {
	tot := 0
	isNine := false

	if len(f) != 9 {
		return "", fmt.Errorf("invalid input f: %d, b: %d", len(f), len(b))
	}

	if len(b) != 9 {
		if len(b) != 0 {
			return "", fmt.Errorf("invalid input f: %d, b: %d", len(f), len(b))
		}
		isNine = true
	}

	for _, s := range f {
		tot += int(s)
	}
	if isNine {
		return strconv.Itoa(tot*2) + "*", nil
	} else {
		for _, s := range b {
			tot += int(s)
		}
		return strconv.Itoa(tot), nil
	}
}

// Calculates the Score Differential for an 18 hole score
//
// - slope Slope Rating
//
// - score Adjusted gross score
//
// - rating Course Rating
//
// - pcc PCC adjustment
func calc18HoleDifferential(slope, score, rating, pcc float64) float64 {
	return (113 / slope) * (score - rating - pcc)
}

// Calculates the Score Differential for a 9 hole score
//
// - slope Slope Rating
//
// - score Adjusted gross score
//
// - rating Course Rating
//
// - pcc PCC adjustment
func calc9HoleDifferential(slope, score, rating, pcc float64) float64 {
	return (113 / slope) * (score - rating - (pcc / 2)) * 2 // should be + expected score, cannot find this calculation
}

func (r Round) CalcScore() float64 {
	score := 0
	for _, h := range r.Front {
		score += h
	}
	for _, h := range r.Back {
		score += h
	}
	slog.Info(fmt.Sprintf("Score: %d", score))
	return float64(score)
}

func (r Round) CalcDifferential() (float64, error) {
	if len(r.Course.Back) != 9 && len(r.Course.Front) == 9 {
		return calc9HoleDifferential(r.Course.Slope, r.CalcScore(), r.Course.Rating, 0), nil
	}
	if len(r.Course.Front) == 9 && len(r.Course.Back) == 9 {
		return calc18HoleDifferential(r.Course.Slope, r.CalcScore(), r.Course.Rating, 0), nil
	}

	return -1, fmt.Errorf("invalid number of holes: %d", len(r.Course.Front)+len(r.Course.Back))
}

func convertNine(input []int32, nine *[9]int, err error) error {
	if err != nil {
		return err
	}
	if len(input) != 9 && len(input) != 0 {
		return fmt.Errorf("invalid nine length: %d", len(input))
	}

	if len(input) == 0 {
		return nil
	}

	for i := 0; i < 9; i++ {
		nine[i] = int(input[i])
	}

	return nil
}
