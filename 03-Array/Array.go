package main

//Types of Arrays in Go
// 1. Single-Dimensional Arrays
// 2. Multi-Dimensional Arrays
// 3. Jagged Arrays
// 4. Dynamic Arrays
// 5. Associative Arrays
// 6. Sparse Arrays
// 7. Fixed-Size Arrays
// 8. Variable-Size Arrays
// 9. Homogeneous Arrays
// 10. Heterogeneous Arrays

import "fmt"

func main() {

	var studentNames [5]string = [5]string{"Ibn Abbas", "Ali", "Ahmed", "Mohammed", "Abdullah"}
	var studentAges [5]int = [5]int{36, 25, 30, 28, 22}
	var studentGrades [5]float64 = [5]float64{3.5, 3.8, 3.2, 3.9, 3.0}

	fmt.Println("Student Names: ", studentNames[1])
	fmt.Println("Student Ages: ", studentAges[1])
	fmt.Println("Student Grades: ", studentGrades[1])

}
