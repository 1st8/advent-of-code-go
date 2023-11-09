package cli

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// DownloadInput downloads the puzzle input for the given day from Advent of Code and saves it to a file.
// The session parameter should be the value of your session cookie from the Advent of Code website.
func DownloadInput(year int, day int, session string) error {
	// Create the URL
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	// Create a request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set the session cookie
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-200 status: %d %s", resp.StatusCode, resp.Status)
	}

	// Ensure the inputs directory exists
	yearDir := filepath.Join("inputs", fmt.Sprintf("%d", year))
	err = os.MkdirAll(yearDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Create the file
	filePath := filepath.Join(yearDir, fmt.Sprintf("%d.txt", day))
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the response body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Input for day %d has been saved to %s\n", day, filePath)
	return nil
}
