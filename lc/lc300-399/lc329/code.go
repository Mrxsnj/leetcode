var (
    cache [][]int
    dirs = [][]int{[]int{0, -1}, []int{0, 1}, []int{-1, 0}, []int{1, 0}}
)

func longestIncreasingPath(matrix [][]int) int {
    ans := 1
    row := len(matrix)
    column := len(matrix[0])
    cache = make([][]int, row)

    for i := 0; i < row; i++ {
        cache[i] = make([]int, column)
    }

    for m, arr := range matrix {
        for n, _ := range arr {
            ans = max(ans, dfs(matrix, m, n))
        }
    }

    return ans
}

func dfs(matrix [][]int, row, column int) int {
    if cache[row][column] != 0 {
        return cache[row][column]
    }

    cache[row][column]++
    delta := 0

    for _, dir := range dirs {
        newRow := row + dir[0]
        newColumn := column + dir[1]

        if newRow >= 0 && newRow < len(matrix) && newColumn >= 0 && newColumn < len(matrix[0]) && matrix[newRow][newColumn] > matrix[row][column] {
            delta = max(delta, dfs(matrix, newRow, newColumn))
        }
    }

    cache[row][column] += delta

    return cache[row][column]
}

func max(x, y int) int {
    if x > y {
        return x
    }

    return y
}
