package utils

func Remove(a []string, i int) []string {
	return append(a[:i], a[i+1:]...)
}

func Unshift(a []string, b string) []string {
	return append([]string{b}, a...)
}