package db

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Write to the database dir
func Write(id string, data string) {
	ioutil.WriteFile(fmt.Sprintf(`database/%s.json`, id), []byte(data), 0644)
}

// Read the database dir
func Read(id string) []byte {
	content, err := ioutil.ReadFile(fmt.Sprintf(`database/%s.json`, id))
	if err != nil {
		return nil
	}
	return content
}

// Remove File
func Remove(id string) {
	os.Remove(fmt.Sprintf(`database/%s.json`, id))
}

// Find a file
func Find(id string) bool {
	_, err := os.Stat(fmt.Sprintf(`database/%s.json`, id))
	return err == nil
}
