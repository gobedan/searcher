package bsearch

import (
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/trace"
	"sort"
	"testing"
	"time"
)

func TestSearches(t *testing.T) {
	tests := []struct {
		name string
		data []int
		item int
		want int
	}{
		{
			name: "#1",
			data: []int{1, 2, 4, 6, 10},
			item: 6,
			want: 3,
		},
		{
			name: "#2",
			data: []int{1, 2, 4, 5, 6, 10},
			item: 6,
			want: 4,
		},
		{
			name: "#3",
			data: []int{1, 2, 4, 5, 6, 10},
			item: 16,
			want: -1,
		},
		{
			name: "#4",
			data: []int{1, 2, 4, 5, 6, 10},
			item: 2,
			want: 1,
		},
		{
			name: "#5",
			data: []int{1, 2, 4, 5, 6, 10, 100},
			item: 4,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Binary(tt.data, tt.item); got != tt.want {
				t.Errorf("Binary() = %v, want %v", got, tt.want)
			}
			if got := Simple(tt.data, tt.item); got != tt.want {
				t.Errorf("Simple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sampleData() []int {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		data[i] = rand.Intn(1000)
	}

	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	return data
}

func BenchmarkBinary(b *testing.B) {
	f, _ := os.Create("traceBin.out")
	defer f.Close()
	trace.Start(f)
	data := sampleData()
	for b.Loop() {
		n := rand.Intn(1000)
		res := Binary(data, n)
		_ = res
	}
	trace.Stop()
	http.ListenAndServe(":80", nil)
}

func BenchmarkSimple(b *testing.B) {
	f, _ := os.Create("traceSimp.out")
	defer f.Close()
	trace.Start(f)
	data := sampleData()
	for b.Loop() {
		n := rand.Intn(1000)
		res := Simple(data, n)
		_ = res
	}
	trace.Stop()
	http.ListenAndServe(":81", nil)
}
