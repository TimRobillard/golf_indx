package seed

import (
	"log"

	"github.com/TimRobillard/handicap_tracker/store"
)

func Seed(db *store.PostgresStore) {
	_, err := db.CreateCourse("manderley on the green central north", "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYsa5s9fz-agjYOZtBTJDSaDV_78gxOiRTQw&usqp=CAU", float32(65.3), float32(110), [9]int{5, 3, 5, 3, 4, 3, 4, 4}, [9]int{5, 4, 4, 3, 5, 4, 3, 4, 4})
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = db.CreateCourse("manderley on the green south north", "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYsa5s9fz-agjYOZtBTJDSaDV_78gxOiRTQw&usqp=CAU", float32(32), float32(32), [9]int{4, 4, 3, 4, 4, 3, 5, 3, 5}, [9]int{5, 4, 4, 3, 5, 4, 3, 4, 4})
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = db.CreateCourse("amberwood village golf", "https://stittsvillecentral.ca/wp-content/uploads/amberwood-village-golf-green.jpg", float32(32), float32(32), [9]int{3, 5, 4, 3, 4, 3, 3, 4, 4}, [9]int{})
	// courses, err := db.SearchCourses("man gree")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// for _, course := range courses {
	// 	fmt.Printf("Id %d, Name %s", course.Id, course.Name)
	// 	for _, par := range course.Front {
	// 		fmt.Printf("Par %d", par)
	// 	}
	// }

	if err != nil {
		log.Fatal(err.Error())
	}
}
