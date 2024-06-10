package helpers

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckTimeIsZeroToPointer(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want *time.Time
	}{
		{
			name: "success",
			args: args{
				t: nowStartDate,
			},
			want: &nowStartDate,
		},
		{
			name: "zero value",
			args: args{
				t: time.Time{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckTimeIsZeroToPointer(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckTimeIsZeroToPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckTimePointerValue(t *testing.T) {
	type args struct {
		t *time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "success",
			args: args{
				t: &nowStartDate,
			},
			want: nowStartDate,
		},
		{
			name: "nil param",
			args: args{
				t: nil,
			},
			want: time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckTimePointerValue(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckTimePointerValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckTimeIsZeroToString(t *testing.T) {
	type args struct {
		t          time.Time
		formatDate string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				t:          now,
				formatDate: FormatDate,
			},
			want: now.Format(FormatDate),
		},
		{
			name: "zero value",
			args: args{
				t:          time.Time{},
				formatDate: FormatDate,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckTimeIsZeroToString(tt.args.t, tt.args.formatDate); got != tt.want {
				t.Errorf("CheckTimeIsZeroToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckTimePointerToString(t *testing.T) {
	type args struct {
		t          *time.Time
		formatDate string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				t:          &now,
				formatDate: FormatDate,
			},
			want: now.Format(FormatDate),
		},
		{
			name: "nil value",
			args: args{
				t:          nil,
				formatDate: FormatDate,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckTimePointerToString(tt.args.t, tt.args.formatDate); got != tt.want {
				t.Errorf("CheckTimePointerToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
