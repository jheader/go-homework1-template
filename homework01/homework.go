package homework01

import "errors"

// 出现一次的数字
// 给定一个非空数数组，除了某个元素只出现一次以外，其他每个元素均出现二次
//出现两次的元素会相互抵消（结果为 0）；
//最终剩下的结果就是只出现一次的元素（0 和该元素异或仍为它本身）

func SingleNumber(nums []int) int {

	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverData := 0

	if reverData < x {
		//后面添加一位，x%10 取余获取最后一位
		reverData = reverData*10 + x%10
		//去掉最后一位
		x = x / 10
	}

	return x == reverData || x == reverData/10
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		char := s[i]
		if stack.IsEmpty() {
			stack.Push(char)
		} else {
			topVal, _ := stack.Top()
			if topVal == char {
				stack.Pop()
			} else {
				stack.Push(char)
			}

		}
	}
	return stack.IsEmpty()
}

type Stack struct {
	elements []interface{}
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
	}
}

func (s *Stack) Push(element interface{}) {

	s.elements = append(s.elements, element)

}

func (s *Stack) Pop() (interface{}, error) {

	if s.IsEmpty() {
		return nil, errors.New("栈为空，无栈顶元素")
	}

	// 获取栈顶元素
	topIdx := len(s.elements) - 1
	topVal := s.elements[topIdx]
	// 移除栈顶元素（切片截取，不修改原切片，仅重新赋值）
	s.elements = s.elements[:topIdx]
	return topVal, nil

}

// Top 取栈顶元素：返回栈顶元素（不删除，栈空时返回错误）
func (s *Stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("栈为空，无栈顶元素")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Size() int {

	return len(s.elements)
}

func (s *Stack) Clear() {
	s.elements = make([]interface{}, 0)
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	prefix := strs[0]
	for _, s := range strs[1:] {

		for len(s) < len(prefix) || s[:len(prefix)] != prefix {
			// 缩短前缀（去掉最后一个字符）
			prefix = prefix[:len(prefix)-1]
			// 前缀为空，提前终止
			if prefix == "" {
				return ""
			}
		}

	}

	return prefix
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	return append([]int{1}, digits...)
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// TODO: implement
	return 0
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// TODO: implement
	return nil
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {

	numMap := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {

		comp := target - nums[i]
		if idx, exists := numMap[comp]; exists {

			return []int{idx, nums[i]}
		}
		numMap[nums[i]] = nums[i]
	}

	return nil
}
