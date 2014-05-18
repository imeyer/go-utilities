package hostname

import (
	"strings"
)

func Reverse(hostname string) string {
	chunks := strings.Split(hostname, ".")

	for i, j := 0, len(chunks)-1; i < j; i, j = i+1, j-1 {
		chunks[i], chunks[j] = chunks[j], chunks[i]
	}

	reverse_hostname := strings.Join(chunks, ".")

	return reverse_hostname
}

func ReverseOffset(hostname string, offset int8) string {
	chunks := strings.Split(hostname, ".")

	for i, j := 0, len(chunks)-1; i < j; i, j = i+1, j-1 {
		chunks[i], chunks[j] = chunks[j], chunks[i]
	}

	reverse_hostname := strings.Join(chunks[offset:], ".")

	return reverse_hostname
}
