package _5__跳跃游戏

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
 @Time    : 2022/2/28 -- 3:14 下午
 @Author  : bishop ❤️ dcm
 @Software: GoLand
 @Description: 55.go
*/
func canJump(nums []int) bool {
	canReach := 0
	for i:=0; i<len(nums); i++ {
		if i>canReach {
			return false
		}
		canReach = max(canReach, i+nums[i])
	}
	return true
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

