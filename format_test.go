package format

import (
	"testing"
)

func TestListOr(t *testing.T) {
	{
		g := ListOr("foo", "bar", "baz")
		w := "foo, bar or baz"
		if g != w {
			t.Errorf("wrong output: want: %q got %q", w, g)
		}
	}
	{
		OrTerm = "oder"
		g := ListOr("foo", "bar", "baz")
		w := "foo, bar oder baz"
		if g != w {
			t.Errorf("wrong output: want: %q got %q", w, g)
		}
	}
}
