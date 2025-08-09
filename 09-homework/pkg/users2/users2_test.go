package users2

import (
	"reflect"
	"testing"
)

func TestElderEmp(t *testing.T) {
	type args struct {
		u []Employee
	}
	tests := []struct {
		name string
		args args
		want Employee
	}{
		{
			"Test #1",
			args{[]Employee{{22}, {33}, {44}, {55}}},
			Employee{55},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Elder(tt.args.u...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Elder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElderCust(t *testing.T) {
	type args struct {
		u []Customer
	}
	tests := []struct {
		name string
		args args
		want Customer
	}{
		{
			"Test #1",
			args{[]Customer{{22}, {33}, {44}, {55}}},
			Customer{55},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Elder(tt.args.u...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Elder() = %v, want %v", got, tt.want)
			}
		})
	}
}
