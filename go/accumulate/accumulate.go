package accumulate

const testVersion = 1

func Accumulate(s []string, converter func(s string) string) []string {
	var vals []string

	for _, v := range s {
		vals = append(vals, converter(v))
	}

	return vals
}
