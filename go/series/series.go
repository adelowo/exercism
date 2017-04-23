package slice

const testVersion = 1

func All(n int, s string) []string {
	var out []string

	for i := 0; i < len(s)-n+1; i++ {
		out = append(out, s[i:i+n])
	}

	return out
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {

	if n > len(s) || n < 0 {
		return "", false
	}

	return s[:n], true
}
