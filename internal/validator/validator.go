package validator

import "regexp"

const titleReg = "^[a-z](?:_?[a-z0-9]+)*$"

// ValidateTitle valides given string
//
// Pre-cond: given title. Title must contain only letters and numbers
// Title can't start with numbers
//
// Post-cond: return true if title is valid
// Otherwise return false
func ValidateTitle(title string) bool {
	b, err := regexp.MatchString(titleReg, title)
	if err != nil {
		return false
	}

	return b
}
