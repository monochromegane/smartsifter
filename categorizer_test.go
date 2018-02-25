package smartsifter

import "testing"

func TestNoOpCategorizerIndex(t *testing.T) {
	c := noOpCategorizer{}

	expect := 10
	idx := c.Index([]int{expect, 20, 30})
	if idx != expect {
		t.Errorf("noOpCategorizer Index should return first element, but %d", idx)
	}
}

func TestNoOpCategorizerSize(t *testing.T) {
	c := noOpCategorizer{}
	size := c.Size(0)
	if size != 1 {
		t.Errorf("noOpCategorizer Size should always return 1, but %d", size)
	}
}
