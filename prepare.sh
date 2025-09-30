#!/bin/bash

SESSION=$AOC_SESSION
YEAR=$1
DAY=$2

# Validate the year and the day are provided
if [[ -z "$DAY" || -z "$YEAR" ]]; then
  echo "Usage: $0 <year> <day>"
  exit 1
fi

BASE_URL="https://adventofcode.com/$YEAR/day/$DAY"

# Validate the year and the day are correct
STATUS_CODE=$(curl -o /dev/null -s -w "%{http_code}" "$BASE_URL")
if [ "$STATUS_CODE" -ne 200 ]; then
  echo "Puzzle not found (year $YEAR day $DAY)"
  exit 1
fi

# Validate session token is valid
STATUS_CODE=$(curl -o /dev/null -s -H "Cookie: session=$SESSION" -w "%{http_code}" "$BASE_URL/input")
if [ "$STATUS_CODE" -ne 200 ]; then
  echo "Session token is missing or invalid (make sure to set AOC_SESSION environment variable with a valid token)"
  exit 1
fi

DAY_PADDED=$(printf "%02d" "$DAY")

echo "Preparing year $YEAR day $DAY"

# Create output directory
OUTPUT_DIR="puzzles"
mkdir -p $OUTPUT_DIR

PUZZLE_BASE_LOWERCASE="year${YEAR}day${DAY_PADDED}"
PUZZLE_BASE_CAMELCASE="year${YEAR}Day${DAY_PADDED}"
PUZZLE_BASE_PASCALCASE="Year${YEAR}Day${DAY_PADDED}"

INPUT_PATH="${OUTPUT_DIR}/${PUZZLE_BASE_LOWERCASE}.txt"

# Check for the existing input file
if [ -e "$INPUT_PATH" ]; then
  echo "Skip creating $INPUT_PATH (file exists)"
else
  # Download the input
  if curl -s -H "Cookie: session=$SESSION" "https://adventofcode.com/$YEAR/day/$DAY/input" -o "$INPUT_PATH"
  then
    echo "Created $INPUT_PATH"
  else
    echo "Failed to create $INPUT_PATH"
    exit 1
  fi
fi

SOLUTION_PATH="$OUTPUT_DIR/${PUZZLE_BASE_LOWERCASE}.go"
# Check for the existing solution file
if [ -e "$SOLUTION_PATH" ]; then
  echo "Skip creating $SOLUTION_PATH (file exists)"
else
  SED_EXTRACT='s/.*--- Day [0-9]\{1,2\}: \(.*\) ---.*/\1/p'
  SED_HTML_UNESCAPE='s/&nbsp;/ /g; s/&amp;/\&/g; s/&lt;/\</g; s/&gt;/\>/g; s/&quot;/\"/g; s/&apos;/\'"'"'/g; s/&ldquo;/\"/g; s/&rdquo;/\"/g;'
  PUZZLE_TITLE=$(curl -s "$BASE_URL" | sed -n "$SED_EXTRACT" | sed "$SED_HTML_UNESCAPE")

  # Creating the solution stub
  echo "package ${OUTPUT_DIR}

import (
	_ \"embed\"
	\"strconv\"
)

//go:embed ${PUZZLE_BASE_LOWERCASE}.txt
var year${YEAR}Day${DAY_PADDED}Input string

func init() {
	// https://adventofcode.com/${YEAR}/day/${DAY}
	addPuzzle(${YEAR}, ${DAY}, \"${PUZZLE_TITLE}\", ${PUZZLE_BASE_CAMELCASE}Input, ${PUZZLE_BASE_CAMELCASE}Part1, ${PUZZLE_BASE_CAMELCASE}Part2)
}

func ${PUZZLE_BASE_CAMELCASE}Part1(input string) string {
	return strconv.Itoa(0)
}

func ${PUZZLE_BASE_CAMELCASE}Part2(input string) string {
	return strconv.Itoa(0)
}" > "$SOLUTION_PATH"

  if [[ $? -eq 0 ]]; then
    echo "Created $SOLUTION_PATH"
  else
    echo "Failed to create $SOLUTION_PATH"
    exit 1
  fi
fi

TESTS_PATH="$OUTPUT_DIR/${PUZZLE_BASE_LOWERCASE}_test.go"
# Check for the existing tests file
if [ -e "$TESTS_PATH" ]; then
  echo "Skip creating $TESTS_PATH (file exists)"
else

  # Creating the tests stub
  echo "package ${OUTPUT_DIR}

import \"testing\"

func Test${PUZZLE_BASE_PASCALCASE}Part1(t *testing.T) {
	runTests(t, ${PUZZLE_BASE_CAMELCASE}Part1, map[string]testCase{
		\"input\": {${PUZZLE_BASE_CAMELCASE}Input, \"0\"},
	})
}

func Test${PUZZLE_BASE_PASCALCASE}Part2(t *testing.T) {
	runTests(t, ${PUZZLE_BASE_CAMELCASE}Part2, map[string]testCase{
		\"input\": {${PUZZLE_BASE_CAMELCASE}Input, \"0\"},
	})
}" > "$TESTS_PATH"

  if [[ $? -eq 0 ]]; then
    echo "Created $TESTS_PATH"
  else
    echo "Failed to create $TESTS_PATH"
    exit 1
  fi
fi
