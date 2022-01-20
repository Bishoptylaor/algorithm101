package _30__被围绕的区域

/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2022/1/20 -- 11:50 上午
 @Author  : bishop ❤️ dcm
 @Software: GoLand
 @Description: 130.go
*/


func solve(board [][]byte)  {
    if len(board) == 0 || len(board[0]) == 0 {
        return
    }
    row, col := len(board), len(board[0])
	gd := InitGridDfs(row, col)
    for i := 0; i < row; i++ {
	   gd.Dfs(board, i, 0)
	   gd.Dfs(board, i, col-1)
    }
	for j := 0; j < col; j++ {
	   gd.Dfs(board, 0, j)
	   gd.Dfs(board, row-1, j)
	}

	for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if board[i][j] == 'A' {
                board[i][j] = 'O'
            } else if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
        }
    }
}

var GridDfs gridDfs

type gridDfs struct {
	Used [][]bool
}

func InitGridDfs(row, col int) *gridDfs {
	used := make([][]bool, row)
	for g := range used {
		used[g] = make([]bool, col)
	}
	return &gridDfs{Used: used}
}

// Dfs dfs core
func (g *gridDfs)Dfs(grid [][]byte, row, col int) {
	if !g.InArea(grid, row, col) {
		return
	}
	if !g.BlockValid(grid, row, col) {
		return
	}
	// if g.IsPaint(row, col) {
	// 	return
	// }
	// g.Paint(row, col)
	grid[row][col] = 'A'

	g.Dfs(grid, row+1, col)
	g.Dfs(grid, row-1, col)
	g.Dfs(grid, row, col+1)
	g.Dfs(grid, row, col-1)
}

// InArea 判断是否在区域内
func (g *gridDfs)InArea(grid [][]byte, row, col int) bool {
	return 0 <= row && row < len(grid) &&
		0 <= col && col < len(grid[0])
}

// Paint 标记已访问
func (g *gridDfs)Paint(row, col int) {
	g.Used[row][col] = true
}

// UnPaint 清除标记
func (g *gridDfs)UnPaint(row, col int) {
	g.Used[row][col] = false
}

// IsPaint 检查已访问
func (g *gridDfs)IsPaint(row, col int) bool {
	return g.Used[row][col]
}

// BlockValid 有效地块，需要按照题目定制
func (g *gridDfs)BlockValid(grid [][]byte, row, col int) bool {
	return grid[row][col] == 'O'
}