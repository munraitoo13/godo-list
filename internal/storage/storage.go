package storage

import (
	"encoding/json"
	"os"
)

// storage struct
type Storage[T any] struct {
	FileName string
}

// create new storage
func NewStorage[T any](fileName string) *Storage[T] {
	// returns a pointer to the storage struct
	return &Storage[T]{FileName: fileName}
}

// save data to file
func (s *Storage[T]) SaveTasks(data T) error {
	// convert data to json
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// write data to file
	return os.WriteFile(s.FileName, fileData, 0644)
}

// read data from file
func (s *Storage[T]) LoadTasks(data *T) error {
	// read file data
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return nil
	}

	// unmarshal fileData to data
	return json.Unmarshal(fileData, data)
}
