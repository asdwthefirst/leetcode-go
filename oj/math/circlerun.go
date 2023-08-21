/*
*

	@author: jiangxi
	@since: 2023/7/23
	@desc: //TODO

*
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var L, R float64
	// fmt.Println("请输入跑步总路程 L 和跑道半径 R：")
	fmt.Scan(&L, &R)

	// 顺时针情况下结束位置的坐标
	xCounterClockwise := R * math.Cos(-L/R)
	yCounterClockwise := R * math.Sin(-L/R)
	fmt.Printf("%.3f %.3f\n", xCounterClockwise, yCounterClockwise)

	// 逆时针情况下结束位置的坐标
	xClockwise := R * math.Cos(L/R)
	yClockwise := R * math.Sin(L/R)
	fmt.Printf("%.3f %.3f\n", xClockwise, yClockwise)
}

//圆的极坐标方程可以表示为：r = R
//假设小明已经跑了L的总路程，那么他的极坐标方程可以表示为：θ = L/R

//我们可以使用以下公式将极坐标转换为直角坐标：
//x = r * cos(θ)
//y = r * sin(θ)
