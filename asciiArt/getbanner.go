package asciiArt

import (
	"fmt"
	"os"
	"strings"
)

// GetBannerFile selects the correct banner file based on user input
func GetBannerFile() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("no banner type specified")
	}

	// Default banner file
	defaultFile := "standard.txt"
	// Check if the user provided an argument like "standard", "shadow", or "thinkertoy"
	arg := strings.ToLower(os.Args[len(os.Args)-1]) // Last argument should be the banner type or text

	switch arg {
	case "standard":
		return "banners/standard.txt", nil
	case "shadow":
		return "banners/shadow.txt", nil
	case "thinkertoy":
		return "banners/thinkertoy.txt", nil
	default:
		// If no specific banner type is found, return the default
		return "banners/" + defaultFile, nil
	}
}

func PrintUsage() {
	fmt.Printf("Usage: go run . [OPTION] [STRING]\n\n")
	fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
}
