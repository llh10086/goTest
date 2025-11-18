package main

import (
	"fmt"
	"sort"
)

func main() {
	//只出现一次的字符
	arr := []string{"a", "a", "a", "b", "b", "b", "c"}
	commandParser(arr)
	//有效的括号
	var s = "([])"
	fmt.Println(isValid(s))

	//最长公链前缀
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

	//基本值类型
	digits := []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits))

	//切片去重
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5}
	newLength := removeDuplicates(nums)
	fmt.Printf("去重后的切片长度为: %d\n", newLength)
	fmt.Printf("去重后的切片内容为: %v\n", nums[:newLength])

	//合并区间
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	mergedIntervals := merge(intervals)
	fmt.Printf("合并后的区间为: %v\n", mergedIntervals)

	// 两数之和
	nums2 := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums2, target)
	fmt.Printf("两数之和的索引为: %v\n", result)

}

// 只出现一次的数字
func commandParser(arr []string) map[string]int {
	map1 := make(map[string]int)
	map2 := make(map[string]int)
	for _, v := range arr {
		if m, ok := map1[v]; ok {
			map1[v] = m + 1
		} else {
			map1[v] = 1
		}
	}
	for k, v := range map1 {
		if v == 1 {
			fmt.Printf("切片中出现次数为1的元素为%v\n", k)
			map2[k] = v
		}
	}
	return map2
}

// 有效的括号
func isValid(s string) bool {
	// 创建一个映射，存储右括号对应的左括号
	mapping := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 用切片实现栈的功能
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		char := s[i]
		// 如果是右括号
		if val, ok := mapping[char]; ok {
			// 栈为空或栈顶元素不匹配，返回false
			if len(stack) == 0 || stack[len(stack)-1] != val {
				return false
			}
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，压入栈中
			stack = append(stack, char)
		}
	}

	// 最后栈必须为空才是有效的
	return len(stack) == 0
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	// 处理空输入的情况
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串为基准，遍历每个字符位置
	for i := 0; i < len(strs[0]); i++ {
		// 取出当前位置的字符
		currentChar := strs[0][i]

		// 与其他字符串的相同位置进行比较
		for j := 1; j < len(strs); j++ {
			// 如果超出其他字符串长度或字符不匹配，返回当前前缀
			if i >= len(strs[j]) || strs[j][i] != currentChar {
				return strs[0][:i]
			}
		}
	}

	// 如果第一个字符串是所有字符串的前缀
	return strs[0]
}

// 基本值类型
func plusOne(digits []int) []int {
	n := len(digits)

	// 从最后一位开始处理
	for i := n - 1; i >= 0; i-- {
		// 当前位加1
		digits[i]++
		// 取模10，若结果不为0，说明无进位，直接返回
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
		// 若结果为0，说明有进位，继续处理前一位
	}

	// 若所有位都有进位（如999 -> 1000），需要增加数组长度
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

// 切片去重
func removeDuplicates(nums []int) int {
	// 处理空数组的情况
	if len(nums) == 0 {
		return 0
	}

	// 慢指针i：记录不重复元素的位置（初始指向第一个元素）
	i := 0

	// 快指针j：遍历整个数组（从第二个元素开始）
	for j := 1; j < len(nums); j++ {
		// 当快指针元素与慢指针元素不同时
		if nums[j] != nums[i] {
			// 慢指针后移一位，将快指针元素赋值到新位置
			i++
			nums[i] = nums[j]
		}
		// 若元素相同则快指针继续后移，不做操作
	}

	// 新长度为慢指针索引+1（因为索引从0开始）
	return i + 1
}

// 合并区间
func merge(intervals [][]int) [][]int {
	// 处理空输入
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按照区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果切片，放入第一个区间
	result := [][]int{intervals[0]}

	// 遍历剩余区间
	for i := 1; i < len(intervals); i++ {
		// 获取结果中最后一个区间
		last := result[len(result)-1]
		// 当前区间的起始和结束
		currentStart, currentEnd := intervals[i][0], intervals[i][1]

		// 如果当前区间的起始 <= 结果中最后一个区间的结束，说明有重叠，需要合并
		if currentStart <= last[1] {
			// 合并后的结束取两个区间的最大值
			if currentEnd > last[1] {
				last[1] = currentEnd
			}
		} else {
			// 无重叠，直接加入结果
			result = append(result, intervals[i])
		}
	}

	return result
}

// 两数之和
func twoSum(nums []int, target int) []int {
	// 创建一个哈希表，用于存储数值到索引的映射
	numMap := make(map[int]int)

	// 遍历数组
	for i, num := range nums {
		// 计算需要寻找的互补数
		complement := target - num

		// 检查互补数是否在哈希表中
		if j, exists := numMap[complement]; exists {
			// 找到则返回两个数的索引
			return []int{j, i}
		}

		// 没找到则将当前数和索引存入哈希表
		numMap[num] = i
	}

	// 题目保证有解，这里只是为了满足函数返回要求
	return []int{}
}
