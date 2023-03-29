/**
  @author: jiangxi
  @since: 2022/11/3
  @desc: //TODO
**/
package level

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) (averages []float64) {
	nextLevel := []*TreeNode{root}
	for len(nextLevel) != 0 {
		curLevel := nextLevel
		nextLevel = nil
		sum := 0
		num := 0
		for _, node := range curLevel {
			sum += node.Val
			num += 1
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		averages = append(averages, float64(sum)/float64(num))
	}
	return
}

