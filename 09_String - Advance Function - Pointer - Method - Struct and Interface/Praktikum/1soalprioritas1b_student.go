package main

import (
	"fmt"
	"strconv"
)

type Student struct {

	name string
	score int
}

func NewStudent(name string, score int) *Student {

	student := Student{name, score}

	return &student
}

func (this *Student)GetName() string {

	return this.name
}

func (this *Student)GetScore() int {

	return this.score
}

func InputStudent(index int) *Student {

	index += 1

	var name string
	var score string
	
	fmt.Printf("Input %d Student`s Name: ", index)
	fmt.Scan(&name)
	fmt.Printf("Input %d Student`s Score: ", index)
	fmt.Scan(&score)

	in, err := strconv.ParseInt(score, 0, 32)

	if err == nil {

		student := NewStudent(name, int(in))
		return student
	}

	return nil
}

func main() {

	fmt.Println("Hello World, Hello Golang World!")
	var students []Student

	var maxStudent string
	fmt.Printf("Input Max Student: ")
	fmt.Scan(&maxStudent)
	in, err := strconv.ParseInt(maxStudent, 0, 32)
	if err == nil {

		n := int(in)

		maxScore := NewStudent("Unknown", -1) 
		minScore := NewStudent("Unknown", -1)
		var avgScore int

		for i := 0; i < n; i++ {

			student := InputStudent(i)
	
			students = append(students, *student)

			score := student.GetScore()

			if maxScore.GetScore() == -1 {

				maxScore = student
				
			} else if (maxScore.GetScore() < score) {

				maxScore = student
			}

			if minScore.GetScore() == -1 {

				minScore = student
				
			} else if (minScore.GetScore() > score) {

				minScore = student
			}
			
			avgScore += score
			
		}

		fmt.Println(students)
		fmt.Printf("Average Score: %f\n", avgScore / n)
		fmt.Printf("Min Score of Students: %s (%d)\n", minScore.GetName(), minScore.GetScore())
		fmt.Printf("Max Score of Students: %s (%d)\n", maxScore.GetName(), maxScore.GetScore())
	}
}
