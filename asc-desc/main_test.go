package main

import (
	"reflect"
	"testing"
)

func TestAsc(t *testing.T) {
	actual := SortAscDesc([]int{10, 3, 4, 11, 8, -1, -10}, true)
	expected := []int{-10, -1, 3, 4, 8, 10, 11}

	if reflect.DeepEqual(actual, expected) {
		t.Logf("Result is Expected")
	} else {
		t.Fatalf("Result is not Expected")
	}
}

func TestDesc(t *testing.T) {
	actual := SortAscDesc([]int{10, 3, 4, 11, 8, -1, -10}, false)
	expected := []int{11, 10, 8, 4, 3, -1, -10}

	if reflect.DeepEqual(actual, expected) {
		t.Logf("Result is Expected")
	} else {
		t.Fatalf("Result is not Expected")
	}
}
