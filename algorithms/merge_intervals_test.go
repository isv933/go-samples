package algorithms

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		name   string
		input  []Interval
		expect []Interval
	}{
		{
			name:   "nil input1",
			input:  nil,
			expect: nil,
		},
		{
			name:   "empty input",
			input:  []Interval{},
			expect: []Interval{},
		},
		{
			name:   "single interval",
			input:  []Interval{{2, 4}},
			expect: []Interval{{2, 4}},
		},
		{
			name:   "full overlapped intervals",
			input:  []Interval{{1, 3}, {2, 4}},
			expect: []Interval{{1, 4}},
		},
		{
			name:   "not overlapped intervals",
			input:  []Interval{{1, 3}, {4, 5}},
			expect: []Interval{{1, 3}, {4, 5}},
		},
		{
			name:   "partly overlapped intervals",
			input:  []Interval{{1, 3}, {4, 6}, {4, 5}, {7, 8}},
			expect: []Interval{{1, 3}, {4, 6}, {7, 8}},
		},
		{
			name:   "not sorted overlapped intervals",
			input:  []Interval{{2, 8}, {1, 4}},
			expect: []Interval{{1, 8}},
		},
		{
			name:   "touching intervals merge",
			input:  []Interval{{1, 3}, {3, 5}},
			expect: []Interval{{1, 5}},
		},
		{
			name:   "not sorted touching intervals merge",
			input:  []Interval{{3, 5}, {1, 3}},
			expect: []Interval{{1, 5}},
		},
		{
			name:   "nested interval is absorbed",
			input:  []Interval{{1, 5}, {2, 3}},
			expect: []Interval{{1, 5}},
		},
		{
			name:   "same left edge keeps max right",
			input:  []Interval{{1, 5}, {1, 3}},
			expect: []Interval{{1, 5}},
		},
		{
			name:   "negative bounds",
			input:  []Interval{{-5, -1}, {-3, 2}},
			expect: []Interval{{-5, 2}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expect, MergeIntervals{}.MergeIntervals(tc.input))
		})
	}
}

func TestMergeIntervalsDoesNotMutateInput(t *testing.T) {
	input := []Interval{{4, 6}, {1, 3}, {2, 5}}
	original := slices.Clone(input)

	MergeIntervals{}.MergeIntervals(input)

	require.Equal(t, original, input)
}
