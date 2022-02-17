package types

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Groups   []Group
}
