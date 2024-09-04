package asciiArt

import (
	"fmt"

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
			fmt.Printf("Invalid color: %s\n", err)
			colorCode = resetCode // Reset to default if color is invalid
		}
	}

	// Create a map to identify which characters should be colored
	colorMap := make(map[rune]bool)
	for _, char := range substring {
		colorMap[char] = true
	}

	// Process each character in the line
	for _, char := range line {
		banner, exists := bannerMap[int(char)]
		if !exists {
			fmt.Printf("Character %c not found in banner map\n", char)
			continue
		}

		for i := 0; i < 8; i++ {
			if substring == "" {
				// Color the entire text
				output[i] += colorCode + banner[i] + resetCode
			} else if colorMap[char] {
				// Color only the specified substring
				output[i] += colorCode + banner[i] + resetCode
			} else {
				// Regular output for other characters
				output[i] += banner[i]
			}
		}
	}

	// Print the final output
	for _, outLine := range output {
		fmt.Println(outLine)
	}
}
