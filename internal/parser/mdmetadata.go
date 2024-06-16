package parser

import (
	"errors"
)

// Read from char 0-2 (inclusive)
// If it contains "---", read until the next "---"
// If it doesn't, return "" (no metadata)
func readMetadata(md []byte) ([]byte, error) {
	// Read from char 0-2 (inclusive)
	for i := 0; i < 2; i++ {
		if md[i] != '-' {
			return nil, errors.New("No metadata found")
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
	return nil, errors.New("error: error reading metadata, no closing delimiter found.\nFormat:\n\t---\n\tmetadata\n\t---\n\tcontent\n")
}

// Parse the metadata into a Metadata struct for later use
func parseMetadata(md []byte) (Metadata, error) {
	m := Metadata{}

	// read until the next "\n"
	for i := 0; i < len(md); i++ {
		if md[i] == '\n' {
			m.Title = string(md[:i])
			break
		}
	}

	// read until the next ":"
	// read until the next "\n"
	// read until the next "---"

	return m, nil
}
