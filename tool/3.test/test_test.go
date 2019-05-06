package main

import "testing"

func TestAdd(t *testing.T) {
	testTables := []struct {
		a int
		b int
		t int
	}{
		{1, 1, 2},
		{2, 2, 4},
	}

	for _, testTable := range testTables {
		total := add(testTable.a, testTable.b)
		if total != testTable.t {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", testTable.a, testTable.b, total, testTable.t)
		}
	}
}
