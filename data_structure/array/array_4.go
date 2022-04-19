package array

/**
  @Author: pika
  @Date: 2022/4/16 3:52 下午
  @Description: 二维数组
*/

/*
	给定一个 n×n 的二维矩阵matrix 表示一个图像。请你将图像顺时针旋转 90 度。
	你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

*/

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < len(matrix); i++ {
		reverseRow(matrix[i])
	}

}

func reverseRow(row []int) {
	l, r := 0, len(row)-1
	for l < r {
		row[l], row[r] = row[r], row[l]
		l++
		r--
	}
}

/*
	给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
*/

//失败
func spiralOrder1(matrix [][]int) []int {
	lenS := len(matrix) * len(matrix[0])
	i, j := 0, 0
	// up上边界、lb左边界、db下边界、rb右边界
	ub, lb, db, rb := 0, 0, len(matrix)-1, len(matrix[0])-1
	var res []int
	for len(res) < lenS {
		// 从左到右
		for i == ub && j < rb {
			res = append(res, matrix[i][j])
			j++
		}
		ub++
		//从上到下
		for j == rb && i < db {
			res = append(res, matrix[i][j])
			i++
		}
		rb--

		//从右到左
		for i == db && j > lb {
			res = append(res, matrix[i][j])
			j--
		}
		db--
		//从下到上
		for j == lb && i > ub {
			res = append(res, matrix[i][j])
			i--
		}
		lb++
	}
	return res
}

//失败
func spiralOrder2(matrix [][]int) []int {
	lenS := len(matrix) * len(matrix[0])
	// up上边界、lb左边界、db下边界、rb右边界
	ub, lb, db, rb := 0, 0, len(matrix)-1, len(matrix[0])-1
	var res []int
	for len(res) < lenS {
		// 从左到右
		if ub < db {
			for j := 0; j < rb; j++ {
				res = append(res, matrix[ub][j])
			}
			// 上边界下移1
			ub++
		}
		//从上到下
		if rb > lb {
			for i := 0; i < db; i++ {
				res = append(res, matrix[i][rb])
			}
			// 右边界左移1
			rb--
		}

		//从右到左
		if db > ub {
			for j := rb; j > lb; j-- {
				res = append(res, matrix[db][j])
			}
			//下边界上移1
			db--
		}

		//从下到上
		if lb < rb {
			for i := db; i > ub; i-- {
				res = append(res, matrix[i][lb])
			}
			//左边界右移1
			lb++
		}

	}
	return res
}

//失败
func spiralOrder3(matrix [][]int) []int {
	// up上边界、lb左边界、db下边界、rb右边界
	ub, lb, db, rb := 0, 0, len(matrix)-1, len(matrix[0])-1
	var res []int
	for ub <= db && lb <= rb {
		// 从左到右
		for j := 0; j < rb; j++ {
			res = append(res, matrix[ub][j])
		}

		//从上到下
		for i := 0; i < db; i++ {
			res = append(res, matrix[i][rb])
		}

		//从右到左
		for j := rb; j > lb; j-- {
			res = append(res, matrix[db][j])
		}

		//从下到上
		for i := db; i > ub; i-- {
			res = append(res, matrix[i][lb])
		}
		// 上边界下移1
		ub++
		// 右边界左移1
		rb--
		//下边界上移1
		lb++
		//左边界右移1
		db--
	}
	return res
}

func spiralOrder(matrix [][]int) []int {
	// up上边界、lb左边界、db下边界、rb右边界
	ub, lb, db, rb := 0, 0, len(matrix)-1, len(matrix[0])-1
	var res []int
	for ub <= db && lb <= rb {
		// 从左到右
		for j := lb; j <= rb; j++ {
			res = append(res, matrix[ub][j])
		}

		//从上到下
		for i := ub + 1; i <= db; i++ {
			res = append(res, matrix[i][rb])
		}

		if ub < db && lb < rb {
			//从右到左
			for j := rb - 1; j > lb; j-- {
				res = append(res, matrix[db][j])
			}

			//从下到上
			for i := db; i > ub; i-- {
				res = append(res, matrix[i][lb])
			}
		}

		// 上边界下移1
		ub++
		// 右边界左移1
		rb--
		//下边界上移1
		lb++
		//左边界右移1
		db--
	}
	return res
}

/*
	给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
*/

func generateMatrix(n int) [][]int {
	if n <= 0 {
		return nil
	}
	// 初始化返回二维数组
	matrix := make([][]int, n)
	for i, _ := range matrix {
		matrix[i] = make([]int, n)
	}

	ub, db, lb, rb := 0, n-1, 0, n-1
	num := 1
	for num <= n*n && ub <= db && lb <= rb {
		// 向右
		for col := lb; col <= rb; col++ {
			matrix[ub][col] = num
			num++
		}

		// 向下
		for row := ub + 1; row <= db; row++ {
			matrix[row][rb] = num
			num++
		}
		if ub < db && lb < rb {
			// 向左
			for col := rb - 1; col > lb; col-- {
				matrix[db][col] = num
				num++
			}

			// 向上
			for row := db; row > ub; row-- {
				matrix[row][lb] = num
				num++
			}
		}

		ub++
		rb--
		db--
		lb++

	}
	return matrix

}
