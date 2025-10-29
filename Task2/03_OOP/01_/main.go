package main

import (
	"fmt"
)

//定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
//考察点 ：接口的定义与实现、面向对象编程风格
var pi float32 = 3.14

type Shape interface {//接口
	Area() //求面机
	Perimeter() //求周长
}

type Rectangle struct {
	x float32
	y float32
}

type Circle struct {
	r float32
}

func (r Rectangle)Area()float32{
	res :=  r.x * r.y
	return res
}

func (c Circle)Area()float32{
	res := float32(c.r) * float32(c.r) * pi
	return res
}

func (r Rectangle)Perimeter()float32{
	res := r.x * r.y
	return res
}

func (c Circle)Perimeter()float32{
	res := 2 * pi * float32(c.r)
	return res	
}

func gotestfunc (){
	r1 := Rectangle{3 , 4}
	areaof_r1 := r1.Area()
	perimeter_of_r1 := r1.Perimeter()
	fmt.Printf("area of r1 is:%v\n", areaof_r1)
	fmt.Printf("perimeter of r1 is:%v\n", perimeter_of_r1)

	c1 := Circle{6}
	areaof_c1 := c1.Area()
	perimeter_of_c1 := c1.Perimeter()
	fmt.Printf("area of c1 is:%v\n", areaof_c1)
	fmt.Printf("perimeter of c1 is:%v\n", perimeter_of_c1)

	
}

func main() {
	fmt.Println("begin test")
	gotestfunc ()
}

//done
