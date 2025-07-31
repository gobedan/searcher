package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortStrings(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		want []string
	}{
		{
			name: "Test #1",
			s:    []string{"abc", "bca", "zx", "c"},
			want: []string{"abc", "bca", "c", "zx"},
		},
		{
			name: "Test #2",
			s:    []string{"абв", "бва", "zx", ""},
			want: []string{"", "zx", "абв", "бва"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Strings(tt.s)
			if got := tt.s; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}

func TestSortInts(t *testing.T) {
	data := []int{3, 33, 22, 1}
	want := []int{1, 3, 22, 33}
	sort.Ints(data)
	if got := data; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v; want: %v", got, want)
	}
}

func BenchmarkSortInts(b *testing.B) {
	data := []int{22, 3432, 23415, 648971, 12123, -2331, 0, 999_999}
	for b.Loop() {
		sort.Ints(data)
	}
}

func BenchmarkSortFloat64s(b *testing.B) {
	data := []float64{22, 34.32, 234.15, 648971, 12123, -2331, 0, 999_999}
	for b.Loop() {
		sort.Float64s(data)
	}
}
