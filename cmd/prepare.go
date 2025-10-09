package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: prepare <year> <day>")
		os.Exit(1)
	}
	year := os.Args[1]
	day := os.Args[2]
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		fmt.Println("Session token is missing or invalid (make sure to set AOC_SESSION environment variable with a valid token)")
		os.Exit(1)
	}

	if err := prepareSolution(year, day, session); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func prepareSolution(year, day, session string) error {
	if !validatePuzzleExists(year, day) {
		return fmt.Errorf("Solution not found (year %s day %s)", year, day)
	}
	if !validateSession(session, year, day) {
		return errors.New("session token is missing or invalid (make sure to set AOC_SESSION environment variable with a valid token)")
	}
	fmt.Printf("Preparing year %s day %s\n", year, day)
	dayPadded := fmt.Sprintf("%02s", day)
	baseNameLower := fmt.Sprintf("year%vday%v", year, dayPadded)
	baseNamePascal := fmt.Sprintf("Year%vDay%v", year, dayPadded)
	outputDir := filepath.Join("solutions", year, baseNameLower)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}
	inputFile := filepath.Join(outputDir, baseNameLower+".txt")
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		if err := downloadInput(session, year, day, inputFile); err != nil {
			return err
		}
		fmt.Println("Created", inputFile)
	} else {
		fmt.Println("Skip creating", inputFile, "(file exists)")
	}
	solutionFile := filepath.Join(outputDir, baseNameLower+".go")
	if _, err := os.Stat(solutionFile); os.IsNotExist(err) {
		if err := createSolutionStub(solutionFile, year, day, baseNameLower); err != nil {
			return err
		}
		fmt.Println("Created", solutionFile)
	} else {
		fmt.Println("Skip creating", solutionFile, "(file exists)")
	}
	testFile := filepath.Join(outputDir, baseNameLower+"_test.go")
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		if err := createTestStub(testFile, baseNameLower, baseNamePascal); err != nil {
			return err
		}
		fmt.Println("Created", testFile)
	} else {
		fmt.Println("Skip creating", testFile, "(file exists)")
	}
	if err := regenerateSolutionLoader(); err != nil {
		return err
	}
	if err := regenerateSolutionReadme(); err != nil {
		return err
	}
	return nil
}

func validatePuzzleExists(year, day string) bool {
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s", year, day)
	resp, err := http.Head(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}

func validateSession(session, year, day string) bool {
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}
	req.Header.Set("Cookie", "session="+session)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}

func downloadInput(session, year, day, outputPath string) error {
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
	defer resp.Body.Close()
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}

func getPuzzleTitle(year, day string) string {
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s", year, day)
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	re := regexp.MustCompile(`--- Day [0-9]{1,2}: (.*?) ---`)
	matches := re.FindSubmatch(body)
	if len(matches) > 1 {
		return htmlUnescape(string(matches[1]))
	}
	return ""
}

func htmlUnescape(s string) string {
	replacer := strings.NewReplacer(
		"&nbsp;", " ", "&amp;", "&", "&lt;", "<", "&gt;", ">", "&quot;", "\"", "&apos;", "'", "&ldquo;", "\"", "&rdquo;", "\"",
	)
	return replacer.Replace(s)
}

func createSolutionStub(path, year, day, baseNameLower string) error {
	title := getPuzzleTitle(year, day)
	stub := fmt.Sprintf(`package %s

import (
    _ "embed"
    "strconv"

    "github.com/polarfish/advent-of-code-go/registry"
)

//go:embed %s.txt
var input string

func init() {
    // https://adventofcode.com/%s/day/%s
    registry.AddSolution(%s, %s, "%s", input, part1, part2)
}

func part1(input string) string {
    return strconv.Itoa(0)
}

func part2(input string) string {
    return strconv.Itoa(0)
}
`, baseNameLower, baseNameLower, year, day, year, day, title)
	return os.WriteFile(path, []byte(stub), 0644)
}

func createTestStub(path, baseNameLower, baseNamePascal string) error {
	stub := fmt.Sprintf(`package %s

import (
    "testing"

    "github.com/polarfish/advent-of-code-go/utils"
)

func Test%sPart1(t *testing.T) {
    utils.Test(t, "0", part1(input))
}

func Test%sPart2(t *testing.T) {
    utils.Test(t, "0", part2(input))
}
`, baseNameLower, baseNamePascal, baseNamePascal)
	return os.WriteFile(path, []byte(stub), 0644)
}

func regenerateSolutionLoader() error {
	loaderFilePath := filepath.Join("loader", "loader.go")
	var subPackages []string
	root := "solutions"
	yearDirs, err := os.ReadDir(root)
	if err != nil {
		return err
	}
	for _, yearDir := range yearDirs {
		if !yearDir.IsDir() || !strings.HasPrefix(yearDir.Name(), "20") {
			continue
		}
		yearPath := filepath.Join(root, yearDir.Name())
		dayDirs, err := os.ReadDir(yearPath)
		if err != nil {
			continue
		}
		for _, dayDir := range dayDirs {
			if !dayDir.IsDir() || !strings.HasPrefix(dayDir.Name(), "year") {
				continue
			}
			subPackages = append(subPackages, filepath.Join(yearDir.Name(), dayDir.Name()))
		}
	}
	sort.Strings(subPackages)
	var buf bytes.Buffer
	buf.WriteString("package loader\n\nimport (\n")
	for _, sub := range subPackages {
		buf.WriteString(fmt.Sprintf("\t_ \"github.com/polarfish/advent-of-code-go/solutions/%s\"\n", sub))
	}
	buf.WriteString(")\n")
	if err := os.WriteFile(loaderFilePath, buf.Bytes(), 0644); err != nil {
		return err
	}
	fmt.Println("Regenerated " + loaderFilePath)
	return nil
}

func regenerateSolutionReadme() error {
	outputFile := filepath.Join("solutions", "README.md")
	var buf bytes.Buffer
	buf.WriteString("# Solutions index\n\n")
	root := "solutions"
	yearDirs, err := os.ReadDir(root)
	if err != nil {
		return err
	}
	var years []string
	for _, yearDir := range yearDirs {
		if yearDir.IsDir() && strings.HasPrefix(yearDir.Name(), "20") {
			years = append(years, yearDir.Name())
		}
	}
	sort.Strings(years)
	for _, year := range years {
		buf.WriteString(fmt.Sprintf("## %s\n\n", year))
		yearPath := filepath.Join(root, year)
		dayDirs, err := os.ReadDir(yearPath)
		if err != nil {
			continue
		}
		var dayLinks []string
		for _, dayDir := range dayDirs {
			if !dayDir.IsDir() || !strings.HasPrefix(dayDir.Name(), "year") {
				continue
			}
			solutionFile := filepath.Join(root, year, dayDir.Name(), dayDir.Name()+".go")
			_, err := os.Stat(solutionFile)
			if err == nil {
				re := regexp.MustCompile(`year[0-9]{4}day([0-9]{2})`)
				matches := re.FindStringSubmatch(dayDir.Name())
				if len(matches) > 1 {
					dayNum := matches[1]
					link := fmt.Sprintf("[Day %s](%s)", dayNum, filepath.Join(year, dayDir.Name(), dayDir.Name()+".go"))
					dayLinks = append(dayLinks, link)
				}
			}
		}
		if len(dayLinks) > 0 {
			buf.WriteString(strings.Join(dayLinks, "\n • "))
			buf.WriteString("\n\n")
		}
	}
	if err := os.WriteFile(outputFile, buf.Bytes(), 0644); err != nil {
		return err
	}
	fmt.Println("Regenerated solutions/README.md")
	return nil
}
