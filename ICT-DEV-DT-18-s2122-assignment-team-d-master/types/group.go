package types

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name        string    `json:"name"`
	TeacherID   uint      `json:"teacher"`
	Description string    `json:"description"`
	Teacher     Teacher   `json:"-"`
	Students    []Student `json:"-"`
}
