package main

func maxMatrixSum(matrix [][]int) int64 {
	minn, sum := 9999999999, 0
	cnt := 0
	cnt0 := 0
  for i := range matrix {
		for j := range matrix[i] {
			sum += abs(matrix[i][j])
			if matrix[i][j] == 0 {
				cnt0++
			}
			if matrix[i][j] < 0 {
				cnt++
			}
			minn = min(minn, abs(matrix[i][j]))
		}
	}
	if cnt % 2 == 0 || cnt0 > 0{
		return int64(sum)
	}
	return int64(sum - 2 * minn)
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
