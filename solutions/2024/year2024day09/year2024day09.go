package year2024day09

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
)

//go:embed year2024day09.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/9
	registry.AddSolution(2024, 9, "Disk Fragmenter", input, part1, part2)
}

func part1(input string) (string, error) {
	diskMap := createDiskMap(input)
	i := 0
	j := len(diskMap) - 1
	jFile := diskMap[j]
	var memoryAddress int
	var checksum int64

	for {
		if i%2 == 0 {
			iFile := diskMap[i]
			for iFile > 0 {
				checksum += int64(memoryAddress) * int64(i/2)
				memoryAddress++
				iFile--
			}
		} else {
			iSpaces := diskMap[i]
			for iSpaces > 0 {
				checksum += int64(memoryAddress) * int64(j/2)
				memoryAddress++
				iSpaces--
				jFile--
				if jFile == 0 {
					j--
					if j == i {
						break
					}
					j--
					jFile = diskMap[j]
				}
			}

			if j == i {
				break
			}
		}

		i++
		if i >= j {
			break
		}
	}
	for jFile > 0 {
		checksum += int64(memoryAddress) * int64(j/2)
		memoryAddress++
		jFile--
	}
	return strconv.FormatInt(checksum, 10), nil
}

type fileBlocks struct {
	id    int
	size  int
	moved bool
}

type space struct {
	freeSize int
	files    []*fileBlocks
}

func (s *space) addFile(f *fileBlocks) {
	s.files = append(s.files, f)
	s.freeSize -= f.size
}

func part2(input string) (string, error) {
	diskMap := createDiskMap(input)

	spaces := make([]*space, 0, len(diskMap)/2)
	files := make([]*fileBlocks, 0, len(diskMap)/2)
	for i := 0; i < len(diskMap); i++ {
		if i%2 == 1 {
			spaces = append(spaces, &space{freeSize: diskMap[i]})
		} else {
			files = append(files, &fileBlocks{id: i / 2, size: diskMap[i]})
		}
	}
	firstSpace := 0
	for i := 0; i < len(files); i++ {
		file := files[len(files)-i-1]
		for j := firstSpace; j < len(spaces)-i; j++ {
			sp := spaces[j]
			if sp.freeSize >= file.size {
				sp.addFile(file)
				file.moved = true
				break
			}
			if j == firstSpace && sp.freeSize == 0 {
				firstSpace++
			}
		}
	}
	memoryAddress := 0
	var checksum int64 = 0
	for i := 0; i < len(diskMap); i++ {
		if i%2 == 0 {
			file := files[i/2]
			if file.moved {
				memoryAddress += file.size
				continue
			}
			for j := 0; j < file.size; j++ {
				checksum += int64(memoryAddress) * int64(file.id)
				memoryAddress++
			}
		} else {
			sp := spaces[(i-1)/2]
			for _, file := range sp.files {
				for j := 0; j < file.size; j++ {
					checksum += int64(memoryAddress) * int64(file.id)
					memoryAddress++
				}
			}
			memoryAddress += sp.freeSize
		}
	}
	return strconv.FormatInt(checksum, 10), nil
}

func createDiskMap(input string) []int {
	input = strings.TrimSpace(input)
	diskMap := make([]int, len(input))
	for i, r := range input {
		diskMap[i] = int(r - 48)
	}
	return diskMap
}
