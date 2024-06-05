package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Character struct {
	ID    int
	Name  string
	Level int
	HP    int
	MP    int
	Str   int
	End   int
	Int   int
	Cha   int
	Dex   int
	Wis   int
}

func main() {
	//cli arg
	fileName := flag.String("file", "data.xlsx", "Excel file to read from")
	flag.Parse()

	// file path shit
	filePath := filepath.Join("data", *fileName)
	fmt.Printf("Attempting to open file: %s\n", filePath) // Debug print

	// does the bitch exist
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("File does not exist: %s\n", filePath)
	}

	//testing----------------------------
	fmt.Printf("Command-line arguments: %v\n", os.Args)

	files, err := os.ReadDir("data")
	if err != nil {
		log.Fatalf("Failed to read the data directory: %v", err)
	}
	fmt.Println("Files in data directory:")
	for _, file := range files {
		fmt.Println(file.Name())
	}

	//-----------------------------------------
	//still testing I just hate myself

	// Open the Excel file
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open the Excel file: %v", err)
	}

	// Get all the rows in the first sheet
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatalf("Failed to get rows: %v", err)
	}

	// Initialize a map to hold the data
	characters := make(map[int]Character)

	// Iterate through the rows and fill out the struct
	for i, row := range rows {
		// Skip the header row
		if i == 0 {
			continue
		}

		// Convert the ID and other int fields from string to int, defaulting to 0 if empty
		id := convertToInt(row[0], 0)
		level := convertToInt(row[2], 0)
		hp := convertToInt(row[3], 0)
		mp := convertToInt(row[4], 0)
		str := convertToInt(row[5], 0)
		end := convertToInt(row[6], 0)
		intel := convertToInt(row[7], 0)
		cha := convertToInt(row[8], 0)
		dex := convertToInt(row[9], 0)
		wis := convertToInt(row[10], 0)

		// Fill the struct with data
		character := Character{
			ID:    id,
			Name:  row[1],
			Level: level,
			HP:    hp,
			MP:    mp,
			Str:   str,
			End:   end,
			Int:   intel,
			Cha:   cha,
			Dex:   dex,
			Wis:   wis,
		}

		// Save the struct to the map
		characters[character.ID] = character
	}

	// Print the map to verify the data
	for id, character := range characters {
		fmt.Printf("ID: %d, Name: %s, Level: %d, HP: %d, MP: %d, Str: %d, End: %d, Int: %d, Cha: %d, Dex: %d, Wis: %d\n", id, character.Name, character.Level, character.HP, character.MP, character.Str, character.End, character.Int, character.Cha, character.Dex, character.Wis)
	}

	fmt.Println(altered_char(characters))

}

func convertToInt(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return value
}

// altered_char returns a randomly selected character and adds random numbers to the values by 0.2
func altered_char(characters map[int]Character) Character {
	keys := make([]int, 0, len(characters))
	for k := range characters {
		keys = append(keys, k)
	}
	randomIndex := rand.Intn(len(keys)) // Use the length of the keys slice as the upper limit
	selectedCharacter := characters[keys[randomIndex]]

	// Add random numbers to the values by 0.2
	selectedCharacter.Level += rand.Intn(3) // Random number between 0 and 2
	selectedCharacter.HP += rand.Intn(3)    // Random number between 0 and 2
	selectedCharacter.MP += rand.Intn(3)    // Random number between 0 and 2
	selectedCharacter.Str += rand.Intn(3)   // Random number between 0 and 2
	selectedCharacter.End += rand.Intn(3)   // Random number between 0 and 2
	selectedCharacter.Int += rand.Intn(3)   // Random number between 0 and 2
	selectedCharacter.Cha += rand.Intn(3)   // Random number between 0 and 2
	selectedCharacter.Dex += rand.Intn(3)   // Random number between 0 and 2
	selectedCharacter.Wis += rand.Intn(3)   // Random number between 0 and 2

	return selectedCharacter
}
