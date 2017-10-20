package calc

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	res := Sum(1, 2)
	if res != 3 {
		t.Errorf("Sum failed, got: %d, want: %d", res, 3)
	}
}

func TestTableSum(t *testing.T) {
	table := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, entry := range table {
		total := Sum(entry.x, entry.y)
		if total != entry.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", entry.x, entry.y, total, entry.n)
		}
	}
}

func TestTableSumParallel(t *testing.T) {
	//t.Parallel()

	table := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, entry := range table {
		t.Run(fmt.Sprintf("Sum of %d and %d", entry.x, entry.y), func(st *testing.T) {
			st.Parallel()

			total := SumSlow(entry.x, entry.y)
			if total != entry.n {
				t.Errorf("Sum failed, got: %d, want: %d", total, entry.n)
			}
		})
	}
}
