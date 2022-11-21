package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Polskiyman/students/pkg/storage"
	"github.com/Polskiyman/students/pkg/student"
)

func main() {
	myGrup := storage.Group{}
	myGrup.Studs = make(map[string]*student.Student)
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
		myGrup.Put(stud)
	}
	fmt.Println("The end! List students:")
	cnt := 1
	for _, v := range myGrup.Studs {
		fmt.Printf("%v. %v\n", cnt, myGrup.Get(v.Name))
		cnt++
	}
}
