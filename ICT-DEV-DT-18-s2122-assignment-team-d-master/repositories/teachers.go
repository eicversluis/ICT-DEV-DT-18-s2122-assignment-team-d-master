package repositories

import (
	. "eindopdracht/types"
	"gorm.io/gorm/clause"
)

func AuthenticateTeacher(username string, password string) (Teacher, error) {
	var result Teacher
	//err := connection().Exec("SELECT id, name FROM teachers WHERE username = '" + username + "' AND password = '" + password + "'").First(&result).Error
	err := connection().Where("username = '" + username + "'").Where("password = '" + password + "'").First(&result).Error
	return result, err
}

func GetTeachers(Preloads ...string) []Teacher {
	var result []Teacher
	query := connection().Preload(clause.Associations)
	for _, preload := range Preloads {
		query.Preload(preload)
	}
	query.Find(&result)
	return result
}
