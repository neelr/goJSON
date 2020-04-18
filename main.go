package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/neelr/gojson/db"
)

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}
	name := strings.Replace(r.URL.Path, "/api/", "", 1)
	if r.Method == "POST" || r.Method == "PUT" {
		if isJSON(string(body)) {
			db.Write(name, string(body))
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Done"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Malformatted Request"))
		return
	} else if r.Method == "DELETE" {
		db.Remove(name)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Done"))
		return
	}
	data := db.Read(name)
	if data != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Record not Found"))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/", indexHandle)
	fmt.Println("Up on port 3000!")
	http.ListenAndServe(":3000", nil)
}
