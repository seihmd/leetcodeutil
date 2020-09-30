# leetcodeutil
Go util for [Leetcode](https://leetcode.com/) problems

## usage

``` go
import . "github.com/seihmd/leetcodeutil"

func problemTakesTreeNode(t *TreeNode) {}
func problemTakesListNode(l *ListNode) {}
func problemTakesMatrix(m [][]int) {}
func problemTakesIntervals(i []Interval) {}
func problemTakesIntSlice(i []int) {}
func problemTakesIntSlices(i [][]int) {}

// solve problems!
func Solve() {
	problemTakesTreeNode(BinaryTree("[1,2,3,null,5]"))
	problemTakesListNode(LinkedList("[1,2,3,4,5]"))
	problemTakesMatrix(Matrix("[[1,2,3],[4,5,6],[7,8,9]]"))
	problemTakesIntervals(Intervals("[[1,2],[5,6],[8,9]]"))
	problemTakesIntSlice(IntSlice("[1,2,3]"))
	problemTakesIntSlices(IntSlices("[[1],[2,3],[4,5,6]]"))
}
```

## godoc

https://godoc.org/github.com/seihmd/leetcodeutil
