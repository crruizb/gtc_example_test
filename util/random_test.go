package util

import (
	"testing"
)

func TestRandomStr(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Hpath",
			want: "xyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomStr(); got != tt.want {
				t.Errorf("RandomStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Happy path",
			args: args{
				a: 2,
				b: 5,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("CalcSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
