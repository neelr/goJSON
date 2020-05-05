package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/neelr/gojson/db"
)

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func cors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Method", "*")

}

func logHandle(w http.ResponseWriter, r *http.Request) {
	cors(&w)
	logs, _ := ioutil.ReadFile("logs.json")
	w.Header().Set("Content-Type", "application/json")
	w.Write(logs)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	cors(&w)
	// Log the request
	content, _ := ioutil.ReadFile(`logs.json`)
	logs, _ := simplejson.NewJson(content)
	currentTime := time.Now().Format("2006-01-02")
	if log, ok := logs.CheckGet(currentTime); ok {
		logs.Set(currentTime, log.MustInt()+1)
	} else {
		logs.Set(currentTime, 1)
	}
	byteLogs, _ := logs.MarshalJSON()
	ioutil.WriteFile(`logs.json`, byteLogs, 0644)

	// End of Logging System
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}
	name := strings.Replace(filepath.Clean(r.URL.Path), "/api/", "", 1)
	keys := strings.Split(name, "/")
	if r.Method == "POST" || r.Method == "PUT" {
		if isJSON(string(body)) {
			var js *simplejson.Json
			if db.Find(keys[0]) {
				js, _ = simplejson.NewJson(db.Read(keys[0]))
			} else {
				js = simplejson.New()
			}
			bodyjs, _ := simplejson.NewJson(body)
			main := js
			before := js

			// Navigate the JSON from the url
			for i := 1; i <= len(keys)-1; i++ {
				if data, ok := js.CheckGet(keys[i]); ok {
					js = data
					if i > 1 {
						before = js.Get(keys[i-1])
					}
				} else {
					js.Set(keys[i], simplejson.New().Interface())
					js = js.Get(keys[i])
				}
			}

			stringChecker, _ := js.MarshalJSON()
			if !strings.Contains(string(stringChecker), "{") {
				before.Set(keys[len(keys)-1], simplejson.New().Interface())
				js = before.Get(keys[len(keys)-1])
			}
			for k, v := range bodyjs.MustMap() {
				js.Set(k, v)
			}
			jstring, _ := main.MarshalJSON()
			db.Write(keys[0], string(jstring))
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Done"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Malformatted Request"))
		return
	} else if r.Method == "DELETE" {
		if len(keys) > 1 {
			js, _ := simplejson.NewJson(db.Read(keys[0]))
			main := js
			fmt.Println(keys)
			for i := 1; i < len(keys)-1; i++ {
				fmt.Println(keys[i])
				if data, ok := js.CheckGet(keys[i]); ok {
					js = data
				} else {
					js.Set(keys[i], simplejson.New().Interface())
					js = js.Get(keys[i])
				}
			}
			js.Del(keys[len(keys)-1])
			jstring, _ := main.MarshalJSON()
			fmt.Println(string(jstring))
			db.Write(keys[0], string(jstring))
			w.WriteHeader(204)
			w.Write([]byte("Done"))
			return
		}
		db.Remove(name)
		w.WriteHeader(204)
		w.Write([]byte("Done"))
		return
	}
	data := db.Read(keys[0])
	if data != nil {
		if len(keys) > 1 {
			js, _ := simplejson.NewJson(data)
			fmt.Println(keys)
			for i := 1; i <= len(keys)-1; i++ {
				fmt.Println(keys[i])
				if data, ok := js.CheckGet(keys[i]); ok {
					js = data
				} else {
					w.WriteHeader(http.StatusNotFound)
					w.Write([]byte("Record not Found"))
					return
				}
			}
			w.Header().Set("Content-Type", "application/json")
			jstring, _ := js.MarshalJSON()
			w.Write([]byte(jstring))
			return
		}
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
	http.HandleFunc("/logs", logHandle)
	fmt.Println("Up on port 3000!")
	http.ListenAndServe(":3000", nil)
}
