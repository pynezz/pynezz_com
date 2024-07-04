package parser

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Read from char 0-2 (inclusive)
// If it contains "---", read until the next "---"
// If it doesn't, return "" (no metadata)
func readMetadata(md []byte) ([]byte, error) {
	// Read from char 0-2 (inclusive)
	for i := 0; i < 2; i++ {
		if md[i] != '-' {
			return nil, errors.New("no metadata found")
		}
	}

	delimiters := 0
	// read until the next "---"
	for i := 3; i < len(md); i++ {
		if md[i] == '-' {
			delimiters++
		}

		if delimiters == 3 {
			return md[3:i], nil
		}
	}

	// If it doesn't exist, return ""
	return nil, errors.New("error: error reading metadata, no closing delimiter found\nFormat:\n\t---\n\tmetadata\n\t---\n\tcontent")
}

// Parse the metadata into a Metadata struct for later use
func ParseMetadata(md []byte) (Metadata, error) {
	m := Metadata{}

	// read until the next "\n"
	for i := 0; i < len(md); i++ {
		if md[i] == '\n' {
			m.Title = string(md[:i])
			break
		}
	}

	lines := strings.Split(string(md), "\n")
	data := make(map[string]string) // key-value pairs

	for _, line := range lines {
		if line == "" {
			continue
		}
		data[strings.Split(line, ":")[0]] = strings.Split(line, ":")[1]
	}

	// Set the metadata fields
	m.Description = data["description"]
	m.Date, _ = time.Parse("2006-01-02", data["date"])
	m.LastModified, _ = time.Parse("2006-01-02", data["last_modified"])
	m.Tags = strings.Split(data["tags"], ",")

	fmt.Println("parsed data:\n", m)

	return m, nil
}
