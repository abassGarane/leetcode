package binarysearch

// You are given an m x n integer matrix matrix with the following two properties:

// Each row is sorted in non-decreasing order.
// The first integer of each row is greater than the last integer of the previous row.
// Given an integer target, return true if target is in matrix or false otherwise.
// You must write a solution in O(log(m * n)) time complexity.

// SearchMatrix(matrix, target) finds a given element in a matrix using binary search
func SearchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		if matrix[i][0] <= target && target <= matrix[i][m-1] {
			if Search(matrix[i], target) != -1 {
				return true
			}
		}
	}
	return false
}

// implicitely converts the 2d array to id array
func SearchMatrixAlt(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	low, high := 0, (n*m)-1
	// row := index / m(num of elements in that row)
	// col := index % m (to get the right column)
	// elem location(row,col) is thus (i / m, i % m)(0-based indexing)

	//	If matrix[row][col] == target: We should return true here, as we have found the ‘target’.
	//
	// If matrix[row][col] < target: In this case, we need bigger elements. So, we will eliminate the left half and consider the right half (low = mid+1).
	// If matrix[row][col] > target: In this case, we need smaller elements. So, we will eliminate the right half and consider the left half (high = mid-1).
	for low <= high {
		mid := (low + high) / 2
		row := mid / m
		col := mid % m
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
