package main

import (
	"bytes"
	"io"
	"testing"
)

var input = `A Y
B X
C Z`

func Test_part1(t *testing.T) {
	buf := bytes.NewBuffer([]byte(input))

	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Test part 1",
			args:    args{r: buf},
			want:    15,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("part1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	buf := bytes.NewBuffer([]byte(input))

	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Test part 2",
			args:    args{r: buf},
			want:    12,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("part2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
