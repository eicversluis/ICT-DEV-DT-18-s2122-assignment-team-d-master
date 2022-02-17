package repositories

import (
	. "eindopdracht/types"
	"gorm.io/gorm/clause"
)

func GetGroups(Preloads ...string) []Group {
	var result []Group
	query := connection().Preload(clause.Associations)
	for _, preload := range Preloads {
		query.Preload(preload)
	}
	query.Find(&result)
	return result
}

func GetGroup(id int, Preloads ...string) (Group, error) {
	var result Group
	query := connection().Where("ID = ?", id).Preload(clause.Associations)
	for _, preload := range Preloads {
		query.Preload(preload)
	}
	if err := query.First(&result).Error; err != nil {
		return Group{}, err
	}
	return result, nil
}

func CreateGroup(group Group) error {
	if err := connection().Create(&group).Error; err != nil {
		return err
	}
	return nil
}
