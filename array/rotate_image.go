package array

func rotateImage(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			// Save top-left element
			temp := matrix[i][j]

			// Move values from right to top
			matrix[i][j] = matrix[n-j-1][i]

			// Move values from bottom to right
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]

			// Move values from left to bottom
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]

			// Move values from top to left
			matrix[j][n-i-1] = temp
		}
	}
}
