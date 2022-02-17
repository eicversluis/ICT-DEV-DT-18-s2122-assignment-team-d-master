package types

import "gorm.io/gorm"

type Module struct {
	gorm.Model
	Name   string  `json:"name"`
	Grades []Grade `json:"-"`
}
