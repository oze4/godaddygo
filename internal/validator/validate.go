package validator

// Validate is a wrapper around validating params (due to no enums)
func Validate(s string, m map[string]string) bool {
	for t := range m {
		if s == t {
			return true
		}
	}
	return false
}