package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Monster struct {
	ID  int    `json:"ID"`
	HP  int    `json:"HP"`
	ATK int    `json:"ATK"`
	DEF int    `json:"DEF"`
	AGI int    `json:"AGI"`
	EVA int    `json:"EVA"`
	WPN string `json:"WPN"`
}

func main() {

	scriptPath := "../MonsterEngine/app.py"

	// Command to run the Python script
	cmd := exec.Command("python3", scriptPath)
	// Get the output of the Python script
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the raw output for debugging
	fmt.Println("Raw output:", string(output))

	// Parse the JSON output into a Monster struct
	var monster Monster
	err = json.Unmarshal(output, &monster)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print the monster struct
	fmt.Printf("Randomly selected monster: %+v\n", monster)
}
