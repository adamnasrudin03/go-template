package helpers

import (
	"testing"
)

func Test_defaultTargetLang(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty input",
			args: args{
				s: "",
			},
			want: LangID,
		},
		{
			name: "not empty input",
			args: args{
				s: LangEn,
			},
			want: LangEn,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defaultTargetLang(tt.args.s); got != tt.want {
				t.Errorf("defaultTargetLang() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultSourceLang(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty input",
			args: args{
				s: "",
			},
			want: Auto,
		},
		{
			name: "not empty input",
			args: args{
				s: LangEn,
			},
			want: LangEn,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defaultSourceLang(tt.args.s); got != tt.want {
				t.Errorf("defaultSourceLang() = %v, want %v", got, tt.want)
			}
		})
	}
}
