package types

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name    string  `json:"name"`
	GroupID uint    `json:"group"`
	Group   Group   `json:"-"`
	Grades  []Grade `json:"-"`
}

func (s *Student) GetAvgGrade() float64 {
	AvgGrade := 0.0
	for _, grade := range s.Grades {
		AvgGrade += grade.Value
	}
	AvgGrade = AvgGrade / float64(len(s.Grades))
	return AvgGrade
}
