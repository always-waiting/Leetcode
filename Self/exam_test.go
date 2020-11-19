package Self

import (
	"testing"
)

func Test_GetSub(t *testing.T) {
	GetSub([]string{"a", "b", "c"})
}

func Test_Match(t *testing.T) {
	t.Log(Match("abcde", "mmabemcdefgm"))
}

func Test_SliceAndArray(t *testing.T) {
	a1 := []string{"m", "m"}
	Slice(a1)
	t.Log(a1)
	b1 := [2]string{"m", "m"}
	Array(b1)
	t.Log(b1)
}
