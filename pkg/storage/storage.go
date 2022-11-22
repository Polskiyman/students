package storage

import (
	"students/pkg/student"
)

type Group struct {
	Studs map[string]*student.Student
}

func (p Group) Put(stud *student.Student) {
	p.Studs[stud.Name] = stud
}

func (p Group) Get(name string) *student.Student {
	stud := p.Studs[name]
	return stud
}
