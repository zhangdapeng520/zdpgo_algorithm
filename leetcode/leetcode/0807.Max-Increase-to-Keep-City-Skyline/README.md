# [807. Max Increase to Keep City Skyline](https://leetcode.com/problems/max-increase-to-keep-city-skyline/)

## 题目

There is a city composed of n x n blocks, where each block contains a single building shaped like a vertical square prism. You are given a 0-indexed n x n integer matrix grid where grid[r][c] represents the height of the building located in the block at row r and column c.

A city's skyline is the the outer contour formed by all the building when viewing the side of the city from a distance. The skyline from each cardinal direction north, east, south, and west may be different.

We are allowed to increase the height of any number of buildings by any amount (the amount can be different per building). The height of a 0-height building can also be increased. However, increasing the height of a building should not affect the city's skyline from any cardinal direction.

Return the maximum total sum that the height of the buildings can be increased by without changing the city's skyline from any cardinal direction.

**Example 1**:

![https://assets.leetcode.com/uploads/2021/06/21/807-ex1.png](https://assets.leetcode.com/uploads/2021/06/21/807-ex1.png)

    Input: grid = [[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]
    Output: 35
    Explanation: The building heights are shown in the center of the above image.
    The skylines when viewed from each cardinal direction are drawn in red.
    The grid after increasing the height of buildings without affecting skylines is:
    gridNew = [ [8, 4, 8, 7],
    [7, 4, 7, 7],
    [9, 4, 8, 7],
    [3, 3, 3, 3] ]

**Example 2**:

    Input: grid = [[0,0,0],[0,0,0],[0,0,0]]
    Output: 0
    Explanation: Increasing the height of any building will result in the skyline changing.

**Constraints:**

- n == grid.length
- n == grid[r].length
- 2 <= n <= 50
- 0 <= grid[r][c] <= 100

## 题目大意

在二维数组grid中，grid[i][j]代表位于某处的建筑物的高度。 我们被允许增加任何数量（不同建筑物的数量可能不同）的建筑物的高度。 高度 0 也被认为是建筑物。

最后，从新数组的所有四个方向（即顶部，底部，左侧和右侧）观看的“天际线”必须与原始数组的天际线相同。 城市的天际线是从远处观看时，由所有建筑物形成的矩形的外部轮廓。 请看下面的例子。

建筑物高度可以增加的最大总和是多少？

## 解题思路

- 从数组竖直方向（即顶部，底部）看“天际线”计算出 topBottomSkyline
- 从数组水平方向（即左侧，右侧）看“天际线”计算出 leftRightSkyline
- 计算 grid 中每个元素与对应的 topBottomSkyline 和 leftRightSkyline 中较小值的差值
- 统计所有差值的总和 ans 并返回

## 代码

```go
package leetcode

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	topBottomSkyline := make([]int, 0, n)
	leftRightSkyline := make([]int, 0, n)
	for i := range grid {
		cur := 0
		for _, v := range grid[i] {
			if cur < v {
				cur = v
			}
		}
		leftRightSkyline = append(leftRightSkyline, cur)
	}
	for j := range grid {
		cur := 0
		for i := 0; i < len(grid[0]); i++ {
			if cur < grid[i][j] {
				cur = grid[i][j]
			}
		}
		topBottomSkyline = append(topBottomSkyline, cur)
	}
	var ans int
	for i := range grid {
		for j := 0; j < len(grid[0]); j++ {
			ans += min(topBottomSkyline[j], leftRightSkyline[i]) - grid[i][j]
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```