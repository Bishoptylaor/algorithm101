package _00__岛屿数量

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
 @Time    : 2022/3/14 -- 5:39 下午
 @Author  : bishop ❤️ dcm
 @Software: GoLand
 @Description: 200.go
*/
func numIslands(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])
	gd := InitGridDfs(row, col)

	var num int
	for i:= 0; i<row; i++ {
		for j:=0; j<col; j++ {
			if !gd.IsPaint(i, j) {
				if grid[i][j] == '1' {
					gd.Dfs(grid, i, j)
					num += 1
				}
			}
		}
	}

	return num
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
	if g.IsPaint(row, col) {
		return
	}
	g.Paint(row, col)

	g.Dfs(grid, row+1, col)
	g.Dfs(grid, row-1, col)
	g.Dfs(grid, row, col+1)
	g.Dfs(grid, row, col-1)
}

// InArea 是否在格子内
func (g *gridDfs)InArea(grid [][]byte, row, col int) bool {
	return 0 <= row && row < len(grid) &&
		0 <= col && col < len(grid[0])
}

// Paint 标记已访问
func (g *gridDfs)Paint(row, col int) {
	g.Used[row][col] = true
}

// UnPaint 标记已访问
func (g *gridDfs)UnPaint(row, col int) {
	g.Used[row][col] = false
}

// IsPaint 标记已访问
func (g *gridDfs)IsPaint(row, col int) bool {
	return g.Used[row][col]
}

// BlockValid 有效地块，需要按照题目定制
func (g *gridDfs)BlockValid(grid [][]byte, row, col int) bool {
	return grid[row][col] == '1'
}