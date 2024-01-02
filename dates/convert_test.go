package dates

import (
	"reflect"
	"testing"
	"time"
)

func TestArgToEpoch(t *testing.T) {
	type args struct {
		arg     string
		backoff int64
	}
	tests := []struct {
		name      string
		args      args
		wantTime  time.Time
		wantEpoch int64
		wantErr   bool
	}{
		{
			name:      "vanilla",
			args:      args{arg: "2023-12-01T00:00:00"},
			wantEpoch: 1701388800, // UTC timestamp in seconds for 2023-12-01 midnight
			wantTime:  time.UnixMilli(1701388800 * 1000),
		},
		{
			name:      "default",
			args:      args{arg: "", backoff: 0},
			wantEpoch: 1701388800,
			wantTime:  time.UnixMilli(1701388800 * 1000),
		},
		{
			name:      "backoff",
			args:      args{arg: "", backoff: 42},
			wantEpoch: 1701388800 - 42,
			wantTime:  time.UnixMilli((1701388800 - 42) * 1000),
		},
		{
			name:      "noTime",
			args:      args{arg: "2023-12-01"},
			wantEpoch: 1701388800,
			wantTime:  time.UnixMilli(1701388800 * 1000),
		},
		{
			name:      "noSecs",
			args:      args{arg: "2023-12-01T00:00"},
			wantEpoch: 1701388800,
			wantTime:  time.UnixMilli(1701388800 * 1000),
		},
		{
			name:      "noMins",
			args:      args{arg: "2023-12-01T00"},
			wantEpoch: 1701388800,
			wantTime:  time.UnixMilli(1701388800 * 1000),
		},
		{
			name:      "noYear",
			args:      args{arg: "12-01"}, // TODO this assumes it's 2023
			wantEpoch: 1701388800,
			wantTime:  time.UnixMilli(1701388800 * 1000),
		},
	}
	areTesting = true
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTime, gotEpoch, err := ConvertTimeString(tt.args.arg, tt.args.backoff)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArgToEpoch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotEpoch != tt.wantEpoch {
				t.Errorf("ArgToEpoch() epoch = %v, want %v", gotEpoch, tt.wantEpoch)
			}
			if !gotTime.Equal(tt.wantTime) {
				t.Errorf("ArgToEpoch() time = %v, want %v", gotTime, tt.wantTime)
			}
		})
	}
}

func Test_getTimeFromEpochString(t *testing.T) {
	type args struct {
		epochString string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "empty",
			args: args{""},
			want: time.Time{},
		},
		{
			name: "zero",
			args: args{"0"},
			want: time.Time{},
		},
		{
			name:    "negtive",
			args:    args{"-42"},
			wantErr: true,
		},
		{
			name: "midnite",
			args: args{"1701388800"},
			want: time.UnixMilli(1701388800 * 1000).UTC(),
		},
		{
			name:    "baddog",
			args:    args{"badbeef"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTimeFromEpochString(tt.args.epochString)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTimeFromEpochString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTimeFromEpochString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTimesFromQuery(t *testing.T) {
	type args struct {
		startTimeString string
		endTimeString   string
		backoff         int
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		want1   time.Time
		wantErr bool
	}{
		{
			name: "good",
			args: args{
				startTimeString: "",
				endTimeString:   "1701388800",
				backoff:         120,
			},
			want:  time.UnixMilli((1701388800 - 120) * 1000).UTC(),
			want1: time.UnixMilli(1701388800 * 1000).UTC(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getTimesFromQuery(tt.args.startTimeString, tt.args.endTimeString, tt.args.backoff)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTimesFromQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTimesFromQuery() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getTimesFromQuery() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
