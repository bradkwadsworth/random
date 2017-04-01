package random

import (
	"math/rand"
	"testing"
)

func TestWord(t *testing.T) {
	rand.Seed(3)
	w := []string{"foo", "bar", "biz", "baz"}
	e := "foo"
	a := Word(w)
	if e != a {
		t.Errorf("%s does not equal %s\n", e, a)
	}
}
