package helpers

import "testing"

func TestIsValidEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success case normal email",
			args: args{
				email: "adam@email.com",
			},
			want: true,
		},
		{
			name: "success case with number symbol",
			args: args{
				email: "adam.nasrudin1234@email.com",
			},
			want: true,
		},
		{
			name: "failed case normal email",
			args: args{
				email: "adam.1234",
			},
			want: false,
		},
		{
			name: "failed case domain",
			args: args{
				email: "go.dev",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmail(tt.args.email); got != tt.want {
				t.Errorf("IsValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
