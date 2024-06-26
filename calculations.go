package main

import (
	"fmt"
	"sort"

	"github.com/TimRobillard/handicap_tracker/store"
)

// Calculates the Score Differential for an 18 hole score
//
// - slope Slope Rating
//
// - score Adjusted gross score
//
// - rating Course Rating
//
// - pcc PCC adjustment
func Calc18HoleDifferential(slope, score, rating, pcc float32) float32 {
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
func Calc9HoleDifferential(slope, score, rating, pcc float32) float32 {
	return (113 / slope) * (score - rating - (pcc / 2)) * 2 // should be + expected score, cannot find this calculation
}

type Round struct {
	holes  []int
	course *store.Course
}

func (r Round) CalcScore() float32 {
	score := 0
	for _, h := range r.holes {
		score += h
	}
	return float32(score)
}

func (r Round) CalcDifferential() (float32, error) {
	if len(r.course.Back) != 9 && len(r.course.Front) == 9 {
		return Calc9HoleDifferential(r.course.Slope, r.CalcScore(), r.course.Rating, 0), nil
	}
	if len(r.course.Front) == 9 && len(r.course.Back) == 9 {

		return Calc18HoleDifferential(r.course.Slope, r.CalcScore(), r.course.Rating, 0), nil
	}

	return -1, fmt.Errorf("invalid number of holes: %d", len(r.course.Front)+len(r.course.Back))
}

func average(v []float32) float32 {
	total := float32(0)
	for _, v := range v {
		total += v
	}
	return total / float32(len(v))
}

func CalculateHandicap(rounds []Round) (float32, error) {
	handicap := float32(0)
	roundsCount := len(rounds)

	if roundsCount < 3 {
		return -1, fmt.Errorf("minimum 3 rounds needed. round count: %d", roundsCount)
	}
	if roundsCount > 20 {
		return -1, fmt.Errorf("too many rounds.  %d/20 max rounds", roundsCount)
	}

	var scoreDiffs []float32

	for _, s := range rounds {
		d, err := s.CalcDifferential()
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
	case 7:
	case 8:
		handicap = average(scoreDiffs[:1])
	case 9:
	case 10:
	case 11:
		handicap = average(scoreDiffs[:2])
	case 12:
	case 13:
	case 14:
		handicap = average(scoreDiffs[:3])
	case 15:
	case 16:
		handicap = average(scoreDiffs[:4])
	case 17:
	case 18:
		handicap = average(scoreDiffs[:5])
	case 19:
		handicap = average(scoreDiffs[:6])
	default:
		handicap = average(scoreDiffs[:7])
	}

	return handicap, nil
}

func Test() {
	manderleyCN := store.Course{
		Id: 1,
		Name:   "Manderley Central North",
		Front:  [9]int{5, 3, 5, 3, 4, 3, 5, 4, 4},
		Back:   [9]int{5, 4, 4, 3, 5, 4, 3, 4, 4},
		Slope:  110,
		Rating: 67.1,
	}
	manderleyNS := store.Course{
		Id: 2,
		Name:   "Manderley North South",
		Front:  [9]int{5, 4, 4, 3, 5, 4, 3, 4, 4},
		Back:   [9]int{4, 4, 3, 4, 4, 3, 5, 3, 5},
		Slope:  112,
		Rating: 65.3,
	}
	dragonFly := store.Course{
		Id: 3,
		Name:   "Dragonfly Golf Links",
		Front:  [9]int{4, 4, 4, 4, 3, 5, 3, 4, 5},
		Back:   [9]int{4, 5, 4, 3, 4, 5, 4, 4, 3},
		Slope:  123,
		Rating: 69.9,
	}
	cedarHill := store.Course{
		Id: 4,
		Name:   "Cedarhill",
		Front:  [9]int{4, 4, 5, 4, 3, 4, 4, 4, 3},
		Back:   [9]int{4, 4, 4, 4, 3, 3, 4, 5, 4},
		Slope:  112,
		Rating: 67.6,
	}
	amberwood := store.Course{
		Id: 5,
		Name:   "Amberwood",
		Front:  [9]int{3, 4, 4, 3, 4, 3, 3, 4, 4},
		Slope:  99,
		Rating: 31.2,
	}
	round1 := Round{
		holes:  []int{8, 4, 6, 5, 6, 4, 9, 6, 6, 7, 6, 6, 2, 7, 4, 6, 4, 6},
		course: &manderleyCN,
	}
	round2 := Round{
		holes:  []int{5, 5, 8, 5, 2, 5, 4, 5, 7, 5, 5, 5, 4, 4, 8, 6, 6, 3},
		course: &dragonFly,
	}
	round3 := Round{
		holes:  []int{7, 5, 5, 8, 3, 4, 6, 7, 5, 3, 7, 5, 3, 3, 4, 7, 6, 8},
		course: &cedarHill,
	}
	round4 := Round{
		holes:  []int{3, 5, 3, 4, 6, 4, 3, 6, 6},
		course: &amberwood,
	}
	round5 := Round{
		holes:  []int{7, 7, 6, 4, 6, 8, 4, 4, 6, 5, 6, 4, 7, 5, 5, 6, 3, 7},
		course: &manderleyNS,
	}
	// round6 := Round{
	// 	holes:  []int{5, 8, 5, 4, 6, 4, 4, 4, 7},
	// 	course: &amberwood,
	// }

	handicap, err := CalculateHandicap([]Round{round1, round2, round3, round4, round5}) //  round6

	if err != nil {
		panic(err)
	}
	fmt.Printf("Diff: %.1f\n\n", handicap)
}
