package main

import (
	"math/rand/v2"
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
	tables := []struct {
		name string
		size int
		data []int
	}{
		{
			name: "1k",
			size: 1_000,
		},
		{
			name: "1m",
			size: 1_000_000,
		},
	}
	for _, tt := range tables {
		tt.data = make([]int, tt.size)
		for i := range tt.data {
			tt.data[i] = rand.Int()
		}
		b.Run(tt.name, func(b *testing.B) {
			for b.Loop() {
				sort.Ints(tt.data)
			}
		})
	}
}

func BenchmarkSortFloat64s(b *testing.B) {
	tables := []struct {
		name string
		size int
		data []float64
	}{
		{
			name: "1k",
			size: 1_000,
		},
		{
			name: "1m",
			size: 1_000_000,
		},
	}
	for _, tt := range tables {
		tt.data = make([]float64, tt.size)
		for i := range tt.data {
			tt.data[i] = rand.Float64()
		}
		b.Run(tt.name, func(b *testing.B) {
			for b.Loop() {
				sort.Float64s(tt.data)
			}
		})
	}
}

// Ints:
//
//	1K 797 ns/op
//	1M 804_345 ns/op
// Floats:
//	1K 1041 ns/op
//  1M 1_101_813 ns/op
