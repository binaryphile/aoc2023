package aoc2023

import (
	"bufio"
	"os"
	"testing"
)

func TestSumWords(t *testing.T) {
	t.Parallel()

	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "eleven",
			args: args{
				words: []string{"a1c1w"},
			},
			want: 11,
		},
		{
			name: "single number eleven",
			args: args{
				words: []string{"a1w"},
			},
			want: 11,
		},
		{
			name: "actual",
			args: args{
				words: loadFile("day1_input.txt"),
			},
			want: 55108,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := SumWords(tt.args.words); got != tt.want {
				t.Errorf("SumWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func loadFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func TestSumNumWords(t *testing.T) {
	t.Parallel()

	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "actual",
			args: args{
				words: loadFile("day1_input.txt"),
			},
			want: 56324,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := SumNumWords(tt.args.words); got != tt.want {
				t.Errorf("SumNumWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
