package _75__爱吃香蕉的珂珂

import "sort"

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
 @Time    : 2022/2/14 -- 8:19 下午
 @Author  : bishop ❤️ dcm
 @Software: GoLand
 @Description: 875.go
*/

func minEatingSpeed(piles []int, h int) int {
	// 最小速度
	low := 1
	sort.Ints(piles)
	// 最大速度 = 一次吃完最大一堆
	high := piles[len(piles)-1]

	for low < high {
		mid := low + (high - low) / 2
		if getTime(piles, mid) > h {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return high
}

// 以速度speed吃光所有堆的时间，注意向上取整
func getTime(piles []int, speed int) int {
	time := 0
	for _, p := range piles {
		// 向上取整
		time += (p-1)/speed + 1
	}
	return time
}
