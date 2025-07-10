package unitTesting

import "testing"

// Singular test
func TestAddTwoNumbers_Single(t *testing.T) {
	a := 2
	b := 3
	actualOutput := AddTwoNumbers(a, b)

	// failingExpectedOutput := 99 // Fail first
	// if actualOutput != failingExpectedOutput {
	// 	t.Errorf("AddTwoNumbers() = %v; want %v.", actualOutput, failingExpectedOutput)
	// }

	passingExpectedOutput := 5
	if actualOutput != passingExpectedOutput {
		t.Errorf("AddTwoNumbers() = %v; want %v.", actualOutput, passingExpectedOutput)
	}
}

func TestAddTwoNumbers_Multiple(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{"Fail first", args{1, 1}, 99},
		{"Sample pass", args{2, 2}, 4},
		{"Zero plus zero", args{0, 0}, 0},
		{"Negative numbers", args{-1, -2}, -3},
		{"Positive and negative", args{5, -3}, 2},
		{"Large numbers", args{1000000, 2000000}, 3000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ------------------------------------------

func TestRestaurantWithCustomEquals(t *testing.T) {
	expected := Restaurant{Name: "Joes Pizza", Cuisine: "Italian"}
	actual := Restaurant{Name: "Joes Pizza", Cuisine: "Italian"}

	if !expected.Equal(actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Negative test: different name
	actual2 := Restaurant{Name: "Mama Putt", Cuisine: "Italian"}
	if expected.Equal(actual2) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Negative test: different cuisine
	actual3 := Restaurant{Name: "Joes Pizza", Cuisine: "Indian"}
	if expected.Equal(actual3) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
