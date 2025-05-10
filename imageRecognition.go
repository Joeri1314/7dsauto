package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func imageRec(match string) bool {
	cmd := exec.Command("python", "match.py", "screen.png", "resources/" + match + ".png")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	result := strings.TrimSpace(string(output))
	if result == "FOUND" {
		fmt.Println("FOUND: " + match)
		return true
	} else {
		fmt.Println("NOT FOUND: " + match)
		return false
	}
}