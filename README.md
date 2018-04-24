# leetcodeutil
Go util for Leetcode problems

## usage

``` go
import . "github.com/seihmd/leetcodeutil"

func problemTakeTreeNode(t *TreeNode) {}
func problemTakeListNode(l *ListNode) {}
func problemTakeMatrix(m [][]int) {}

func main() {
    problemTakeTreeNode(BinaryTree("[1,2,3,null,5]"))
    problemTakeListNode(LinkedList([]int{1,2,3,4,5}))
    problemTakeMatrix(Matrix("[[1,2,3],[4,5,6],[7,8,9]]"))
}
```

## godoc

https://godoc.org/github.com/seihmd/leetcodeutil
