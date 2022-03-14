package _99__掉落的方块

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
 @Time    : 2022/3/14 -- 5:20 下午
 @Author  : bishop ❤️ dcm
 @Software: GoLand
 @Description: 699.go
*/

func fallingSquares(positions [][]int) []int {
	qans := make([]int, len(positions))
	for i := 0; i < len(positions); i++ {
		left := positions[i][0]
		size := positions[i][1]
		right := left + size
		qans[i] += size

		for j := i+1; j < len(positions); j++ {
			left2 := positions[j][0]
			size2 := positions[j][1]
			right2 := left2 + size2
			if left2 < right && left < right2 {
				qans[j] = max(qans[j], qans[i])
			}
		}
	}

	var ans []int
	cur := -1
	for _, x := range qans {
		cur = max(cur, x)
		ans = append(ans, cur)
	}
	return ans
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}