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
 @Time    : 2022/1/10 -- 8:53 下午
 @Author  : bishop ❤️ dcm
 @Software: GoLand
 @Description: 21.go
https://leetcode-cn.com/problems/merge-two-sorted-lists/
*/

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    cur := head
    for l1 != nil || l2 != nil {
        if l1 != nil && l2 != nil {
            if l1.Val < l2.Val {
                cur.Next = l1
                l1 = l1.Next
            } else {
                cur.Next = l2
                l2 = l2.Next
            }
            cur = cur.Next
        } else if l1 != nil {
            cur.Next = l1
            break
        } else {
            cur.Next = l2
            break
        }
    }
    return head.Next
}