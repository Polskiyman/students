package student

import (
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

func Parse(line string) (string, int, int, error) {
	fields := strings.Fields(line)
	if len(fields) < 3 {
		return "", 0, 0, fmt.Errorf("enter name, grade, age! Please, try again")
	}
	name := fields[0]
	age, err := strconv.Atoi(fields[1])
	if err != nil {
		return "", 0, 0, fmt.Errorf("error processing age:%s", err)
	}
	grade, err := strconv.Atoi(fields[2])
	if err != nil {
		return "", 0, 0, fmt.Errorf("error processing grade:%s", err)
	}
	return name, age, grade, nil
}

func NewStudent(name string, age, grade int) (*Student, error) {
	return &Student{
		name:  name,
		age:   age,
		grade: grade,
	}, nil
}

func (s *Student) SetName(name string) {
	s.name = name
}
func (s *Student) GetName() string {
	return s.name
}

func (s *Student) SetAge(age int) {
	s.age = age
}
func (s *Student) GetAge() int {
	return s.age
}

func (s *Student) SetGrade(grade int) {
	s.grade = grade
}
func (s *Student) GetGrade() int {
	return s.grade
}
