package store

import (
	"strconv"
	"strings"

	"github.com/lib/pq"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CourseStore interface {
	CreateCourse(name, thumbnail string, rating, slope float32, front, back [9]int) (*Course, error)
	SearchCourses(text string) ([]*UICourse, error)
}

type Course struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Front     [9]int  `json:"front"`
	Back      [9]int  `json:"back"`
	Slope     float32 `json:"slope"`
	Rating    float32 `json:"rating"`
	Thumbnail string  `json:"thumbnail"`
}

type UICourse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Par       string `json:"par"`
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

func (c Course) GetName() UICourse {
	caser := cases.Title(language.Und)
	return UICourse{
		Id:        c.Id,
		Name:      caser.String(c.Name),
		Thumbnail: c.Thumbnail,
		Par:       strconv.Itoa(c.Par()),
	}
}

func (c Course) Par() int {
	return c.FrontPar() + c.BackPar()
}

func (pg PostgresStore) CreateCourse(name, thumbnail string, rating, slope float32, front, back [9]int) (*Course, error) {
	name = strings.ToLower(name)

	query := `INSERT INTO course(name, thumbnail, front, back, rating, slope)
	VALUES($1, $2, $3, $4, $5, $6) 
	RETURNING id;`

	var id int

	if err := pg.db.QueryRow(query, name, thumbnail, pq.Array(front), pq.Array(back), rating, slope).Scan(&id); err != nil {
		return nil, err
	}

	return &Course{
		Id:        id,
		Name:      name,
		Thumbnail: thumbnail,
		Front:     front,
		Back:      back,
		Rating:    rating,
		Slope:     slope,
	}, nil
}

func (pg PostgresStore) SearchCourses(text string) ([]*UICourse, error) {
	query := `SELECT 
		id,
		name,
		thumbnail,
		similarity(name, $1) as rank
		FROM course
		ORDER BY rank DESC
		LIMIT 10;
	`

	var courses []*UICourse
	rows, err := pg.db.Query(query, text)

	if err != nil {
		return courses, err
	}

	defer rows.Close()

	for rows.Next() {
		var c UICourse
		var rank float32
		if err := rows.Scan(&c.Id, &c.Name, &c.Thumbnail, &rank); err != nil {
			return courses, err
		}

		courses = append(courses, &c)
	}

	if err = rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}
