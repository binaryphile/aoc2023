package aoc2023

import (
	"github.com/mariomac/gostream/item"
	"strconv"
	"strings"
)

type Slice[T any] []T

func (s Slice[T]) Filter(p func(T) bool) Slice[T] {
	result := make(Slice[T], 0, len(s))

	for _, elem := range s {
		if p(elem) {
			result = append(result, elem)
		}
	}

	return result
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
	result := make(Slice[T], 0, len(s.Slice))

	for _, elem := range s.Slice {
		result = append(result, f(elem))
	}

	return result
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

	first := 'a'

	var last rune

	for i := 0; i < len(runes); i++ {
		w := word[i:]

		if runes[i] >= '0' && runes[i] <= '9' {
			if first == 'a' {
				first = runes[i]
			}

			last = runes[i]
		} else if strings.HasPrefix(w, "one") {
			if first == 'a' {
				first = '1'
			}

			last = '1'
		} else if strings.HasPrefix(w, "two") {
			if first == 'a' {
				first = '2'
			}

			last = '2'
		} else if strings.HasPrefix(w, "three") {
			if first == 'a' {
				first = '3'
			}

			last = '3'
		} else if strings.HasPrefix(w, "four") {
			if first == 'a' {
				first = '4'
			}

			last = '4'
		} else if strings.HasPrefix(w, "five") {
			if first == 'a' {
				first = '5'
			}

			last = '5'
		} else if strings.HasPrefix(w, "six") {
			if first == 'a' {
				first = '6'
			}

			last = '6'
		} else if strings.HasPrefix(w, "seven") {
			if first == 'a' {
				first = '7'
			}

			last = '7'
		} else if strings.HasPrefix(w, "eight") {
			if first == 'a' {
				first = '8'
			}

			last = '8'
		} else if strings.HasPrefix(w, "nine") {
			if first == 'a' {
				first = '9'
			}

			last = '9'
		}
	}

	if first == 'a' {
		return 0
	}

	result := toInt(first)*10 + toInt(last)

	return result
}

func SumNumWords(words []string) int {
	sum, ok := StringSliceFrom[int](words).Map(numWordToNumber).Reduce(item.Add[int])
	if !ok {
		panic("no words")
	}

	return sum
}
