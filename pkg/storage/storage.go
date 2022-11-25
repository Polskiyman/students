package storage

import (
	"students/pkg/student"
)

type Storage struct {
	studs map[string]*student.Student
}

func New(s map[string]*student.Student) Storage {
	return Storage{
		studs: s,
	}
}

func (p Storage) Put(stud *student.Student) {
	p.studs[stud.Name()] = stud
}

func (p Storage) GetStudent(name string) *student.Student {
	stud := p.studs[name]
	return stud
}
func (p Storage) GetAll() (res []*student.Student) {
	for _, s := range p.studs {
		res = append(res, s)
	}
	return
}
