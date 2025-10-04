#!/bin/bash

function refresh_puzzle_loader() {
  local loader_file_path="puzzles/loader/loader.go"
  local puzzles_sub_packages; puzzles_sub_packages=$(find puzzles -mindepth 2 -maxdepth 2 | grep "year20" | sort)

  echo "package loader" > "${loader_file_path}"
  echo "" >> "${loader_file_path}"
  echo "import (" >> "${loader_file_path}"
  # shellcheck disable=SC2068
  for sub_package in ${puzzles_sub_packages[@]}; do
      echo "	_ \"github.com/polarfish/advent-of-code-go/${sub_package}\"" >> "${loader_file_path}"
  done
  echo ")" >> "${loader_file_path}"

  echo "Refreshed puzzles/loader.go"
}

function main() {
  local session=$AOC_SESSION
  local year=$1
  local day=$2
  local status_code

  # Validate the year and the day are provided
  if [[ -z "$day" || -z "$year" ]]; then
    echo "Usage: $0 <year> <day>"
    exit 1
  fi

  local puzzle_base_url="https://adventofcode.com/$year/day/$day"

  # Validate the year and the day are correct
  status_code=$(curl -o /dev/null -s -w "%{http_code}" "$puzzle_base_url")
  if [ "$status_code" -ne 200 ]; then
    echo "Puzzle not found (year $year day $day)"
    exit 1
  fi

  # Validate session token is valid
  status_code=$(curl -o /dev/null -s -H "Cookie: session=$session" -w "%{http_code}" "$puzzle_base_url/input")
  if [ "$status_code" -ne 200 ]; then
    echo "Session token is missing or invalid (make sure to set AOC_SESSION environment variable with a valid token)"
    exit 1
  fi

  echo "Preparing year $year day $day"

  local day_padded; day_padded=$(printf "%02d" "$day")
  local base_name_lowercase="year${year}day${day_padded}"
  local base_name_pascalcase="Year${year}Day${day_padded}"

  # Create output directory
  local output_dir="puzzles/${year}/${base_name_lowercase}"
  mkdir -p "${output_dir}"

  # Create input file
  local input_file_path="${output_dir}/${base_name_lowercase}.txt"
  # Check for the existing input file
  if [ -e "$input_file_path" ]; then
    echo "Skip creating $input_file_path (file exists)"
  else
    # Download the input file
    if curl -s -H "Cookie: session=$session" "https://adventofcode.com/$year/day/$day/input" -o "$input_file_path"
    then
      echo "Created $input_file_path"
    else
      echo "Failed to create $input_file_path"
      exit 1
    fi
  fi

  # Create solution file
  local solution_file_path="$output_dir/${base_name_lowercase}.go"
  # Check for the existing solution file
  if [ -e "$solution_file_path" ]; then
    echo "Skip creating $solution_file_path (file exists)"
  else
    local sed_extract='s/.*--- Day [0-9]\{1,2\}: \(.*\) ---.*/\1/p'
    local sed_html_unescape='s/&nbsp;/ /g; s/&amp;/\&/g; s/&lt;/\</g; s/&gt;/\>/g; s/&quot;/\"/g; s/&apos;/\'"'"'/g; s/&ldquo;/\"/g; s/&rdquo;/\"/g;'
    local puzzle_title; puzzle_title=$(curl -s "$puzzle_base_url" | sed -n "$sed_extract" | sed "$sed_html_unescape")

    # Creating the solution stub
    echo "package ${base_name_lowercase}

import (
    _ \"embed\"
    \"strconv\"

    \"github.com/polarfish/advent-of-code-go/puzzles/registry\"
)

//go:embed ${base_name_lowercase}.txt
var input string

func init() {
    // https://adventofcode.com/${year}/day/${day}
    registry.AddPuzzle(${year}, ${day}, \"${puzzle_title}\", input, part1, part2)
}

func part1(input string) string {
    return strconv.Itoa(0)
}

func part2(input string) string {
    return strconv.Itoa(0)
}" > "$solution_file_path"

    if [[ $? -eq 0 ]]; then
      echo "Created $solution_file_path"
    else
      echo "Failed to create $solution_file_path"
      exit 1
    fi
  fi

  # Create test file
  local test_file_path="$output_dir/${base_name_lowercase}_test.go"
  # Check for the existing tests file
  if [ -e "$test_file_path" ]; then
    echo "Skip creating $test_file_path (file exists)"
  else

    # Creating the test stub
    # shellcheck disable=SC2028
    echo "package ${base_name_lowercase}

import (
    \"testing\"

    \"github.com/polarfish/advent-of-code-go/puzzles/utils\"
)

func Test${base_name_pascalcase}Part1(t *testing.T) {
    utils.Test(t, \"0\", part1(input))
}

func Test${base_name_pascalcase}Part2(t *testing.T) {
    utils.Test(t, \"0\", part2(input))
}" > "$test_file_path"

    if [[ $? -eq 0 ]]; then
      echo "Created $test_file_path"
    else
      echo "Failed to create $test_file_path"
      exit 1
    fi
  fi

  refresh_puzzle_loader
}

main "$@"