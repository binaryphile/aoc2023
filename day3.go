package aoc2023

import (
	"github.com/mariomac/gostream/item"
	"regexp"
	"strconv"
)

var digitRegex = regexp.MustCompile(`\d+`)

type Part struct {
	isActive bool
	id       string
}

type PartOrSymbol struct {
	Part
	isPart bool
}

func toParts(field [][]*PartOrSymbol) func(int, string) []int {
	return func(i int, line string) []int {
		length := len(line)

		result := make([]int, 0, length)

		for j := 0; j < length; j++ {
			if line[j] == '.' {
				continue
			}

			if line[j] >= '0' && line[j] <= '9' {
				digits := digitRegex.FindString(line[j:])

				part := &PartOrSymbol{
					Part: Part{
						id: digits,
					},
					isPart: true,
				}

				part.isActive = isPartActiveAt(i, j, len(digits), field)

				if part.isActive {
					id, err := strconv.Atoi(digits)
					if err != nil {
						panic(err)
					}

					result = append(result, id)
				}

				field[i][j] = part

				j += len(digits) - 1

				field[i][j] = part
			} else {
				symbol := &PartOrSymbol{
					isPart: false,
				}

				result = append(result, symbolMakesNewActiveParts(i, j, field)...)

				field[i][j] = symbol
			}
		}

		return result
	}
}

func SumParts(schematics []string) int {
	field := make([][]*PartOrSymbol, len(schematics))

	for i, schematic := range schematics {
		field[i] = make([]*PartOrSymbol, len(schematic))
	}

	sum, ok := FlatMapWithIndex(schematics, toParts(field)).Reduce(item.Add[int])
	if !ok {
		panic("no parts")
	}

	return sum
}

func isPartActiveAt(r, c, length int, field [][]*PartOrSymbol) bool {
	for i := r - 1; i <= r; i++ {
		if i < 0 {
			continue
		}

		row := field[i]

		for j := c - 1; j <= c+length; j++ {
			if j < 0 {
				continue
			}

			if j >= len(row) {
				break
			}

			symbol := row[j]

			if symbol != nil && !symbol.isPart {
				return true
			}
		}
	}

	return false
}

func symbolMakesNewActiveParts(r, c int, field [][]*PartOrSymbol) []int {
	result := make([]int, 0)

	for i := r - 1; i <= r; i++ {
		if i < 0 {
			continue
		}

		row := field[i]

		for j := c - 1; j <= c+1; j++ {
			if j < 0 {
				continue
			}

			if j >= len(row) {
				break
			}

			part := row[j]

			if part != nil && part.isPart && !part.isActive {
				part.isActive = true

				id, err := strconv.Atoi(part.id)
				if err != nil {
					panic(err)
				}

				result = append(result, id)
			}
		}
	}

	return result
}
