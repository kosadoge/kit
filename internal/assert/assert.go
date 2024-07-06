package assert

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Equal(t *testing.T, expect any, actual any) {
	t.Helper()

	if diff := cmp.Diff(expect, actual); diff != "" {
		t.Fatalf("mismatch (-expect +actual)\n%s", diff)
	}
}
