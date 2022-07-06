package util

func MapReverse(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[v] = k
	}
	return out
}
