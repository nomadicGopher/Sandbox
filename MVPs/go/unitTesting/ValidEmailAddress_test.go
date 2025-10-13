package unitTesting

import (
	"testing"
)

func TestValidEmailAddress(t *testing.T) {
	unitTests := []struct {
		Description string
		Address     string
		Pass        bool
	}{
		{"Invalid email address", "somebody", false},
		{"Valid email address", "somebody@gmail.com", true},
	}

	for _, unitTest := range unitTests {
		t.Run(unitTest.Description, func(t *testing.T) {
			if actual := ValidEmailAddress(unitTest.Address); actual != unitTest.Pass {
				t.Errorf("%s: %s is not valid.", unitTest.Description, unitTest.Address)
			}
		})
	}
}
