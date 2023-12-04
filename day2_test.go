package aoc2023

import (
	"testing"
)

func TestSumGames(t *testing.T) {
	t.Parallel()

	type args struct {
		games []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "actual",
			args: args{
				games: loadFile("day2_input.txt"),
			},
			want: 2207,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := SumGames(tt.args.games); got != tt.want {
				t.Errorf("SumGames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumPowerGames(t *testing.T) {
	t.Parallel()

	type args struct {
		games []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "actual",
			args: args{
				games: loadFile("day2_input.txt"),
			},
			want: 62241,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := SumPowerGames(tt.args.games); got != tt.want {
				t.Errorf("SumPowerGames() = %v, want %v", got, tt.want)
			}
		})
	}
}
