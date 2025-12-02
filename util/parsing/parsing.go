package parsing

import (
	"aoc25/util/xslices"
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	data string
}

type LineSlice []string
type WordSlice []string
type StringGrid [][]string

func FromFile(file string) *Parser {
	b, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return &Parser{strings.TrimSpace((string(b)))}
}

func FromString(s string) *Parser {
	return &Parser{strings.TrimSpace(s)}
}

func (p *Parser) Lines() LineSlice {
	return strings.Split(p.data, "\n")
}

func (p *Parser) Words() WordSlice {
	return strings.Fields(strings.ReplaceAll(p.data, "\n", " "))
}

func (p *Parser) Grid(sep string) [][]string {
	lines := p.Lines()
	grid := make([][]string, len(lines))
	for i, l := range lines {
		grid[i] = strings.Split(l, sep)
	}
	return grid
}

func (slice WordSlice) AsInts() []int {
	return xslices.Map(slice, Stoi)
}

func (slice LineSlice) AsInts() [][]int {
	result := make([][]int, len(slice))
	for i, line := range slice {
		words := strings.Fields(line)
		result[i] = xslices.Map(words, Stoi)
	}
	return result
}

func Stoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
