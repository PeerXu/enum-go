package enum

import (
	"fmt"
	"strings"
)

func GenerateGetByName[T interface {
	~int
	fmt.Stringer
}](min, max T) func(string) (T, bool) {
	reversed := map[string]T{}
	for i := min; i <= max; i++ {
		reversed[strings.ToLower(i.String())] = i
	}
	return func(s string) (v T, ok bool) {
		v, ok = reversed[strings.ToLower(s)]
		return
	}
}
