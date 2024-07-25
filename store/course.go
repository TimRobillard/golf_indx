package store

import (
	"context"
	"strconv"
	"strings"

	"github.com/lib/pq"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CourseStore interface {
	CreateCourse(ctx context.Context, name, thumbnail string, rating, slope float64, front, back [9]int) (*Course, error)
	GetCourseById(context.Context, int) (*Course, error)
	SearchCourses(context.Context, string) ([]*UICourse, error)
}

type Course struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Front     [9]int  `json:"front"`
	Back      [9]int  `json:"back"`
	Slope     float64 `json:"slope"`
	Rating    float64 `json:"rating"`
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

func (pg PostgresStore) CreateCourse(ctx context.Context, name, thumbnail string, rating, slope float64, front, back [9]int) (*Course, error) {
	name = strings.ToLower(name)

	query := `INSERT INTO course(name, thumbnail, front, back, rating, slope)
	VALUES($1, $2, $3, $4, $5, $6) 
	RETURNING id;`

	var id int

	if err := pg.db.QueryRowContext(ctx, query, name, thumbnail, pq.Array(front), pq.Array(back), rating, slope).Scan(&id); err != nil {
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

func (pg PostgresStore) GetCourseById(ctx context.Context, id int) (*Course, error) {
	query := `
	SELECT 
		id,
		name,
		front,
		back,
		rating,
		slope,
		thumbnail
	FROM course
	WHERE id = $1;`

	rows, err := pg.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	course := &Course{}

	for rows.Next() {
		var f, b []int32
		if err = rows.Scan(&course.Id, &course.Name, (*pq.Int32Array)(&f), (*pq.Int32Array)(&b), &course.Rating, &course.Slope, &course.Thumbnail); err != nil {
			return nil, err
		}

		err = convertNine(f, &course.Front, err)
		err = convertNine(b, &course.Back, err)
		if err != nil {
			return nil, err
		}
	}

	return course, nil
}

func (pg PostgresStore) SearchCourses(ctx context.Context, text string) ([]*UICourse, error) {
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
	rows, err := pg.db.QueryContext(ctx, query, text)

	if err != nil {
		return courses, err
	}

	defer rows.Close()
	caser := cases.Title(language.Und)

	for rows.Next() {
		var c UICourse
		var rank float64
		if err := rows.Scan(&c.Id, &c.Name, &c.Thumbnail, &rank); err != nil {
			return courses, err
		}

		c.Name = caser.String(c.Name)

		courses = append(courses, &c)
	}

	if err = rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}
