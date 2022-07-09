package lc

import (
	"reflect"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name      string
		args      args
		wantIndex int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			wantIndex: 3,
		},
		{
			name: "2",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			wantIndex: -1,
		},
		{
			name: "2",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			wantIndex: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndex := binarySearchLeftFirst(tt.args.nums, tt.args.target); gotIndex != tt.wantIndex {
				t.Errorf("binarySearchLeftFirst() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums:   nil,
				target: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetMiddle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMiddle(3068, 2341)
	}
}

func BenchmarkGetMiddleB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMiddleB(3068, 2341)
	}
}
