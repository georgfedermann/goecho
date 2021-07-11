package echolib_test

import (
	"testing"

	"github.com/georgfedermann/goecho/echolib"
)

func TestReadArgsIntoString(t *testing.T) {
	expect := "1 2 3 4 5"
	actual := echolib.ReadArgsIntoString([]string{"1", "2", "3", "4", "5"})
	if expect != actual {
		t.Fatal("expected:", expect, "but was:", actual)
	}
}
