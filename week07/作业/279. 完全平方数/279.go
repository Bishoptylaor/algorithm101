package _79__完全平方数

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
 @Time    : 2022/2/28 -- 3:38 下午
 @Author  : bishop ❤️ dcm
 @Software: GoLand
 @Description: 279.go
*/
func numSquares(n int) int {
    dp := make([]int,n+1)
    for i:=1; i<=n; i++ {
        minn := math.MaxInt32
        for j:=1; j*j<=i; j++ {
            minn = min(minn, dp[i-j*j])
        }
        dp[i] = minn + 1
    }
    return dp[n]
}

func min(a,b int) int {
    if a<b {
        return a
    }
    return b
}

