package db

import (
	"github.com/replit/database-go"
)

// Write to the database dir
func Write(id string, data string) {
	database.Set(id, data)
}

// Read the database dir
func Read(id string) []byte {
	content, err := database.Get(id)
	if err != nil {
		return nil
	}
	return []byte(content)
}

// Remove File
func Remove(id string) {
	database.Delete(id)
}

// Find a file
func Find(id string) bool {
	_, err := database.Get(id)
	return err == nil
}
