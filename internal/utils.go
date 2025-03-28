package utils

// some returns true if at least one element in the slice is true
func Some(bools []bool) bool {
	for _, b := range bools {
		if b {
			return true
		}
	}
	return false
}

// any is an alias for some
func Any(bools []bool) bool {
	return Some(bools)
}
