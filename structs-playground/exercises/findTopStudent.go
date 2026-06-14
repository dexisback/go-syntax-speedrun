package main

type Student struct {
	Name  string
	Marks int
}

func findTopStudents() Student {
	//input is a slice of type Student
	//output = student with highest marks
	//so iterate through the slice(Student.Marks) using range and find the highest marks
	input := []Student{
		{Name: "Amaan", Marks: 80},
		{Name: "Ayaan", Marks: 81},
	}
	//assume first index as the top student
	topStudent := input[0]

	//now iterate
	for _, student := range input {
		if student.Marks > topStudent.Marks {
			topStudent = student
		}
	}

	return topStudent

}
