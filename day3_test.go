package aoc2023

import (
	"testing"
)

func TestSumParts(t *testing.T) {
	t.Parallel()

	type args struct {
		schematics []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "actual",
			args: args{
				schematics: loadFile("day3_input.txt"),
			},
			want: 2207,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := SumParts(tt.args.schematics); got != tt.want {
				t.Errorf("SumParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
