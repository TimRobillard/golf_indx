package store

type Course struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Front  [9]int  `json:"front"`
	Back   [9]int  `json:"back"`
	Slope  float32 `json:"slope"`
	Rating float32 `json:"rating"`
}

func (c Course) FrontPar() int {
	t := 0
	for _, par := range c.Front {
		t += par
	}
	return t
}

func (c Course) BackPar() int {
	t := 0
	for _, par := range c.Back {
		t += par
	}
	return t
}

func (c Course) CalculateNine(scores [9]int) int {
	t := 0
	for _, s := range scores {
		t += s
	}
	return t
}
