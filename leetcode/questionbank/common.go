package questionbank

import (
	"container/heap"
	"fmt"
	"math"
	"sync"
)

type ListNode struct {
	Var  int
	Next *ListNode
}

func (l *ListNode) Print() {
	if l != nil && l.Var != -1 {
		fmt.Print(l.Var, " ")
	}
	if l.Next != nil {
		l.Next.Print()
	}
}

type minHeapForList struct {
	data []*ListNode
}

func (h *minHeapForList) Print() {
	fmt.Println()
	for _, v := range h.data {
		fmt.Print(v.Var, " ")
	}
	fmt.Println()
}

func (h *minHeapForList) Less(i, j int) bool {
	return h.data[i].Var < h.data[j].Var
}

func (h *minHeapForList) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *minHeapForList) Len() int {
	return len(h.data)
}

func (h *minHeapForList) Pop() any {
	tail := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return tail
}

func (h *minHeapForList) Push(item any) {
	h.data = append(h.data, item.(*ListNode))
}

// MergeList
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func MergeList(l1, l2 *ListNode) *ListNode {
	head := &ListNode{-1, nil}
	p := head
	p1, p2 := l1, l2
	for p1 != nil || p2 != nil {
		if p1 == nil {
			p.Next = p2
			break
		}
		if p2 == nil {
			p.Next = p1
			break
		}
		if p1.Var < p2.Var {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	return head
}

// PartitionList
// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
// 你应当 保留 两个分区中每个节点的初始相对位置。
func PartitionList(l *ListNode, target int) *ListNode {
	p1, p2 := &ListNode{-1, nil}, &ListNode{-1, nil}
	l1, l2 := p1, p2
	cur := l
	for cur != nil {
		if cur.Var < target {
			l1.Next = cur
			l1 = l1.Next
		} else {
			l2.Next = cur
			l2 = l2.Next
		}
		cur = cur.Next
	}
	l2.Next = nil
	l1.Next = p2
	return p1
}

// MergeKList
// 给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
func MergeKList(in []*ListNode) *ListNode {
	h := &minHeapForList{}
	heap.Init(h)
	for _, v := range in {
		temp := v
		for temp != nil {
			heap.Push(h, temp)
			temp = temp.Next
		}
	}
	res := &ListNode{-1, nil}
	p := res

	h.Print()
	for h.Len() > 0 {
		v := heap.Pop(h).(*ListNode)
		v.Next = nil
		p.Next = v
		p = p.Next
		h.Print()
	}
	return res
}

// FindKthLastInList
// 单链表的倒数第 k 个节点
func FindKthLastInList(list *ListNode, k int) *ListNode {
	if k <= 0 {
		return nil
	}
	p := list
	kl := list
	counter := 1
	for p != nil {
		p = p.Next
		if counter > k {
			kl = kl.Next
		}
		counter++
	}
	kl.Next = nil
	return kl
}

// FindMidNode
//
// 查找单链表的中点
func FindMidNode(l *ListNode) *ListNode {
	slow, fast := l, l
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// HasCycle
//
// 检查链表中是否存在环
func HasCycle(l *ListNode) bool {
	slow, fast := l, l
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// FindCycleStartNode
//
// 判断是否有环，若有，则返回环的起点
func FindCycleStartNode(l *ListNode) *ListNode {
	slow, fast := l, l
	hasCycle := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			hasCycle = true
			break
		}
	}
	if hasCycle {
		// fast指针从头和slow一样的速率前进
		// 当两者相遇时，即为环的起点
		fast = l
		for fast != slow {
			slow = slow.Next
			fast = fast.Next
		}
		return fast
	}
	return nil
}

// RemoveDuplicates
//
// 删除非严格递增数组中的重复元素
func RemoveDuplicates(nums []int) int {
	p1, p2 := 0, 0
	for p2 <= len(nums)-1 {
		if nums[p1] != nums[p2] {
			p1++
			nums[p1] = nums[p2]
		}
		p2++
	}
	return p1 + 1
}

func FindKSum(a1, a2 []int, k int) []int {
	// get top k
	h := &myHeap{}
	heap.Init(h)
	for i := len(a1) - 1; i >= 0; i-- {
		for j := len(a2) - 1; j >= 0; j-- {
			heap.Push(h, a1[i]+a2[j])
		}
	}
	var result []int
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(h).(int))
	}

	return result
}

type myHeap []int

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *myHeap) Pop() interface{} {
	tail := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tail
}

type Item struct {
	sum int
	i   int
	j   int
}

type MaxHeap []Item

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i].sum > (*h)[j].sum
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i].sum, (*h)[j].sum = (*h)[j].sum, (*h)[i].sum
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Push(v any) {
	*h = append(*h, v.(Item))
}

func (h *MaxHeap) Pop() any {
	tail := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return tail
}

func FindKLargestSums(A, B []int, k int) []int {
	var results []int
	if len(A) == 0 || len(B) == 0 || k <= 0 {
		return results
	}
	m, n := len(A)-1, len(B)-1
	maxHeap := &MaxHeap{
		{
			sum: A[m] + B[n],
			i:   m,
			j:   n,
		},
	}
	heap.Init(maxHeap)
	top := (*maxHeap)[0]
	visited := make(map[[2]int]bool)
	visited[[2]int{m, n}] = true
	for maxHeap.Len() > 0 && k > 0 {
		top = heap.Pop(maxHeap).(Item)
		results = append(results, top.sum)
		k--

		if top.i > 0 && !visited[[2]int{top.i - 1, top.j}] {
			heap.Push(maxHeap, Item{
				sum: A[top.i-1] + B[top.j],
				i:   top.i - 1,
				j:   top.j,
			})
			visited[[2]int{top.i - 1, top.j}] = true
		}
		if top.j > 0 && !visited[[2]int{top.i, top.j - 1}] {
			heap.Push(maxHeap, Item{
				sum: A[top.i] + B[top.j-1],
				i:   top.i,
				j:   top.j - 1,
			})
			visited[[2]int{top.i, top.j - 1}] = true
		}

	}
	return results
}

// SumTarget
//
// 给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。
//
// 如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
//
// 以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
//
// 你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
//
// 你所设计的解决方案必须只使用常量级的额外空间。
func SumTarget(nums []int, target int) []int {
	var result []int
	left, right := 1, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == target {
			result = append(result, left, right)
			return result
		}
		if nums[left]+nums[right] < target {
			left++
			continue
		} else {
			right--
		}
	}
	return result
}

// LongestPalindromeSubStr 找出最长的回文串
func LongestPalindromeSubStr(str string) string {
	var maxStr string
	for i := 0; i < len(str); i++ {
		r1 := findPalindromeStrFromMid(str, i, i)
		if len(maxStr) < len(r1) {
			maxStr = r1
		}
		r2 := findPalindromeStrFromMid(str, i, i+1)
		if len(maxStr) < len(r2) {
			maxStr = r2
		}
	}
	return maxStr
}

func findPalindromeStrFromMid(s string, m, n int) string {
	var res string
	for n < len(s) && m >= 0 {
		if s[m] != s[n] {
			break
		}
		m--
		n++
	}
	res = s[m+1 : n]
	return res
}

// FindMinCoveringSubStr
//
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
// 注意：
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
func FindMinCoveringSubStr(s, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)
	for i, _ := range t {
		need[t[i]]++
	}
	left := 0
	right := 0
	validNums := 0
	var res string
	for right < len(s) {
		a := s[right]
		if _, ok := need[a]; ok {
			window[a]++
			if window[a] == need[a] {
				validNums++
			}
		}
		right++

		// 所有的字符都已经包含
		for validNums == len(need) && left < right {
			if right-left < len(res) || len(res) == 0 {
				res = s[left:right]
			}
			b := s[left]
			left++
			if _, ok := need[b]; ok {
				window[b]--
				if window[b] != need[b] {
					validNums--
				}
			}
		}
	}
	return res
}

func MinWindow(s string, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0
	// 记录最小覆盖子串的起始索引及长度
	start, length := 0, math.MaxInt32
	for right < len(s) {
		// c 是将移入窗口的字符
		c := s[right]
		// 扩大窗口
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			// 在这里更新最小覆盖子串
			if right-left < length {
				start = left
				length = right - left
			}
			// d 是将移出窗口的字符
			d := s[left]
			// 缩小窗口
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	// 返回最小覆盖子串
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

// FindSubStr
//
// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
//
// 换句话说，s1 的排列之一是 s2 的 子串 。
func FindSubStr(s1, s2 string) bool {
	res := FindMinCoveringSubStr(s2, s1)
	return len(res) == len(s1)
}

// MaxProfit
// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
//
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
//
// 示例 1：
//
// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
// 注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//
// 示例 2：
//
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
//
// 提示：
//
// 1 <= prices.length <= 105
//
// 0 <= prices[i] <= 104
func MaxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		}
		profit := v - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

// CoinsChange
//
// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
// 你可以认为每种硬币的数量是无限的。
//
// 示例 1：
//
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
// 示例 2：
//
// 输入：coins = [2], amount = 3
// 输出：-1
// 示例 3：
//
// 输入：coins = [1], amount = 0
// 输出：0
//
// 提示：
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 231 - 1
// 0 <= amount <= 104
func CoinsChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	var res []int
	for _, v := range coins {
		if amount-v > 0 {
			if r := CoinsChange(coins, amount-v); r > 0 {
				res = append(res, r+1)
			}
		}
		if amount-v == 0 {
			res = append(res, 1)
		}
		if amount-v < 0 {
			res = append(res, -1)
		}
	}
	r := math.MaxInt
	for _, v := range res {
		if v < r {
			r = v
		}
	}
	if r == math.MaxInt {
		return -1
	}
	return r
}

// MaxProfit2
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
//
// 返回 你能获得的 最大 利润 。
//
// 示例 1：
//
// 输入：prices = [7,1,5,3,6,4]
// 输出：7
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
// 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3。
// 最大总利润为 4 + 3 = 7 。
// 示例 2：
//
// 输入：prices = [1,2,3,4,5]
// 输出：4
// 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
// 最大总利润为 4 。
//
// 示例 3：
//
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0。
//
// 提示：
//
// 1 <= prices.length <= 3 * 104
//
// 0 <= prices[i] <= 104
func MaxProfit2(prices []int) int {
	totalProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			totalProfit += prices[i] - prices[i-1]
		}
	}

	return totalProfit
}

//方法一：动态规划
//考虑到「不能同时参与多笔交易」，因此每天交易结束后只可能存在手里有一支股票或者没有股票的状态。
//
//定义状态 dp[i][0] 表示第 i 天交易完后手里没有股票的最大利润，dp[i][1] 表示第 i 天交易完后手里持有一支股票的最大利润（i 从 0 开始）。
//
//考虑 dp[i][0] 的转移方程，如果这一天交易完后手里没有股票，那么可能的转移状态为前一天已经没有股票，即 dp[i−1][0]，
//或者前一天结束的时候手里持有一支股票，即 dp[i−1][1]，这时候我们要将其卖出，并获得 prices[i] 的收益。因此为了收益最大化，我们列出如下的转移方程：
//
//dp[i][0]=max{dp[i−1][0],dp[i−1][1]+prices[i]}
//再来考虑 dp[i][1]，按照同样的方式考虑转移状态，那么可能的转移状态为前一天已经持有一支股票，即 dp[i−1][1]，或者前一天结束时还没有股票，即 dp[i−1][0]，这时候我们要将其买入，并减少 prices[i] 的收益。可以列出如下的转移方程：
//
//dp[i][1]=max{dp[i−1][1],dp[i−1][0]−prices[i]}
//对于初始状态，根据状态定义我们可以知道第 0 天交易结束的时候 dp[0][0]=0，dp[0][1]=−prices[0]。
//
//因此，我们只要从前往后依次计算状态即可。由于全部交易结束后，持有股票的收益一定低于不持有股票的收益，因此这时候 dp[n−1][0] 的收益必然是大于 dp[n−1][1] 的，最后的答案即为 dp[n−1][0]。
//
//C++
//Java
//JavaScript
//Golang
//C
//class Solution {
//public:
//int maxProfit(vector<int>& prices) {
//int n = prices.size();
//int dp[n][2];
//dp[0][0] = 0, dp[0][1] = -prices[0];
//for (int i = 1; i < n; ++i) {
//dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i]);
//dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i]);
//}
//return dp[n - 1][0];
//}
//};
//注意到上面的状态转移方程中，每一天的状态只与前一天的状态有关，而与更早的状态都无关，因此我们不必存储这些无关的状态，只需要将 dp[i−1][0] 和 dp[i−1][1] 存放在两个变量中，通过它们计算出 dp[i][0] 和 dp[i][1] 并存回对应的变量，以便于第 i+1 天的状态转移即可。
//
//C++
//Java
//JavaScript
//Golang
//C
//class Solution {
//public:
//int maxProfit(vector<int>& prices) {
//int n = prices.size();
//int dp0 = 0, dp1 = -prices[0];
//for (int i = 1; i < n; ++i) {
//int newDp0 = max(dp0, dp1 + prices[i]);
//int newDp1 = max(dp1, dp0 - prices[i]);
//dp0 = newDp0;
//dp1 = newDp1;
//}
//return dp0;
//}
//};
//复杂度分析
//
//时间复杂度：O(n)，其中 n 为数组的长度。一共有 2n 个状态，每次状态转移的时间复杂度为 O(1)，因此时间复杂度为 O(2n)=O(n)。
//
//空间复杂度：O(n)。我们需要开辟 O(n) 空间存储动态规划中的所有状态。如果使用空间优化，空间复杂度可以优化至 O(1)。
//
//方法二：贪心
//由于股票的购买没有限制，因此整个问题等价于寻找 x 个不相交的区间 (l
//i
//​
//,r
//i
//​
//] 使得如下的等式最大化
//
//i=1
//∑
//x
//​
//a[r
//i
//​
//]−a[l
//i
//​
//]
//其中 l
//i
//​
//表示在第 l
//i
//​
//天买入，r
//i
//​
//表示在第 r
//i
//​
//天卖出。
//
//同时我们注意到对于 (l
//i
//​
//,r
//i
//​
//] 这一个区间贡献的价值 a[r
//i
//​
//]−a[l
//i
//​
//]，其实等价于 (l
//i
//​
//,l
//i
//​
//+1],(l
//i
//​
//+1,l
//i
//​
//+2],…,(r
//i
//​
//−1,r
//i
//​
//] 这若干个区间长度为 1 的区间的价值和，即
//
//a[r
//i
//​
//]−a[l
//i
//​
//]=(a[r
//i
//​
//]−a[r
//i
//​
//−1])+(a[r
//i
//​
//−1]−a[r
//i
//​
//−2])+…+(a[l
//i
//​
//+1]−a[l
//i
//​
//])
//因此问题可以简化为找 x 个长度为 1 的区间 (l
//i
//​
//,l
//i
//​
//+1] 使得 ∑
//i=1
//x
//​
//a[l
//i
//​
//+1]−a[l
//i
//​
//] 价值最大化。
//
//贪心的角度考虑我们每次选择贡献大于 0 的区间即能使得答案最大化，因此最后答案为
//
//ans=
//i=1
//∑
//n−1
//​
//max{0,a[i]−a[i−1]}
//其中 n 为数组的长度。
//
//需要说明的是，贪心算法只能用于计算最大利润，计算的过程并不是实际的交易过程。
//
//考虑题目中的例子 [1,2,3,4,5]，数组的长度 n=5，由于对所有的 1≤i<n 都有 a[i]>a[i−1]，因此答案为
//
//ans=
//i=1
//∑
//n−1
//​
//a[i]−a[i−1]=4
//但是实际的交易过程并不是进行 4 次买入和 4 次卖出，而是在第 1 天买入，第 5 天卖出。
//
//作者：力扣官方题解
//链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/solutions/476791/mai-mai-gu-piao-de-zui-jia-shi-ji-ii-by-leetcode-s/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// run
// 1. 开启M个协程执行M个任务,等待所有任务执行结束后退出(不可以使用waitGroup)
// 2. 开启M个协程执行N个任务,等待所有任务执行结束后退出
func run(num int, job func()) {
	result := make(chan struct{})

	for i := 0; i < num; i++ {
		go func() {
			defer func() {
				result <- struct{}{}
			}()
			job()
		}()
	}

	doneJob := 0
	for range result {
		doneJob++
		if doneJob == num {
			break
		}
	}
}

// run2
// 1. 开启M个协程执行M个任务,等待所有任务执行结束后退出(不可以使用waitGroup)
// 2. 开启M个协程执行N个任务,等待所有任务执行结束后退出
func run2(m, n int, job func()) {
	jobChan := make(chan func(), n)
	wg := sync.WaitGroup{}
	wg.Add(n)
	done := make(chan struct{})

	for k := 0; k < n; k++ {
		jobChan <- job
	}

	for i := 0; i < m; i++ {
		go func() {
			for {
				select {
				case j := <-jobChan:
					j()
					wg.Done()
				case <-done:
					return
				}
			}
		}()
	}

	wg.Wait()
	done <- struct{}{}
}

// run3
// 3个goroutine，分别输出a、b、c
func run3() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		println("a")
	}()
	go func() {
		defer wg.Done()
		println("b")
	}()
	go func() {
		defer wg.Done()
		println("c")
	}()

	wg.Wait()
}

// 3个goroutine，分别输出a、b、c，一共打印10次
// func run4() {
// 	wg := sync.WaitGroup{}
// 	wg.Add(10)
// 	jobs := make(chan int, 10)
// 	done := make(chan struct{})

// 	sendA, sendB, sendC := make(chan struct{}), make(chan struct{}), make(chan struct{})
// 	go func() {
// 		for {
// 			select {
// 			case <-sendA:
// 				_, ok := <-jobs
// 				println("a")
// 				sendB <- struct{}{}
// 			case <-done:
// 				return
// 			}
// 		}
// 	}()
// 	go func() {
// 		for {
// 			_, ok := <-jobs
// 			if ok {
// 				println("b")
// 				wg.Done()
// 			} else {
// 				return
// 			}
// 		}
// 	}()
// 	go func() {
// 		for {
// 			_, ok := <-jobs
// 			if ok {
// 				println("c")
// 				wg.Done()
// 			} else {
// 				return
// 			}
// 		}
// 	}()

// 	for i := 0; i < 10; i++ {
// 		jobs <- i
// 	}

// 	wg.Wait()
// 	close(jobs)
// }
