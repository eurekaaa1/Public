package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 { //如果是空的直接返回结果
		return ""
	} else if length == 1 { //如果只有一个直接返回第一个作为结果
		return strs[0]
	} else {
		prefix := strs[0]
		for i := 0; i < length; i++ { //将每一个元素都过一遍
			//先假设第一个就是最长公共前缀
			index := 0 //指向公共前缀位置的指标值
			for index < len(strs[i]) && index < len(prefix) && prefix[index] == strs[i][index] {
				index++
			}
			prefix = prefix[:index]
		}
		return prefix
	}
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("begin test")
	ret := longestCommonPrefix(strs)
	fmt.Println(ret)
}