package unitTesting

func AddTwoNumbers(a, b int) int {
	return a + b
}

// ------------------------------------------

type Restaurant struct {
	Name    string
	Cuisine string
}

func (r Restaurant) Equal(other Restaurant) bool {
	return r.Name == other.Name && r.Cuisine == other.Cuisine
}
