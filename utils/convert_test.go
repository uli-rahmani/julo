package utils

import "testing"

func TestGetInt(t *testing.T) {
	type args struct {
		x string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "error",
			args: args{
				x: "lala",
			},
			want: 0,
		},
		{
			name: "success",
			args: args{
				x: "2",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInt(tt.args.x); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
