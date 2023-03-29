/**
  @author: jiangxi
  @since: 2022/11/3
  @desc: //TODO
**/
package level

func findBottomLeftValue(root *TreeNode) int {
	nextLevel := []*TreeNode{root}
	var left int
	for len(nextLevel) != 0 {
		curLevel := nextLevel
		nextLevel = nil
		left = curLevel[0].Val
		for _, node := range curLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
	}
	return left
}
