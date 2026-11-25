//
// Объединение интервалов.
//
// Два интервала объединяются, если начало следующего не больше конца текущего:
// пересечение и касание концами считаются одним интервалом.
// Вложенный интервал поглощается внешним.
//
// Примеры:
//
//	[1,3], [2,5] -> [1,5]          перекрытие
//	[1,3], [3,5] -> [1,5]          касание — объединение
//	[1,5], [2,3] -> [1,5]          вложенность
//	[1,3], [4,5] -> [1,3], [4,5]   разрыв: концы не совпадают и не пересекаются

package algorithms

import (
	"cmp"
	"slices"
)

type Interval struct {
	left  int
	right int
}

type MergeIntervals struct{}

func (m MergeIntervals) MergeIntervals(input []Interval) []Interval {

	if input == nil {
		return nil
	}

	if len(input) < 1 {
		return []Interval{}
	}

	sorted := slices.Clone(input)

	slices.SortFunc(sorted, func(a Interval, b Interval) int {
		return cmp.Compare(a.left, b.left)
	})

	left := sorted[0].left
	right := sorted[0].right

	result := make([]Interval, 0)

	for i := 1; i < len(sorted); i++ {
		if right < sorted[i].left {
			result = append(result, Interval{left, right})
			left = sorted[i].left
			right = sorted[i].right
		} else {
			right = max(right, sorted[i].right)
		}
	}

	result = append(result, Interval{left, right})

	return result
}
