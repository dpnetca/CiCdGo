package hello

import "testing"

func TestGreet(t *testing.T) {
	result := Greet("World")
	expectedResult := "Hello World!"
	if result != expectedResult {
		t.Errorf("Greet(\"World\") = %s; wanted %s", result, expectedResult)
	}
}
