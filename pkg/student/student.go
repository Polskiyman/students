package student

import (
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Age   int
	Grade int
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
		Name:  name,
		Age:   age,
		Grade: grade,
	}, nil
}
