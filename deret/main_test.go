package main

import (
	"reflect"
	"testing"
)

func TestDeretGenap(t *testing.T) {
	actual := Deret(2, 4, 5)
	expected := []int{2, 4, 6, 8, 10}

	if reflect.DeepEqual(actual, expected) {
		t.Logf("Result is Expected")
	} else {
		t.Fatalf("Result is not Expected")
	}
}

func TestDeretGanjil(t *testing.T) {
	actual := Deret(5, 8, 7)
	expected := []int{5, 8, 11, 14, 17, 20, 23}

	if reflect.DeepEqual(actual, expected) {
		t.Logf("Result is Expected")
	} else {
		t.Fatalf("Result is not Expected")
	}
}
