package db

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Write(id string, data string) {
	ioutil.WriteFile(fmt.Sprintf(`database/%s.json`, id), []byte(data), 0644)
}

func Read(id string) []byte {
	content, err := ioutil.ReadFile(fmt.Sprintf(`database/%s.json`, id))
	if err != nil {
		return nil
	}
	return content
}

func Remove(id string) {
	os.Remove(fmt.Sprintf(`database/%s.json`, id))
}

func Find(id string) bool {
	_, err := os.Stat(fmt.Sprintf(`database/%s.json`, id))
	return err == nil
}
