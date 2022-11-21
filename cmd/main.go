package main

import (
	"./pkg/storage"
	"./pkg/student"
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	myStudMap := storage.Group{}
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
		myStudMap.
		(storage.Put(stud))
	}
	fmt.Println("The end! List students:")
	cnt := 1
	for _, v := range myStudMap.(storage.studs) {
		fmt.Printf("%v. %v\n", cnt, myStudMap.get(v.name))
		cnt++
	}
}
