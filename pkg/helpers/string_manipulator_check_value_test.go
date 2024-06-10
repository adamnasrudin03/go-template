package helpers

import (
	"reflect"
	"testing"
)

func TestCheckStringValueToPointer(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "empty string",
			args: args{
				data: "",
			},
			want: nil,
		},
		{
			name: "success",
			args: args{
				data: helloWorld,
			},
			want: &helloWorld,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckStringValueToPointer(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckStringValueToPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckStringValue(t *testing.T) {
	type args struct {
		data *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nil value",
			args: args{
				data: nil,
			},
			want: "",
		},
		{
			name: "success",
			args: args{
				data: &helloWorld,
			},
			want: helloWorld,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckStringValue(tt.args.data); got != tt.want {
				t.Errorf("CheckStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
