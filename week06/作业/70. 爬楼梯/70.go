package _0__爬楼梯

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
 @Time    : 2022/2/28 -- 2:55 下午
 @Author  : bishop ❤️ dcm
 @Software: GoLand
 @Description: 70.go
*/

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n + 1)
	dp[1] = 1
	dp[2] = 2
	for s := 3; s <= n; s++ {
		dp[s] = dp[s-1] + dp[s-2]
	}
	return dp[n]
}