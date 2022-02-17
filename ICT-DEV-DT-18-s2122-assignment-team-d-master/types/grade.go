package types

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	Value     float64 `json:"value"`
	ModuleID  uint    `json:"module" gorm:"index:one_grade,unique"`
	Module    Module  `json:"-"`
	StudentID uint    `json:"student" gorm:"index:one_grade,unique"`
	Student   Student `json:"-"`
}
