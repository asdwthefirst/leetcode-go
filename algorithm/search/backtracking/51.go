/**
  @author: jiangxi
  @since: 2022/12/30
  @desc: //TODO
**/
package backtracking

import "fmt"

func solveNQueens(n int) (ret [][]string) {
	pad := make([][]bool, n)
	for i := range pad {
		pad[i] = make([]bool, n)
	}
	var backtrack51 func(n int, x, y int)
	var canStand func(int, int) bool
	backtrack51 = func(remain int, x, y int) {
		pad[x][y] = true
		//fmt.Println("stand:", x, y)
		//fmt.Println(pad)
		if remain == 1 {
			ret = append(ret, padToStrArr(pad))
		} else {
			for j := 0; j < n; j++ {
				//fmt.Println(j)
				// fmt.Println(pad, x, y, x+1, j)
				//fmt.Println(x+1, j)
				//fmt.Println(&x, &j)
				if canStand(x+1, j) {
					fmt.Println("canstand:", x+1, j)
					backtrack51(remain-1, x+1, j)
				}
			}
		}
		pad[x][y] = false
	}
	//pad = [][]bool{{false, true, false, false}, {false, false, false, true}, {true, false, false, false}, {false, false, false, false}}
	canStand = func(x, y int) bool {
		for i := 0; i <= x; i++ {
			if pad[i][y] {
				return false
			}
		}
		for i := 1; i <= n-1; i++ {
			if x-i >= 0 && y-i >= 0 && pad[x-i][y-i] {
				return false
			}
			if x+i < n && y+i < n && pad[x+i][y+i] {

				return false
			}
			if x-i >= 0 && y+i < n && pad[x-i][y+i] {
				return false
			}
			if x+i < n && y-i >= 0 && pad[x+i][y-i] {
				return false
			}
		}
		return true
	}
	//fmt.Println(canStand(3, 0), canStand(3, 1), canStand(3, 2), canStand(3, 3))

	for j := 0; j < n; j++ {
		backtrack51(n, 0, j)
	}
	return

}

func padToStrArr(pad [][]bool) (ret []string) {
	n := len(pad)
	for i := 0; i < n; i++ {
		str := ""
		for j := 0; j < n; j++ {
			if pad[i][j] {
				str += "Q"
			} else {
				str += "."
			}
		}
		ret = append(ret, str)
	}
	return
}
