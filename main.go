package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/asciiArt"
)

func main() {
	fileName, err := asciiArt.GetBannerFile()
	if err != nil {
		fmt.Println(err)
		PrintUsage()
		return
	}

	// validate accepted number of os arguments
	if len(os.Args[1:]) < 1 || len(os.Args[1:]) > 4 || (!strings.HasPrefix(os.Args[1], "--color=") && len(os.Args[1:]) != 1) || (strings.HasPrefix(os.Args[1], "--color=") && len(os.Args[1:]) == 1) {
		fmt.Printf("Usage: go run . [OPTION] [STRING]\n\n")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}
	validBannerFiles := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}

	var lastArg string

	// Check if there are any arguments provided at the 4th index, else allocate it to standard
	if len(os.Args) == 4 {
		lastArg = "standard" // Default to "standard"
	} else {
		lastArg = os.Args[len(os.Args)-1]
	}

	if !validBannerFiles[lastArg] {
		PrintUsage()
		return
	}
	// print a new line and exit in case argument is a new line character only
	if os.Args[1] == "\\n" {
		fmt.Println()
		return
	}

	color := ""
	substring := ""
	text := ""

	// Check for the --color flag and extract the color
	if strings.HasPrefix(os.Args[1], "--color=") {
		color = strings.TrimPrefix(os.Args[1], "--color=")
		if len(os.Args) > 3 {
			substring = os.Args[2] // Letters to be colored
			text = os.Args[3]      // The actual text
		} else {
			text = os.Args[2]
		}
	} else {
		// No color flag, just the text
		text = os.Args[1]
	}

	// Load the banner map from the file
	bannerMap, err := asciiArt.LoadBannerMap(fileName)
	if err != nil {
		fmt.Println("Error loading banner map:", err)
		return
	}

	// Process the provided text and apply colors where necessary
	args := strings.ReplaceAll(text, "\\n", "\n")
	args = strings.ReplaceAll(args, "\\t", "    ")
	lines := strings.Split(args, "\n")

	// Generate the ASCII art for each line and apply color
	for _, line := range lines {
		asciiArt.PrintLineBanner(line, substring, color, bannerMap)
	}
}

func PrintUsage() {
	fmt.Printf("Usage: go run . [OPTION] [STRING]\n\n")
	fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
}
