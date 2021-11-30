package utils

// InSlice is find string in the string slice
func InSlice(slice []string, str string) (index int, exists bool) {
	for i, val := range slice {
		if val == str {
			return i, true
		}
	}
	return -1, false
}
