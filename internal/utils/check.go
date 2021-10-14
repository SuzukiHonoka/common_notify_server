package utils

func IsEmpty(s ...string) bool {
	for _, ss := range s {
		if len(ss) == 0 {
			return true
		}
	}
	return false
}

func IsNotEmpty(s ...string) bool {
	for _, ss := range s {
		if len(ss) == 0 {
			return false
		}
	}
	return true
}
