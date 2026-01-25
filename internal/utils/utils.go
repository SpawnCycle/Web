package utils

// Iterators still suck in go so we're stuck with this :((
func Map[T, U any](in []T, f func(T) U) (out []U) {
	for _, val := range in {
		out = append(out, f(val))
	}
	return out
}
