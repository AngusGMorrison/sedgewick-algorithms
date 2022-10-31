package binary

import "testing"

func Test_toBinaryString(t *testing.T) {
	t.Parallel()

	want := "110010"
	got := toBinaryString(50)
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
