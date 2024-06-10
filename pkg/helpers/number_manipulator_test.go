package helpers

import (
	"reflect"
	"testing"
)

func TestRoundFloat(t *testing.T) {
	type args struct {
		val       float64
		precision uint
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "case 1 convert decimal 1 digit val 1 digit",
			args: args{
				val:       12.5,
				precision: 1,
			},
			want: 12.5,
		},
		{
			name: "case 2 convert decimal 1 digit val 1 digit",
			args: args{
				val:       12.3,
				precision: 1,
			},
			want: 12.3,
		},
		{
			name: "case 3 convert decimal 1 digit val 1 digit",
			args: args{
				val:       12.9,
				precision: 1,
			},
			want: 12.9,
		},
		{
			name: "case 4 convert decimal 1 digit val 1 digit",
			args: args{
				val:       12.0,
				precision: 1,
			},
			want: 12,
		},
		{
			name: "case convert decimal 1 digit",
			args: args{
				val:       12.3456789,
				precision: 1,
			},
			want: 12.3,
		},
		{
			name: "case convert decimal 2 digits",
			args: args{
				val:       12.3456789,
				precision: 2,
			},
			want: 12.35,
		},
		{
			name: "case convert decimal 3 digits",
			args: args{
				val:       12.3456789,
				precision: 3,
			},
			want: 12.346,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundFloat(tt.args.val, tt.args.precision); got != tt.want {
				t.Errorf("RoundFloat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckArrayIntNil(t *testing.T) {
	input := []int{5, 4, 3, 2, 1, 0}
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty value",
			args: args{
				data: []int{},
			},
			want: []int{},
		},
		{
			name: "success",
			args: args{
				data: input,
			},
			want: input,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckArrayIntNil(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckArrayIntNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckIntValue(t *testing.T) {
	input := 5
	type args struct {
		data *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nil value",
			args: args{
				data: nil,
			},
			want: 0,
		},
		{
			name: "success",
			args: args{
				data: &input,
			},
			want: input,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckIntValue(tt.args.data); got != tt.want {
				t.Errorf("CheckIntValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
