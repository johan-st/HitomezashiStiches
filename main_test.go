package main

import (
	"reflect"
	"testing"
)

func Test_randSeed(t *testing.T) {
	mathRandSeed := int64(0)
	type args struct {
		randSeed int64
		size     int
		oddsTrue float32
	}
	tests := []struct {
		name string
		args args
		want seed
	}{
		{"all false", args{mathRandSeed, 5, 0}, []bool{false, false, false, false, false}},
		{"all true", args{mathRandSeed, 5, 1}, []bool{true, true, true, true, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randSeed(tt.args.randSeed, tt.args.size, tt.args.oddsTrue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_normSeed(t *testing.T) {
	type args struct {
		s      seed
		length int
	}
	tests := []struct {
		name string
		args args
		want seed
	}{
		{"seed with correct length is preserved", args{seed{true, true, false, true, false}, 5}, seed{true, true, false, true, false}},
		{"seed too long are cut", args{seed{true, true, false, true, false}, 2}, seed{true, true}},
		{"seed too short is repeated", args{seed{true, true, false, true, false}, 13}, seed{true, true, false, true, false, true, true, false, true, false, true, true, false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normSeed(tt.args.s, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
