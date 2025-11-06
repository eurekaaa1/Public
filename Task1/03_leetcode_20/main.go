package main
import "fmt"

func isValid(s string) bool {

	n := len(s)
	if n %2 != 0 || len(s) == 0 || s[0] == ')' || s[0] == '}' || s[0] == ']'{
		return  false
	}
	
	stack := [] byte {}

	chara_map := map[byte] byte{
		']' : '[',
		'}' : '{',
		')' : '(',
	}

	for i := 0 ; i < n ; i++{
		if chara_map[s[i]] > 0{//字符串的右括号都出栈
			if len(stack) == 0 || chara_map[s[i]] != stack[len(stack) - 1] {
				return false //如果右括号没法匹配栈顶的左括号 就不对
			}
			stack = stack[:len(stack) - 1]//出栈之后栈的个数减少
		}else{
			stack = append(stack, s[i])
		}//左括号入栈
	}//循环整个字符串	
	return (len(stack) == 0)//最后栈如果是空的  那就是对的，否则栈中留有没处理完的左括号就是不通过
	
}

func gotestfunc(){
	//str := "{}()(){}"
	//str2 := "{(])"
	str3 := "(("
	ret := isValid(str3)
	fmt.Println(ret)
}


func main(){
	//fmt.Println("hello world")
	gotestfunc()
}
//done

