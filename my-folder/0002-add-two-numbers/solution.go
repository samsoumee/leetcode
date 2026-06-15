/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    return &ListNode{Val: (l1.Val + l2.Val) % 10, Next: solve(l1.Next, l2.Next, (l1.Val + l2.Val) / 10)}
}


func solve(l1 *ListNode, l2 *ListNode, carry int) *ListNode{
    if (l1 == nil && l2 == nil){
        if (carry != 0){
            return &ListNode{Val: carry}
        } else {
            return nil
        }
    } else if (l1 == nil){
        return &ListNode{Val: (l2.Val + carry) % 10, Next: solve(nil, l2.Next, (l2.Val + carry) / 10)}
    } else if (l2 == nil){
        return &ListNode{Val: (l1.Val + carry) % 10, Next: solve(nil, l1.Next, (l1.Val + carry) / 10)}
    } else {
        return &ListNode{Val: (l1.Val + l2.Val + carry) % 10, Next: solve(l1.Next, l2.Next, (l1.Val + l2.Val + carry) / 10)}
    }
}
