package string_test

import (
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))

	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s))

	c := []rune(s)

	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

	testTasks := []string{}
	t.Log(testTasks == nil)
}
