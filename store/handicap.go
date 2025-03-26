package store

import (
	"context"
	"database/sql"
	"errors"
	"math"
	"slices"
	"sort"
	"time"
)

type ChartData struct {
	Labels []string  `json:"labels"`
	Data   []float64 `json:"data"`
	Min    int       `json:"min"`
	Max    int       `json:"max"`
}

type HandicapStore interface {
	GetChartDataForUser(ctx context.Context, userId, limit int) (*ChartData, error)
	GetIndxByUserId(ctx context.Context, userId int) (float64, error)
	SaveHandicap(ctx context.Context, userId int, rounds [20]*Round, date time.Time) error
}

var ErrNotEnoughRounds = errors.New("not enough rounds for handicap")

func (pg PostgresStore) GetChartDataForUser(ctx context.Context, userId, limit int) (*ChartData, error) {
	query := `
	SELECT indx, date
	FROM handicap
	WHERE user_id = $1
	ORDER BY date DESC, id DESC
	LIMIT $2;`

	chart := &ChartData{
		Labels: []string{},
		Data:   []float64{},
		Max:    -999999,
		Min:    999999,
	}
	dates := []time.Time{}

	rows, err := pg.db.QueryContext(ctx, query, userId, limit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return chart, nil
		}
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var indx float64
		var date time.Time
		if err = rows.Scan(&indx, &date); err != nil {
			return nil, err
		}

		chart.Data = append(chart.Data, indx)
		if float64(chart.Max) < indx {
			chart.Max = int(math.Ceil(float64(indx)))
		}
		if float64(chart.Min) > indx {
			chart.Min = int(math.Floor(float64(indx)))
		}

		dates = append(dates, date)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	slices.Reverse(chart.Data)

	latestMonth := ""
	for i := len(dates) - 1; i >= 0; i-- {
		month := dates[i].Month().String()
		if latestMonth != month {
			latestMonth = month
			chart.Labels = append(chart.Labels, month)
		} else {
			chart.Labels = append(chart.Labels, "")
		}
	}

	// Add padding to min + max
	chart.Max = chart.Max + 1
	chart.Min = chart.Min - 1

	return chart, nil
}

func (pg PostgresStore) GetIndxByUserId(ctx context.Context, userId int) (float64, error) {
	query := `
	SELECT indx 
	FROM handicap
	WHERE user_id = $1
	ORDER BY date DESC, id DESC
	LIMIT 1;`

	var indx float64
	err := pg.db.QueryRowContext(ctx, query, userId).Scan(&indx)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}

	return indx, err
}

func (pg PostgresStore) SaveHandicap(ctx context.Context, userId int, rounds [20]*Round, date time.Time) error {
	query := `
	INSERT INTO handicap(
		user_id, 
		indx, 
		date
	) VALUES (
		$1, ROUND($2, 1), $3
	);`

	indx, err := CalculateHandicap(rounds)
	if err != nil {
		if errors.Is(err, ErrNotEnoughRounds) {
			return nil
		}
		return err
	}
	tx := ctx.Value(transactionKey).(*sql.Tx)

	if tx != nil {
		_, err = tx.ExecContext(ctx, query, userId, indx, date)
		return err
	}
	_, err = pg.db.ExecContext(ctx, query, userId, indx, date)

	return err
}

func CalculateHandicap(rounds [20]*Round) (float64, error) {
	handicap := float64(0)
	roundsCount := countRounds(rounds)

	if roundsCount < 3 {
		return 0, ErrNotEnoughRounds
	}

	var scoreDiffs []float64

	for i := 0; i < roundsCount; i++ {
		round := rounds[i]
		d, err := round.CalcDifferential()
		if err != nil {
			return -1, err
		}
		scoreDiffs = append(scoreDiffs, d)
	}

	sort.Slice(scoreDiffs, func(i, j int) bool {
		return scoreDiffs[i] < scoreDiffs[j]
	})

	switch roundsCount {
	case 3:
		handicap = scoreDiffs[0] - 2
	case 4:
		handicap = scoreDiffs[0] - 1
	case 5:
		handicap = scoreDiffs[0]
	case 6:
		handicap = average(scoreDiffs[:1]) - 1
	case 7, 8:
		handicap = average(scoreDiffs[:1])
	case 9, 10, 11:
		handicap = average(scoreDiffs[:2])
	case 12, 13, 14:
		handicap = average(scoreDiffs[:3])
	case 15, 16:
		handicap = average(scoreDiffs[:4])
	case 17, 18:
		handicap = average(scoreDiffs[:5])
	case 19:
		handicap = average(scoreDiffs[:6])
	default:
		handicap = average(scoreDiffs[:7])
	}

	return handicap, nil
}

func average(v []float64) float64 {
	total := float64(0)
	for _, v := range v {
		total += v
	}
	return total / float64(len(v))
}

func countRounds(rounds [20]*Round) int {
	count := 0
	for i := range rounds {
		if rounds[i] == nil {
			return count
		}
		count += 1
	}
	return count
}
