package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3}, 100, nil},
		{[]int{}, 5, nil},
		{[]int{-3, 4, 3, 90}, 0, []int{0, 2}},
	}
	for _, c := range cases {
		if got := TwoSum(c.in, c.target); !reflect.DeepEqual(got, c.want) {
			t.Errorf("TwoSum(%v,%d) = %v, want %v", c.in, c.target, got, c.want)
		}
	}
}
