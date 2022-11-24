package storage

import (
	"students/pkg/student"
)

type Group struct {
	studs map[string]*student.Student
}

func NewGroup(s map[string]*student.Student) Group {
	return Group{
		studs: s,
	}
}

func (p Group) Put(stud *student.Student) {
	p.studs[stud.Name()] = stud
}

func (p Group) GetStudent(name string) *student.Student {
	stud := p.studs[name]
	return stud
}
func (p Group) GetAll() (res []*student.Student) {
	for _, s := range p.studs {
		res = append(res, s)
	}
	return
}
