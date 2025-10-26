package main
import "fmt"

func isValid(s string) bool {

	if len(s) %2 != 0 || len(s) == 0 || s[0] == ')' || s[0] == '}'{
		return  false
	}
	a byte 
	
    

}

func gotestfunc(){
	str := "{}()(){}"
	isValid(str)
}

func main(){
	fmt.Println("hello world")
	gotestfunc()
}
