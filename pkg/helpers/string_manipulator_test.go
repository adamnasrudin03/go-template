package helpers

import (
	"testing"
)

var (
	helloWorld         = "Hello World"
	helloWorldLower    = "hello world"
	helloWorldUpper    = "HELLO WORLD"
	helloWorldSentence = "Hello world"
)

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

func TestToLower(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				input: helloWorld,
			},
			want: helloWorldLower,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLower(tt.args.input); got != tt.want {
				t.Errorf("ToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToTitle(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				input: helloWorldLower,
			},
			want: helloWorld,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToTitle(tt.args.input); got != tt.want {
				t.Errorf("ToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				input: helloWorld,
			},
			want: helloWorldUpper,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpper(tt.args.input); got != tt.want {
				t.Errorf("ToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSentenceCase(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				input: helloWorld,
			},
			want: helloWorldSentence,
		},
		{
			name: "empty strings",
			args: args{
				input: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSentenceCase(tt.args.input); got != tt.want {
				t.Errorf("ToSentenceCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
