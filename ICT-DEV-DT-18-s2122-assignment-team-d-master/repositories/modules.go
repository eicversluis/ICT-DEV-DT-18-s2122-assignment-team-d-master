package repositories

import (
	. "eindopdracht/types"
	"gorm.io/gorm/clause"
)

func GetModules(Preloads ...string) []Module {
	var result []Module
	query := connection().Preload(clause.Associations)
	for _, preload := range Preloads {
		query.Preload(preload)
	}
	query.Find(&result)
	return result
}

func GetModule(id int, Preloads ...string) (Module, error) {
	var result Module
	query := connection().Preload(clause.Associations).Where("id = ?", id)
	for _, preload := range Preloads {
		query.Preload(preload)
	}
	if err := query.First(&result).Error; err != nil {
		return Module{}, err
	}
	return result, nil
}

func CreateModule(module Module) error {
	if err := connection().Create(&module).Error; err != nil {
		return err
	}
	return nil
}
