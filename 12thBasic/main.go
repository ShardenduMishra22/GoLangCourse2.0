package main

import "fmt"

func main() {
	// GradesMap := make(map[string]int)

	// GradesMap["Math"] = 90
	// GradesMap["Science"] = 80
	// GradesMap["English"] = 70

	// fmt.Println(GradesMap)
	// fmt.Println("Grades of Math: ", GradesMap["Math"])

	// delete(GradesMap, "Science")
	// fmt.Println(GradesMap)

	// GradeInMath, exists := GradesMap["Math"]
	// if exists {
	// 	fmt.Println("Grade in Math: ", GradeInMath)
	// } else {
	// 	fmt.Println("Subject Doesent Exist")
	// }

	StudentGrades := map[string]int {
		"Math": 90,
		"Science": 80,
		"English": 70,
		"History": 60,
	}

	fmt.Println(StudentGrades)
}