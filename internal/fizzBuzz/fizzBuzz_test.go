package fizzBuzz

import (
	"reflect"
	"testing"
)

// TestFizzBuzz calls fizzBuzz.GetFizzBuzz with simple params, checking
// for a valid return value.
func TestFizzBuzz(t *testing.T) {
	int1 := 3
	int2 := 5
	limit := 10
	str1 := "fizz"
	str2 := "buzz"

	want := []string{
		"1",
		"2",
		"fizz",
		"4",
		"buzz",
		"fizz",
		"7",
		"8",
		"fizz",
	}
	rslt, err := GetFizzBuzz(int1, int2, limit, str1, str2)
	if !reflect.DeepEqual(want, rslt) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, rslt, err, want)
	}
}
