// Package tools for working with data
package tools

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"bzabek/task-tracker/model"
)

const filename = "db.json"

func CreateFileIfNotExist() {
	_, error := os.Stat(filename)
	if errors.Is(error, os.ErrNotExist) {
		SaveChanges(
			model.DB{
				NextID: 1,
				Tasks:  make(map[int64]model.Task),
			})
	}
}

func SaveChanges(toDoList model.DB) error {
	j, errors := json.MarshalIndent(toDoList, "", "  ")

	if errors != nil {
		log.Println(errors)
		return errors
	}

	errors = os.WriteFile(filename, j, 0o666)
	if errors != nil {
		log.Println(errors)
	}

	return errors
}

func GetDB() (model.DB, error) {
	file, errors := os.ReadFile(filename)
	result := model.DB{}

	if errors != nil {
		log.Fatal(errors)
		return result, errors
	}

	errors = json.Unmarshal(file, &result)

	if errors != nil {
		log.Println(errors)
	}

	return result, errors
}
