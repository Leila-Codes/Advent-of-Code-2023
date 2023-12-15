package trebuchet

import "testing"

func TestFirstAndLast(t *testing.T) {
	first, last := FirstAndLast([]int{1, 6, 4, 3, 9, 2})

	if first != 1 {
		t.Errorf("expected first to be '%d' but received '%d'\n", 1, first)
		t.FailNow()
	}

	if last != 2 {
		t.Errorf("expected last to be '%d' but received '%d'\n", 2, last)
		t.FailNow()
	}
}
