package parser

import (
	"errors"
	"fmt"
	"strconv"
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

	m.Title = data["Title"]
	m.Description = data["Description"]
	if d := data["Date"]; d != "" { // if date is provided, parse it
		m.Date = parseDate(d)
	} else { // if not, default to today's date
		m.Date = parseDate(time.Now().Format("02.01.2006")) //passing to parseDate to ensure time.Time struct is created correctly)
	}
	m.LastModified, _ = time.Parse("02.01.2006", data["last_modified"])
	if m.LastModified.IsZero() {
		m.LastModified = time.Now()
	}

	m.Tags = strings.Split(strings.Trim(data["Tags"], "[]"), ",")

	fmt.Println("parsed data:\n", m)
	return m, nil
}

// parseDate parses a date string into a time.Time struct.
// The date string MUST be in the format "dd.mm.yyyy".
func parseDate(date string) time.Time {
	// If any parts of the date are missing or wrong, we'll fix it in the parseMoDaYr function
	d, m, y := parseDaMoYr(date)
	dmy := fmt.Sprintf("%02s.%02s.%04s", d, m, y)
	res, err := time.Parse("02.01.2006", dmy)
	if err != nil {
		ansi.PrintError("Error parsing date: " + err.Error())
		return time.Time{}
	}

	ansi.PrintDebug("parsed date: " + res.String())

	return res
}

func parseDaMoYr(date string) (string, string, string) {
	dArr := strings.Split(date, ".")
	if len(dArr) != 3 {
		ansi.PrintError("dd.mm.yyyy not provided - Invalid date: " + date +
			"\nPlease make sure the date is in the format \"dd.mm.yyyy\".")
		return time.Now().Format("02"), time.Now().Format("01"), time.Now().Format("2006")
	}
	m := dArr[1]
	if len(m) == 1 {
		m = "0" + m
	} // month
	if i, _ := strconv.Atoi(m); i > 12 {
		ansi.PrintError("Invalid month: " + m +
			"\nMonth can't be greater than 12.\nPlease make sure the date is in the format \"dd.mm.yyyy\".")
		m = time.Now().Format("01")
	}
	d := dArr[0]
	if len(d) == 1 {
		d = "0" + d
	} // day
	if i, _ := strconv.Atoi(d); i > 31 {
		ansi.PrintError("Invalid day: " + d +
			"\nDay can't be greater than 31.\nPlease make sure the date is in the format \"dd.mm.yyyy\".")
		d = time.Now().Format("02")
	}

	// A bit over the top, but why not
	if i, _ := strconv.Atoi(d); i > 30 && (m == "04" || m == "06" || m == "09" || m == "11") {
		ansi.PrintError("Invalid day: " + d +
			"\nDay can't be greater than 30 in this month.\nPlease make sure the date is in the format \"dd.mm.yyyy\".")
		d = time.Now().Format("02")
	}
	y := dArr[2]
	if len(y) == 2 {
		y = "20" + y // year (this will be a problem we'll have to deal with in year 2100)
	}

	if i, _ := strconv.Atoi(d); i > 29 && m == "02" {
		ansi.PrintError("Invalid day: " + d +
			"\nDay can't be greater than 29 in this month.\nPlease make sure the date is in the format \"dd.mm.yyyy\".")
		d = time.Now().Format("02")
	}
	if i, _ := strconv.Atoi(y); i > 2100 {
		ansi.PrintError("Invalid year: " + y +
			"\nYear can't be greater than 2100.\nPlease make sure the date is in the format \"dd.mm.yyyy\".")
		y = time.Now().Format("2006")
	}
	fmt.Println("parsed date: " + d + "." + m + "." + y)
	return d, m, y
}

// extractMetadata separates metadata from content
func extractMetadata(mdContent string) (string, string) {
	if strings.HasPrefix(mdContent, "---") {
		end := strings.Index(mdContent[3:], "---")
		if end != -1 {
			return mdContent[:end+6], mdContent[end+6:]
		}
	}
	return "", mdContent
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
