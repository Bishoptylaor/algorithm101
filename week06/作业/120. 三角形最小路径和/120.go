package _20__三角形最小路径和

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
 @Time    : 2022/2/28 -- 2:56 下午
 @Author  : bishop ❤️ dcm
 @Software: GoLand
 @Description: 120.go
*/

func findNumberOfLIS(nums []int) (ans int) {
    maxLen := 0
    n := len(nums)
    dp := make([]int, n)
    cnt := make([]int, n)
    for i, x := range nums {
        dp[i] = 1
        cnt[i] = 1
        for j, y := range nums[:i] {
            if x > y {
                if dp[j]+1 > dp[i] {
                    dp[i] = dp[j] + 1
                    cnt[i] = cnt[j] // 重置计数
                } else if dp[j]+1 == dp[i] {
                    cnt[i] += cnt[j]
                }
            }
        }
        if dp[i] > maxLen {
            maxLen = dp[i]
            ans = cnt[i] // 重置计数
        } else if dp[i] == maxLen {
            ans += cnt[i]
        }
    }
    return
}