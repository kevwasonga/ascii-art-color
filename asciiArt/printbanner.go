package asciiArt

import (
	"fmt"
	"strings"

	"ascii/utils"
)

// PrintLineBanner prints the banner for a line of text and applies color if specified.
func PrintLineBanner(line, substring, color string, bannerMap map[int][]string) {
	if line == "" {
		fmt.Println()
		return
	}

	output := make([]string, 8)
	colorCode := ""
	resetCode := "\033[0m"

	// Find the color ANSI code if color is provided
	if color != "" {
		var err error
		colorCode, err = utils.FindColorANSI(color)
		if err != nil {
			fmt.Printf("Invalid color, try another valid color: %s\n", err)
			return
			// colorCode = resetCode // Reset to default if color is invalid
		}
	}

	// Process the entire string if no substring is provided
	if substring == "" {
		for _, char := range line {
			banner, exists := bannerMap[int(char)]
			if !exists {
				fmt.Printf("Character %c not found in banner map\n", char)
				continue
			}
			for i := 0; i < 8; i++ {
				output[i] += colorCode + banner[i] + resetCode
			}
		}
	} else {
		// Process each character in the line
		for i := 0; i < len(line); i++ {
			char := line[i]
			banner, exists := bannerMap[int(char)]
			if !exists {
				fmt.Printf("Character %c not found in banner map\n", char)
				continue
			}

			// Check if the current position is part of the substring
			subStrIndex := strings.Index(line[i:], substring)
			if subStrIndex == 0 {
				// Color the substring
				for j := 0; j < len(substring); j++ {
					if i+j < len(line) {
						char := line[i+j]
						banner, exists := bannerMap[int(char)]
						if !exists {
							fmt.Printf("Character %c not found in banner map\n", char)
							continue
						}
						for k := 0; k < 8; k++ {
							output[k] += colorCode + banner[k] + resetCode
						}
					}
				}
				i += len(substring) - 1 // Skip past the substring
			} else {
				// Regular output for other characters
				for k := 0; k < 8; k++ {
					output[k] += banner[k]
				}
			}
		}
	}

	// Print the final output
	for _, outLine := range output {
		fmt.Println(outLine)
	}
}
