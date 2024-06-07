package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tealeg/xlsx"
)

type Character struct {
	Name   string         `json:"name"`
	Level  int            `json:"level"`
	HP     int            `json:"hp"`
	ExpVal int            `json:"exp_val"`
	Stats  map[string]int `json:"stats"`
}

func main() {
	http.HandleFunc("/api/characters", handleCharacters)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCharacters(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		characters, err := readExcelFile("characters.xlsx")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Pass the characters to the Python API
		jsonData, err := json.Marshal(characters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := http.Post("http://localhost:5000/api/characters", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, string(body))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func readExcelFile(filename string) ([]Character, error) {
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	sheet := xlFile.Sheets[0]
	var characters []Character

	for _, row := range sheet.Rows[1:] {
		var character Character
		err := row.ReadStruct(&character)
		if err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}

	return characters, nil
}
