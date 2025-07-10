package unitTesting

import (
	"strings"
)

func ValidEmailAddress(address string) bool {
	return strings.Contains(address, "@")
}
