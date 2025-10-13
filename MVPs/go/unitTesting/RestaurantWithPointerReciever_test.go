package unitTesting

import "testing"

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
