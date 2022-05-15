package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	val := "Hello world!"

	expected := "!dlrow olleH"
	got := ReverseString(val)

	if expected != got {
		t.Errorf("reverseString(), got: %v, expected: %v", got, expected)
	}
}
