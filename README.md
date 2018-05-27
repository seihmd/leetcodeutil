# leetcodeutil
Go util for Leetcode problems

## usage

``` go
import . "github.com/seihmd/leetcodeutil"

// solve problems!
func problemTakesTreeNode(t *TreeNode) {}
func problemTakesListNode(l *ListNode) {}
func problemTakesMatrix(m [][]int) {}
func problemTakesIntervals(i []Interval) {}

func main() {
    problemTakesTreeNode(BinaryTree("[1,2,3,null,5]"))
    problemTakesListNode(LinkedList("[1,2,3,4,5]"))
    problemTakesMatrix(Matrix("[[1,2,3],[4,5,6],[7,8,9]]"))
    problemTakesIntervals(Intervals("[[1,2],[5,6],[8,9]]"))
}
```

## godoc

https://godoc.org/github.com/seihmd/leetcodeutil
