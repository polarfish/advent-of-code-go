package year2024day23

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day23.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/23
	registry.AddSolution(2024, 23, "LAN Party", input, part1, part2)
}

func part1(input string) (string, error) {
	g, err := parseInput(input)
	if err != nil {
		return "", err
	}

	triangles := map[string]struct{}{}
	for from, connsMap := range g {
		// collect neighbors as slice
		conns := make([]string, 0, len(connsMap))
		for k := range connsMap {
			conns = append(conns, k)
		}

		for i := 0; i < len(conns)-1; i++ {
			for j := i; j < len(conns); j++ {
				ci := conns[i]
				cj := conns[j]
				if ci == cj {
					continue
				}
				if g[ci] != nil && g[ci][cj] && g[cj] != nil && g[cj][ci] && (strings.HasPrefix(from, "t") || strings.HasPrefix(ci, "t") || strings.HasPrefix(cj, "t")) {
					triple := []string{from, ci, cj}
					sort.Strings(triple)
					key := strings.Join(triple, ",")
					triangles[key] = struct{}{}
				}
			}
		}
	}

	return strconv.Itoa(len(triangles)), nil
}

func part2(input string) (string, error) {
	g, err := parseInput(input)
	if err != nil {
		return "", err
	}

	maxPossibleGroup := 0
	keys := make([]string, len(g))
	i := 0
	for k, v := range g {
		if len(v) > maxPossibleGroup {
			maxPossibleGroup = len(v)
		}
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	checked := make(map[string]struct{}, len(keys))
	memo := make(map[string]struct{})
	var best []string

	for _, computer := range keys {
		if len(best) >= maxPossibleGroup {
			break
		}
		checked[computer] = struct{}{}
		res := dfs(g, []string{computer}, checked, memo)
		if len(res) > len(best) {
			best = append([]string{}, res...)
		}
	}
	return strings.Join(best, ","), nil
}

func parseInput(input string) (map[string]map[string]bool, error) {
	g := map[string]map[string]bool{}
	for _, line := range utils.Lines(input) {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return nil, utils.ErrBadInput
		}
		a, b := parts[0], parts[1]
		if g[a] == nil {
			g[a] = map[string]bool{}
		}
		if g[b] == nil {
			g[b] = map[string]bool{}
		}
		g[a][b] = true
		g[b][a] = true
	}
	return g, nil
}

func dfs(g map[string]map[string]bool, group []string, checked map[string]struct{}, memo map[string]struct{}) []string {
	if len(group) == 13 {
		return group
	}

	// build intersection of neighbors
	if g[group[0]] == nil {
		return group
	}
	connections := make(map[string]struct{})
	for k := range g[group[0]] {
		connections[k] = struct{}{}
	}
	for i := 1; i < len(group); i++ {
		connectionsI := g[group[i]]
		for k := range connections {
			if _, ok := connectionsI[k]; !ok {
				delete(connections, k)
			}
		}
	}

	var best []string
	for conn := range connections {
		if _, ok := checked[conn]; ok {
			continue
		}
		nextGroup := make([]string, len(group)+1)
		copy(nextGroup, group)
		nextGroup[len(group)] = conn
		sort.Strings(nextGroup)
		key := strings.Join(nextGroup, ",")
		if _, seen := memo[key]; seen {
			continue
		}
		memo[key] = struct{}{}
		res := dfs(g, nextGroup, checked, memo)
		if len(res) > len(best) {
			best = res
		}
	}

	if len(best) > len(group) {
		return best
	}
	return group
}
