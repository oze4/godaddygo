package godaddygo

func validate(s string, m map[string]string) bool {
	for t := range m {
		if s == t {
			return true
		}
	}
	return false
}
