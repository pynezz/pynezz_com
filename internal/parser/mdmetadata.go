package parser

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/pynezz/pynezzentials/ansi"
)

// Read from char 0-2 (inclusive)
// If it contains "---", read until the next "---"
// If it doesn't, return "" (no metadata)
// func readMetadata(md []byte) ([]byte, error) {
// 	// Read from char 0-2 (inclusive)
// 	for i := 0; i < 2; i++ {
// 		if md[i] != '-' {
// 			return nil, errors.New("no metadata found")
// 		}
// 	}

// 	delimiters := 0
// 	// read until the next "---"
// 	for i := 3; i < len(md); i++ {
// 		if md[i] == '-' {
// 			delimiters++
// 		}

// 		if delimiters == 3 {
// 			return md[3:i], nil
// 		}
// 	}

// 	// If it doesn't exist, return ""
// 	return nil, errors.New("error: error reading metadata, no closing delimiter found\nFormat:\n\t---\n\tmetadata\n\t---\n\tcontent")
// }

// readMetadata extracts metadata from the markdown content.
func readMetadata(md []byte) ([]byte, error) {
	if len(md) < 3 || string(md[:3]) != "---" {
		return nil, errors.New("no metadata found")
	}

	end := strings.Index(string(md[3:]), "---")
	if end == -1 {
		return nil, errors.New("error: error reading metadata, no closing delimiter found\nFormat:\n\t---\n\tmetadata\n\t---\n\tcontent")
	}

	return md[3 : end+3], nil
}

// ParseMetadata parses the metadata bytes into a Metadata struct.
func ParseMetadata(md []byte) (Metadata, error) {
	m := Metadata{}
	lines := strings.Split(string(md), "\n")
	data := make(map[string]string)

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			data[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	for k, v := range data {
		ansi.PrintColor(ansi.DarkGreen, k+":"+v)
	}

	m.Title = data["Title"]
	m.Description = data["Description"]
	m.Date, _ = time.Parse("2006-01-02", data["Date"])
	m.LastModified, _ = time.Parse("2006-01-02", data["last_modified"])
	m.Tags = strings.Split(data["tags"], ",")

	fmt.Println("parsed data:\n", m)
	return m, nil
}

// // Parse the metadata into a Metadata struct for later use
// func ParseMetadata(md []byte) (Metadata, error) {
// 	m := Metadata{}

// 	// read until the next "\n"
// 	for i := 0; i < len(md); i++ {
// 		if md[i] == '\n' {
// 			m.Title = string(md[:i])
// 			break
// 		}
// 	}

// 	lines := strings.Split(string(md), "\n")
// 	data := make(map[string]string) // key-value pairs

// 	for _, line := range lines {
// 		if line == "" {
// 			continue
// 		}
// 		data[strings.Split(line, ":")[0]] = strings.Split(line, ":")[1]
// 	}

// 	// Set the metadata fields
// 	m.Description = data["description"]
// 	m.Date, _ = time.Parse("2006-01-02", data["date"])
// 	m.LastModified, _ = time.Parse("2006-01-02", data["last_modified"])
// 	m.Tags = strings.Split(data["tags"], ",")

// 	fmt.Println("parsed data:\n", m)

// 	return m, nil
// }
