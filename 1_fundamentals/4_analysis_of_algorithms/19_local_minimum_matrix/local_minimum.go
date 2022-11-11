package localminimum

import (
	"math"
)

// DivideAndConquer works by dividing the input matrix into quadrants along its central row and
// column. After identifying the minimum element in the central row, column or boundary of the
// matrix (where such a boundary exists), it checks whether this minimum element has any smaller
// neighbors. If the minimum element has no smaller neighbors, a local minimum has been found. A
// smaller neighbor indicates that there must be a local minimum in that quadrant, so the algorithm
// recurses into the neighbor's quadrant.
//
// The process of breaking the matrix into progressively smaller quadrants avoids the pitfall of the
// "roll downhill" method, which may visit up to half the elements in the matrix (N^2) if the path
// to the local minimum meanders significantly.
//
// See http://courses.csail.mit.edu/6.006/spring11/lectures/lec02.pdf for an illustration.
func DivideAndConquer(m [][]int) []int {
	if len(m) == 0 {
		return nil
	}

	return divideAndConquer(m, 0, len(m)-1, 0, len(m)-1)
}

func divideAndConquer(m [][]int, top, bottom, left, right int) []int {
	if bottom-top == 0 && right-left == 0 { // single-element matrix; must be local minimum
		return []int{top, left}
	}

	minRow, minCol := frameMin(m, top, bottom, left, right)
	if minRow > top && m[minRow-1][minCol] < m[minRow][minCol] { // check if top neighbor is smaller
		minRow--
	} else if minRow < bottom && m[minRow+1][minCol] < m[minRow][minCol] { // check if bottom neighbor is smaller
		minRow++
	} else if minCol > left && m[minRow][minCol-1] < m[minRow][minCol] { // check if left neighbor is smaller
		minCol--
	} else if minCol < right && m[minRow][minCol+1] < m[minRow][minCol] { // check if right neighbor is smaller
		minCol++
	} else { // no neighbors are smaller; local minimum found
		return []int{minRow, minCol}
	}

	// Determine new frame.
	mid := top + (bottom-top)/2
	if minRow < mid {
		bottom = mid - 1
	} else if minRow > mid {
		top = mid + 1
	}
	if minCol < mid {
		right = mid - 1
	} else if minCol > mid {
		left = mid + 1
	}

	return divideAndConquer(m, top, bottom, left, right)
}

// frameMin returns the index (i,j) of the minimum value in frame (the central row, central column or
// boundary) of the submatrix specified by top, bottom, left and right.
func frameMin(m [][]int, top, bottom, left, right int) (int, int) {
	mid := top + (bottom-top)/2
	minVal := math.MaxInt
	var minRow, minCol int
	for col := left; col <= right; col++ { // middle row
		if m[mid][col] < minVal {
			minVal = m[mid][col]
			minRow = mid
			minCol = col
		}
	}
	for row := top; row <= bottom; row++ { // middle column
		if m[row][mid] < minVal {
			minVal = m[row][mid]
			minRow = row
			minCol = mid
		}
	}

	if top > 0 { // boundary top edge (inc. left and right boundary corners)
		topBound := top - 1
		leftBound := left
		if leftBound > 0 {
			leftBound--
		}
		rightBound := right
		if rightBound < len(m)-1 {
			rightBound++
		}
		for col := leftBound; col <= rightBound; col++ {
			if m[topBound][col] < minVal {
				minVal = m[topBound][col]
				minRow = topBound
				minCol = col
			}
		}
	}

	if bottom < len(m)-1 { // boundary bottom edge (inc. left and right boundary corners)
		bottomBound := bottom + 1
		leftBound := left
		if leftBound > 0 {
			leftBound--
		}
		rightBound := right
		if rightBound < len(m)-1 {
			rightBound++
		}
		for col := leftBound; col <= rightBound; col++ {
			if m[bottomBound][col] < minVal {
				minVal = m[bottomBound][col]
				minRow = bottomBound
				minCol = col
			}
		}
	}

	if left > 0 { // boundary left edge
		leftBound := left - 1
		for row := top; row <= bottom; row++ {
			if m[row][leftBound] < minVal {
				minVal = m[row][leftBound]
				minRow = row
				minCol = leftBound
			}
		}
	}

	if right < len(m)-1 { // boundary right edge
		rightBound := right + 1
		for row := top; row <= bottom; row++ {
			if m[row][rightBound] < minVal {
				minVal = m[row][rightBound]
				minRow = row
				minCol = rightBound
			}
		}
	}

	return minRow, minCol
}

// RollDownhill starts from the middle element of the matrix and proceeds to its smallest neighbor
// on each loop iteration until the current element is smaller than each of its neighbors. This approach may visit up to half the elements in the array, so it is ~N^2.
func RollDownhill(m [][]int) []int {
	size := len(m)
	if size == 0 {
		return nil
	}
	if size == 1 {
		return []int{0, 0}
	}

	row := (len(m) - 1) / 2
	col := row
	for elem := m[row][col]; ; elem = m[row][col] {
		// Handle corners.
		if row == 0 && col == 0 { // top-left
			bottom := m[row+1][col]
			right := m[row][col+1]

			if elem < bottom && elem < right {
				return []int{row, col} // local minimum found
			}

			if bottom < right { // move to the smallest neighbor
				row++
			} else {
				col++
			}
			continue
		}

		if row == size-1 && col == 0 { // bottom-left
			top := m[row-1][col]
			right := m[row][col+1]

			if elem < top && elem < right {
				return []int{row, col} // local minimum found
			}

			if top < right {
				row--
			} else {
				col++
			}
			continue
		}

		if row == 0 && col == size-1 { // bottom-right
			bottom := m[row+1][col]
			left := m[row][col-1]

			if elem < bottom && elem < left {
				return []int{row, col} // local minimum found
			}

			if bottom < left {
				row++
			} else {
				col--
			}
			continue
		}

		if row == size-1 && col == size-1 { // bottom-right
			top := m[row-1][col]
			left := m[row][col-1]

			if elem < top && elem < left {
				return []int{row, col} // local minimum found
			}

			if top < left { // move to the smallest neighbor
				row--
			} else {
				col--
			}
			continue
		}

		// Handle sides.
		if row == 0 { // top
			bottom := m[row+1][col]
			left := m[row][col-1]
			right := m[row][col+1]

			if elem < bottom && elem < left && elem < right {
				return []int{row, col} // local minimum found
			}

			switch min(bottom, left, right) {
			case bottom:
				row++
			case left:
				col--
			case right:
				col++
			}
			continue
		}

		if row == size-1 { // bottom
			top := m[row-1][col]
			left := m[row][col-1]
			right := m[row][col+1]

			if elem < top && elem < left && elem < right {
				return []int{row, col} // local minimum found
			}

			switch min(top, left, right) {
			case top:
				row--
			case left:
				col--
			case right:
				col++
			}
			continue
		}

		if col == 0 { // left
			top := m[row-1][col]
			bottom := m[row+1][col]
			right := m[row][col+1]

			if elem < top && elem < bottom && elem < right {
				return []int{row, col} // local minimum found
			}

			switch min(top, bottom, right) {
			case top:
				row--
			case bottom:
				row++
			case right:
				col++
			}
			continue
		}

		if col == size-1 { // right
			top := m[row-1][col]
			bottom := m[row+1][col]
			left := m[row][col-1]

			if elem < top && elem < bottom && elem < left {
				return []int{row, col} // local minimum found
			}

			switch min(top, bottom, left) {
			case top:
				row--
			case bottom:
				row++
			case left:
				col--
			}
			continue
		}

		// Handle middle element.
		top := m[row-1][col]
		bottom := m[row+1][col]
		left := m[row][col-1]
		right := m[row][col+1]

		if elem < top && elem < bottom && elem < left && elem < right {
			return []int{row, col} // local minimum found
		}

		switch min(top, bottom, left, right) {
		case top:
			row--
		case bottom:
			row++
		case left:
			col--
		case right:
			col++
		}
	}
}

func min(ints ...int) int {
	min := math.MaxInt
	for _, n := range ints {
		if n < min {
			min = n
		}
	}

	return min
}
