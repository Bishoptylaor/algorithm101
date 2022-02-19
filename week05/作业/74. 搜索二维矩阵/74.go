package _4__搜索二维矩阵

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
 @Time    : 2022/2/14 -- 8:07 下午
 @Author  : bishop ❤️ dcm
 @Software: GoLand
 @Description: 74.go
*/

// 一次二分查找

func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    i := sort.Search(m*n, func(i int) bool { return matrix[i/n][i%n] >= target })
    return i < m*n && matrix[i/n][i%n] == target
}


