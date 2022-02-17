package main

import (
	"eindopdracht/repositories"
	. "eindopdracht/types"
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type JsonData struct {
	Teachers []Teacher `json:"teachers"`
	Groups   []Group   `json:"groups"`
	Students []Student `json:"students"`
	Modules  []Module  `json:"modules"`
	Grades   []Grade   `json:"grades"`
}

func main() {
	if os.Getenv("APP_MODE") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	var Data JsonData
	file, _ := os.ReadFile("./seeder/seed.json")
	json.Unmarshal(file, &Data)
	if repositories.Connected() {
		repositories.Create(Data.Teachers)
		repositories.Create(Data.Groups)
		repositories.Create(Data.Modules)
		repositories.Create(Data.Students)
		repositories.Create(Data.Grades)
	}
}
