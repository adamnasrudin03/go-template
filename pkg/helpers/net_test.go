package helpers

import "testing"

func TestQueryEscape(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				s: "hello world",
			},
			want: "hello+world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryEscape(tt.args.s); got != tt.want {
				t.Errorf("QueryEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}
