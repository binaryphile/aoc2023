package aoc2023

import (
	"github.com/mariomac/gostream/item"
	"strconv"
	"strings"
)

type Slice[T any] []T

func (s Slice[T]) Filter(p func(T) bool) Slice[T] {
	return Filter(s, p)
}

func (s Slice[T]) Map(f func(T) T) Slice[T] {
	return Map(s, f)
}

func (s Slice[T]) Reduce(f func(T, T) T) (_ T, ok bool) {
	if len(s) == 0 {
		return
	}

	result := s[0]

	for _, elem := range s[1:] {
		result = f(result, elem)
	}

	return result, true
}

type RuneSlice[R any] struct {
	Slice[rune]
}

func (s RuneSlice[R]) Filter(p func(rune) bool) RuneSlice[R] {
	return RuneSlice[R]{
		Slice: s.Slice.Filter(p),
	}
}

func (s RuneSlice[R]) Map(f func(rune) R) Slice[R] {
	return Map(s.Slice, f)
}

func RuneSliceFromString[R any](s string) RuneSlice[R] {
	return RuneSlice[R]{
		Slice: []rune(s),
	}
}

func Filter[T any](s []T, p func(T) bool) Slice[T] {
	result := make([]T, 0, len(s))

	for _, elem := range s {
		if p(elem) {
			result = append(result, elem)
		}
	}

	return result
}

func Map[T, R any](s []T, f func(T) R) Slice[R] {
	result := make(Slice[R], 0, len(s))

	for _, elem := range s {
		result = append(result, f(elem))
	}

	return result
}

type StringSlice[T any] struct {
	Slice[string]
}

func StringSliceFrom[T any](s []string) StringSlice[T] {
	return StringSlice[T]{
		Slice: s,
	}
}

func (s StringSlice[T]) Map(f func(string) T) Slice[T] {
	return Map(s.Slice, f)
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func toInt(r rune) int {
	i, err := strconv.Atoi(string(r))
	if err != nil {
		panic(err)
	}

	return i
}

func wordToNumber(word string) int {
	digits := RuneSliceFromString[int](word).Filter(isDigit).Map(toInt)

	if len(digits) == 0 {
		return 0
	}

	return digits[0]*10 + digits[len(digits)-1]
}

func SumWords(words []string) int {
	sum, ok := StringSliceFrom[int](words).Map(wordToNumber).Reduce(item.Add[int])
	if !ok {
		panic("no words")
	}

	return sum
}

func numWordToNumber(word string) int {
	runes := []rune(word)

	first := -1

	var last int

	for i := 0; i < len(runes); i++ {
		w := word[i:]

		if runes[i] >= '0' && runes[i] <= '9' {
			last = toInt(runes[i])

			if first == -1 {
				first = 10 * last
			}
		} else if strings.HasPrefix(w, "one") {
			last = 1

			if first == -1 {
				first = 10
			}
		} else if strings.HasPrefix(w, "two") {
			last = 2

			if first == -1 {
				first = 20
			}
		} else if strings.HasPrefix(w, "three") {
			last = 3

			if first == -1 {
				first = 30
			}
		} else if strings.HasPrefix(w, "four") {
			last = 4

			if first == -1 {
				first = 40
			}
		} else if strings.HasPrefix(w, "five") {
			last = 5

			if first == -1 {
				first = 50
			}
		} else if strings.HasPrefix(w, "six") {
			last = 6

			if first == -1 {
				first = 60
			}
		} else if strings.HasPrefix(w, "seven") {
			last = 7

			if first == -1 {
				first = 70
			}
		} else if strings.HasPrefix(w, "eight") {
			last = 8

			if first == -1 {
				first = 80
			}
		} else if strings.HasPrefix(w, "nine") {
			last = 9

			if first == -1 {
				first = 90
			}
		}
	}

	if first == -1 {
		return 0
	}

	result := first + last

	return result
}

func SumNumWords(words []string) int {
	sum, ok := StringSliceFrom[int](words).Map(numWordToNumber).Reduce(item.Add[int])
	if !ok {
		panic("no words")
	}

	return sum
}
