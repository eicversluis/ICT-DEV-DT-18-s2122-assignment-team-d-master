package repositories

import (
	. "eindopdracht/types"
	"gorm.io/gorm/clause"
	"log"
)

func GetStudents(Preloads ...string) []Student {
	var Students []Student
	query := connection().Preload(clause.Associations)
	for _, preload := range Preloads {
		query.Preload(preload)
	}
	query.Find(&Students)
	return Students
}

func GetStudentById(Id int, Preloads ...string) Student {
	var Student Student
	query := connection().Preload(clause.Associations)
	for _, preload := range Preloads {
		query.Preload(preload)
	}
	query.Where("id = ?", Id).First(&Student)
	return Student
}

func CreateStudent(student Student) error {
	if err := connection().Create(&student).Error; err != nil {
		return err
	}
	return nil
}

func GetAverageGrade(Id int) int {
	var Student Student
	err := connection().Preload("Grades").Where("id = ?", Id).First(&Student)
	if err != nil{
		log.Println(err)
	}
	AvgGrade := 0
	for _, grade := range Student.Grades {
		AvgGrade += int	(grade.Value)
	}
	AvgGrade = AvgGrade / len(Student.Grades)
	return AvgGrade
}
