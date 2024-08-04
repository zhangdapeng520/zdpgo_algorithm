# [1518. Water Bottles](https://leetcode.com/problems/water-bottles/)

## 题目

Given numBottles full water bottles, you can exchange numExchange empty water bottles for one full water bottle.

The operation of drinking a full water bottle turns it into an empty bottle.

Return the maximum number of water bottles you can drink.

**Example 1**:

![https://assets.leetcode.com/uploads/2020/07/01/sample_1_1875.png](https://assets.leetcode.com/uploads/2020/07/01/sample_1_1875.png)

    Input: numBottles = 9, numExchange = 3
    Output: 13
    Explanation: You can exchange 3 empty bottles to get 1 full water bottle.
    Number of water bottles you can drink: 9 + 3 + 1 = 13.

**Example 2**:

![https://assets.leetcode.com/uploads/2020/07/01/sample_2_1875.png](https://assets.leetcode.com/uploads/2020/07/01/sample_2_1875.png)

    Input: numBottles = 15, numExchange = 4
    Output: 19
    Explanation: You can exchange 4 empty bottles to get 1 full water bottle.
    Number of water bottles you can drink: 15 + 3 + 1 = 19.

**Example 3**:

    Input: numBottles = 5, numExchange = 5
    Output: 6

**Example 4**:

    Input: numBottles = 2, numExchange = 3
    Output: 2

**Constraints:**

- 1 <= numBottles <= 100
- 2 <= numExchange <= 100

## 题目大意

小区便利店正在促销，用 numExchange 个空酒瓶可以兑换一瓶新酒。你购入了 numBottles 瓶酒。

如果喝掉了酒瓶中的酒，那么酒瓶就会变成空的。

请你计算 最多 能喝到多少瓶酒。

## 解题思路

- 模拟。首先我们一定可以喝到 numBottles 瓶酒，剩下 numBottles 个空瓶。接下来我们可以拿空瓶子换酒，每次拿出 numExchange 个瓶子换一瓶酒，然后再喝完这瓶酒，得到一个空瓶。这样模拟下去，直到所有的空瓶子小于numExchange结束。

## 代码

```go
package leetcode

func numWaterBottles(numBottles int, numExchange int) int {
	if numBottles < numExchange {
		return numBottles
	}
	quotient := numBottles / numExchange
	reminder := numBottles % numExchange
	ans := numBottles + quotient
	for quotient+reminder >= numExchange {
		quotient, reminder = (quotient+reminder)/numExchange, (quotient+reminder)%numExchange
		ans += quotient
	}
	return ans
}
```