package seed

import (
	"log"

	"github.com/TimRobillard/handicap_tracker/store"
)

type courseInput struct {
	Name      string
	Thumbnail string
	Rating    float32
	Slope     float32
	Front     [9]int
	Back      [9]int
}

type roundInput struct {
	User   *store.UIUser
	Course store.Course
	Front  [9]int
	Back   [9]int
	Date   string
}

func createCourse(pg *store.PostgresStore, err error, input courseInput) (*store.Course, error) {
	if err != nil {
		return nil, err
	}
	return pg.CreateCourse(input.Name, input.Thumbnail, input.Rating, input.Slope, input.Front, input.Back)
}

func createRound(pg *store.PostgresStore, err error, input roundInput) (*store.Round, error) {
	if err != nil {
		return nil, err
	}
	return pg.CreateRound(input.User, input.Course, input.Front, input.Back, input.Date)
}

func Seed(db *store.PostgresStore) {
	user, err := db.CreateUser("albatrooss", "$2a$14$MhmlE7Na84k55C.5R6r/9.TrEgmlhHgI/Ukswy47xeJP5XE1gvxz")
	manderleyCN, err := createCourse(
		db,
		err,
		courseInput{
			"Manderley Central North",
			"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYsa5s9fz-agjYOZtBTJDSaDV_78gxOiRTQw&usqp=CAU",
			110,
			67.1,
			[9]int{5, 3, 5, 3, 4, 3, 5, 4, 4},
			[9]int{5, 4, 4, 3, 5, 4, 3, 4, 4},
		})
	manderleyNS, err := createCourse(
		db,
		err,
		courseInput{
			Name:      "Manderley North South",
			Front:     [9]int{5, 4, 4, 3, 5, 4, 3, 4, 4},
			Back:      [9]int{4, 4, 3, 4, 4, 3, 5, 3, 5},
			Thumbnail: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYsa5s9fz-agjYOZtBTJDSaDV_78gxOiRTQw&usqp=CAU",
			Slope:     112,
			Rating:    65.3,
		})
	dragonFly, err := createCourse(
		db, err,
		courseInput{
			Name:      "Dragonfly Golf Links",
			Front:     [9]int{4, 4, 4, 4, 3, 5, 3, 4, 5},
			Back:      [9]int{4, 5, 4, 3, 4, 5, 4, 4, 3},
			Thumbnail: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYsa5s9fz-agjYOZtBTJDSaDV_78gxOiRTQw&usqp=CAU",
			Slope:     123,
			Rating:    69.9,
		})
	cedarHill, err := createCourse(db, err,
		courseInput{
			Name:      "Cedarhill",
			Front:     [9]int{4, 4, 5, 4, 3, 4, 4, 4, 3},
			Back:      [9]int{4, 4, 4, 4, 3, 3, 4, 5, 4},
			Thumbnail: "https://stittsvillecentral.ca/wp-content/uploads/amberwood-village-golf-green.jpg",
			Slope:     112,
			Rating:    67.6,
		})
	amberwood, err := createCourse(
		db, err, courseInput{
			Name:      "Amberwood",
			Front:     [9]int{3, 4, 4, 3, 4, 3, 3, 4, 4},
			Thumbnail: "https://stittsvillecentral.ca/wp-content/uploads/amberwood-village-golf-green.jpg",
			Slope:     99,
			Rating:    31.2,
		})

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = createRound(
		db, err, roundInput{
			User:   user.ToUI(),
			Course: *manderleyCN,
			Front:  [9]int{8, 4, 6, 5, 6, 4, 9, 6, 6},
			Back:   [9]int{7, 6, 6, 2, 7, 4, 6, 4, 6},
			Date:   "2024-06-30",
		})
	_, err = createRound(db, err, roundInput{
		Front:  [9]int{5, 5, 8, 5, 2, 5, 4, 5},
		Back:   [9]int{5, 5, 5, 4, 4, 8, 6, 6, 3},
		Course: *dragonFly,
		Date:   "2024-06-24",
		User:   user.ToUI(),
	})
	_, err = createRound(db, err, roundInput{
		Front:  [9]int{7, 5, 5, 8, 3, 4, 6, 7, 5},
		Back:   [9]int{3, 7, 5, 3, 3, 4, 7, 6, 8},
		Course: *cedarHill,
		Date:   "2024-06-24",
		User:   user.ToUI(),
	})
	_, err = createRound(db, err, roundInput{
		Front:  [9]int{3, 5, 3, 4, 6, 4, 3, 6, 6},
		Back:   [9]int{3, 5, 3, 4, 6, 4, 3, 6, 6},
		Course: *amberwood,
		Date:   "2024-06-24",
		User:   user.ToUI(),
	})
	_, err = createRound(db, err, roundInput{
		Front:  [9]int{7, 7, 6, 4, 6, 8, 4, 4, 6},
		Back:   [9]int{5, 6, 4, 7, 5, 5, 6, 3, 7},
		Course: *manderleyNS,
		Date:   "2024-06-24",
		User:   user.ToUI(),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Default().Println("Done seeding")
}
