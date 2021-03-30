package stringscmp

import (
	"testing"
)

func TestContainsStringArrays(t *testing.T) {
	small := []string{"BBB", "CCC", "AAAA"}
	large := []string{"FFF", "BBB", "AAAA", "DDD", "CCC"}

	yes := containsStringArrays(small, large)

	if !yes {
		t.Fatalf("Should %v be in %v", small, large)
	}
}

func TestCompareStringArrays(t *testing.T) {
	small := []string{"AAAA", "BBB", "CCC"}
	large := []string{"AAAA", "BBB", "CCC"}

	yes := compareStringArrays(small, large)

	if !yes {
		t.Fatalf("Should %v be in %v", small, large)
	}
}
