package helpers

import (
	"reflect"
	"testing"
	"time"
)

var (
	now          = time.Now()
	nowStartDate = StartDate(now)
	nowEndDate   = nowStartDate.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	nowUTC7, _   = time.ParseInLocation(FormatDateTime, now.Format(FormatDateTime), loc)
)

func TestParseUTC7(t *testing.T) {
	type args struct {
		timeFormat string
		value      string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				timeFormat: FormatDateTime,
				value:      now.Format(FormatDateTime),
			},
			want:    nowUTC7,
			wantErr: false,
		},
		{
			name: "err value invalid format",
			args: args{
				timeFormat: FormatDate,
				value:      now.Format(FormatDateTime),
			},
			want:    now,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUTC7(tt.args.timeFormat, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUTC7() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUTC7() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndDate(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "success",
			args: args{
				t: now,
			},
			want: nowEndDate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndDate(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndDateString(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "success",
			args: args{
				t: now.Format(FormatDate),
			},
			want: nowEndDate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndDateString(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndDateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
