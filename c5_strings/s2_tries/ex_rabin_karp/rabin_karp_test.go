package ex_rabin_karp

import "testing"

func Test_IndexOf(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		pat     string
		text    string
		wantIdx int
	}{
		{
			name:    "pattern longer than text",
			pat:     "abcd",
			text:    "abc",
			wantIdx: -1,
		},
		{
			name:    "pattern not present",
			pat:     "cv",
			text:    "abc",
			wantIdx: -1,
		},
		{
			name:    "pattern present at start",
			pat:     "ab",
			text:    "abc",
			wantIdx: 0,
		},
		{
			name:    "pattern present at end",
			pat:     "bc",
			text:    "abc",
			wantIdx: 1,
		},
		{
			name:    "pattern present in middle",
			pat:     "bc",
			text:    "abcd",
			wantIdx: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := IndexOf(tc.pat, tc.text)
			if tc.wantIdx != got {
				t.Errorf("want %d, got %d", tc.wantIdx, got)
			}

		})
	}
}

func BenchmarkIndexOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexOf("xyz", "abcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwxyz")
	}
}

func BenchmarkIndexOfAlternative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexOf("xyz", "abcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwabcdefghijklmnopqrstuvwxyz")
	}
}
