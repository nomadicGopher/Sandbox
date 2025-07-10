package unitTesting

import "testing"

// Singular test
func TestAddTwoNumbers_Single(t *testing.T) {
	a := 2
	b := 3
	actualOutput := AddTwoNumbers(a, b)

	// failingExpectedOutput := 99 // Fail first
	// if actualOutput != failingExpectedOutput {
	// 	t.Errorf("AddTwoNumbers() = %v; expected %v.", actualOutput, failingExpectedOutput)
	// }

	passingExpectedOutput := 5
	if actualOutput != passingExpectedOutput {
		t.Errorf("AddTwoNumbers() = %v; expected %v.", actualOutput, passingExpectedOutput)
	}
}

func TestAddTwoNumbers_Multiple(t *testing.T) {
	type args struct {
		a int
		b int
	}
	testCases := []struct {
		name     string
		args     args
		expected int
		pass     bool
	}{
		{"Fail first", args{1, 1}, 99, false},
		{"Sample pass", args{2, 2}, 4, true},
		{"Zero plus zero", args{0, 0}, 0, true},
		{"Negative numbers", args{-1, -2}, -3, true},
		{"Positive and negative", args{5, -3}, 2, true},
		{"Large numbers", args{1000000, 2000000}, 3000000, true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if actual := AddTwoNumbers(testCase.args.a, testCase.args.b); actual != testCase.expected && testCase.pass != false {
				t.Errorf("AddTwoNumbers() = %v, expected %v", actual, testCase.expected)
			}
		})
	}
}

// ------------------------------------------

func TestRestaurantWithPointerReciever(t *testing.T) {
	expected := Restaurant{Name: "Joes Pizza", Cuisine: "Italian"}
	actual := Restaurant{Name: "Joes Pizza", Cuisine: "Italian"}

	if !expected.Equal(actual) {
		t.Errorf("Expected restaurants to be equal: expected=%v, actual=%v", expected, actual)
	}

	// Negative test: different name
	diffName := Restaurant{Name: "Mama Putt", Cuisine: "Italian"}
	if expected.Equal(diffName) {
		t.Errorf("Expected restaurants to be different by Name: expected=%v, actual=%v", expected, diffName)
	}
}
