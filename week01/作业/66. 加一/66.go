package main

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
 @Time    : 2022/1/10 -- 8:18 下午
 @Author  : bishop ❤️ dcm
 @Software: GoLand
 @Description: 66.go
*/


func plusOne(digits []int) []int {
    length := len(digits)
    carry := 1
    for i:= length-1; i>=0; i-- {
        cur := digits[i]+carry
        carry = cur/10
        digits[i] = cur%10
    }

    if carry>0 {
        digits = append([]int{1}, digits...)
    }
    return digits
}
