package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"students/pkg/storage"
	"students/pkg/student"
)

func main() {
	s := make(map[string]*student.Student)
	myGroup := storage.New(s)

	in := bufio.NewReader(os.Stdin)

	for {
		line, err := in.ReadString('\n')
		if err == io.EOF {
			break
		}

		name, age, grade, err := student.Parse(line)
		if err != nil {
			fmt.Println(err)
			continue
		}

		stud := student.New(name, age, grade)

		myGroup.Put(stud)
	}

	fmt.Println("The end! List students:")

	list := myGroup.GetAll()

	for i, v := range list {
		fmt.Printf("%v. %v. %v. %v. \n", i+1, v.Name(), v.Age(), v.Grade())
	}
}
