package users

import "testing"

func TestElder(t *testing.T) {
	type args struct {
		u []User
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test #1",
			args{[]User{Employee{33}, Customer{44}}},
			44,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Elder(tt.args.u...); got != tt.want {
				t.Errorf("Elder() = %v, want %v", got, tt.want)
			}
		})
	}
}
