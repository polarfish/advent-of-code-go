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

# Create output directory
OUTPUT_DIR="year$YEAR/day$DAY_PADDED"
mkdir -p $OUTPUT_DIR

INPUT_PATH="$OUTPUT_DIR/day${DAY_PADDED}.txt"

# Check for the existing input file
if [ -e "$INPUT_PATH" ]; then
  echo "Skip creating $INPUT_PATH (file exists)"
else
  # Download the input
  curl -s -H "Cookie: session=$SESSION" "https://adventofcode.com/$YEAR/day/$DAY/input" -o "$OUTPUT_DIR/day${DAY_PADDED}.txt"
  if [[ $? -eq 0 ]]; then
    echo "Created $OUTPUT_DIR/day${DAY_PADDED}.txt"
  else
    echo "Failed to create $OUTPUT_DIR/day${DAY_PADDED}.txt"
    exit 1
  fi
fi
