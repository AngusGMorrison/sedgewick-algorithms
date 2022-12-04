package q21_find

import "testing"

func Test_find(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input *node[string]
		key   string
		want  bool
	}{
		{
			name:  "list is nil",
			input: nil,
			key:   "anything",
			want:  false,
		},
		{
			name: "key is not present",
			input: &node[string]{
				data: "a",
				next: &node[string]{
					data: "b",
					next: &node[string]{
						data: "c",
					},
				},
			},
			key:  "d",
			want: false,
		},
		{
			name: "key is present in first element",
			input: &node[string]{
				data: "a",
				next: &node[string]{
					data: "b",
					next: &node[string]{
						data: "c",
					},
				},
			},
			key:  "a",
			want: true,
		},
		{
			name: "key is present in middle element",
			input: &node[string]{
				data: "a",
				next: &node[string]{
					data: "b",
					next: &node[string]{
						data: "c",
					},
				},
			},
			key:  "b",
			want: true,
		},
		{
			name: "key is present in last element",
			input: &node[string]{
				data: "a",
				next: &node[string]{
					data: "b",
					next: &node[string]{
						data: "c",
					},
				},
			},
			key:  "c",
			want: true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := find(tc.input, tc.key); got != tc.want {
				t.Errorf("want %t, got %t", tc.want, got)
			}
		})
	}
}
