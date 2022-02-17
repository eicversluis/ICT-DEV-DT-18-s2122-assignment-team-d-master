package repositories

import (
	. "eindopdracht/types"
	"gorm.io/gorm/clause"
)

func InsertGrades(grades []Grade) error {
	if err := connection().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "module_id"}, {Name: "student_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(&grades).Error; err != nil {
		return err
	}
	return nil
}

func GetModuleGradesForGroup(ModuleId int, GroupId int) ([]Grade, error) {
	var result []Grade
	if err := connection().Preload(clause.Associations).Preload("Student", "group_id = ?", GroupId).Where("module_id = ?", ModuleId).Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
