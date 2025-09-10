package hw

import "testing"

func TestGeom_Distance(t *testing.T) {
	tests := []struct {
		name         string
		points       [4]float64
		wantDistance float64
	}{
		{
			name:         "#1",
			points:       [4]float64{1, 1, 4, 5},
			wantDistance: 5,
		},
		{
			name:         "#2",
			points:       [4]float64{-1, -1, 2, 3},
			wantDistance: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := Distance(tt.points[0], tt.points[1], tt.points[2], tt.points[3]); gotDistance != tt.wantDistance {
				t.Errorf("Geom.CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
