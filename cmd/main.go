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
	myGroup := storage.NewGroup(s)

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
		stud, err := student.NewStudent(name, age, grade)
		if err != nil {
			fmt.Println(err)
		}
		myGroup.Put(stud)
	}
	fmt.Println("The end! List students:")

	list := myGroup.GetAll()

	cnt := 1
	for _, v := range list {
		fmt.Printf("%v. %v. %v. %v. \n", cnt, v.GetName(), v.GetAge(), v.GetGrade())
		cnt++
	}

}
