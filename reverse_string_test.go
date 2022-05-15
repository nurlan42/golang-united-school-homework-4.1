package reverse_string

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	val := "привет"
	expected := []rune("тевирп")

	got := []rune(ReverseString(val))

	for i := 0; i < len(got); i++ {
		for j := i; j < len(expected)-1; j++ {
			if got[i] == expected[j] {
				break
			} else {
				t.Errorf("reverseString(), got: %v, expected: %v", got, expected)
			}
		}
	}
}
