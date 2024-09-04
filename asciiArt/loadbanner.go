package asciiArt

import (
	"bufio"
	"fmt"
	"os"
)

// CheckFileSize checks if the banner file size matches the expected size.
func CheckFileSize(fileName string, expectedSize int64) error {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return fmt.Errorf("error stating file: %v", err)
	}
	if fileInfo.Size() != expectedSize {
		return fmt.Errorf("file size does not match expected size: got %d, expected %d", fileInfo.Size(), expectedSize)
	}
	return nil
}

// LoadBannerMap loads the banner map from a file.
func LoadBannerMap(fileName string) (map[int][]string, error) {
	// Check file size before loading
	expectedSize := int64(6623) // Actual expected file size in bytes
	if err := CheckFileSize(fileName, expectedSize); err != nil {
		return nil, fmt.Errorf("file corruption detected: %v", err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	bannerMap := make(map[int][]string)
	key := 32
	lineCount := 0
	chunk := []string{}

	for scanner.Scan() {
		lines := scanner.Text()

		if lines != "" {
			chunk = append(chunk, lines)
			lineCount++
		}

		if lineCount == 8 {
			bannerMap[key] = chunk
			key++
			chunk = []string{}
			lineCount = 0
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return bannerMap, nil
}
