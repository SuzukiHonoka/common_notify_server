package utils

// IsEmpty check any string is empty
func IsEmpty(s ...string) bool {
	for _, ss := range s {
		if len(ss) == 0 {
			return true
		}
	}
	return false
}

// IsNotEmpty check if all string not empty
func IsNotEmpty(s ...string) bool {
	for _, ss := range s {
		if len(ss) == 0 {
			return false
		}
	}
	return true
}
