package main

//import "fmt"
import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length == 0{
		return nil//切片返回空值
	}

	// 按区间起始位置升序排序
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    // 初始化结果列表，加入第一个区间
    result := [][]int{intervals[0]}
    
    // 遍历剩余区间并合并
    for i := 1; i < len(intervals); i++ {
        last := result[len(result)-1] // 结果列表中最后一个区间
        current := intervals[i]       // 当前区间
        
        if current[0] <= last[1] {
            // 重叠，合并区间（更新结束位置为较大值）
            if current[1] > last[1] {
                last[1] = current[1]
            }
        } else {
            // 不重叠，直接加入结果
            result = append(result, current)
        }
    }
    
    return result
}

func main(){
	num := [][]int {{1,4} , {1,5}}
	ret := merge(num)
	fmt.Println(ret)
}