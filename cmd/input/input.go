package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		fmt.Println("Session token is missing or invalid (make sure to set AOC_SESSION environment variable with a valid token)")
		os.Exit(1)
	}

	inputFilesCount := 0
	root := "solutions"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".go") && strings.Contains(path, "day") {
			base := filepath.Base(path)
			if strings.HasSuffix(base, ".go") && !strings.HasSuffix(base, "_test.go") {
				inputFile := strings.TrimSuffix(path, ".go") + ".txt"
				if _, err := os.Stat(inputFile); os.IsNotExist(err) {
					year, day, ok := extractYearDayFromPath(path)
					if !ok {
						return nil
					}
					fmt.Println("Downloading", inputFile)
					err := downloadInput(session, year, day, inputFile)
					if err != nil {
						fmt.Printf("  Error: %v\n", err)
					} else {
						inputFilesCount++
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking solutions directory:", err)
		os.Exit(1)
	}
	fmt.Printf("Total input files downloaded: %d\n", inputFilesCount)
}

func extractYearDayFromPath(path string) (string, string, bool) {
	// path: solutions/2024/year2024day01/year2024day01.go
	parts := strings.Split(path, string(filepath.Separator))
	if len(parts) < 4 {
		return "", "", false
	}
	year := parts[1]
	base := filepath.Base(path)
	base = strings.TrimSuffix(base, ".go")
	if len(base) != len("year2024day01") {
		return "", "", false
	}
	day := base[len(base)-2:]
	return year, day, true
}

func downloadInput(session, year, day, outputPath string) error {
	if len(day) == 2 && day[0] == '0' {
		day = day[1:]
	}
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Cookie", "session="+session)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to download input: status %d", resp.StatusCode)
	}
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	_, err = io.Copy(f, resp.Body)
	return err
}
