package main

import (
"fmt"

//"unsafe"
)

func main1(){

	str := `
	aaaaaaaaaaaaaaa
	bbbbbbbbbbbbbbb
	ccccccccccccccc
	ddddddddddddddd
	fffffffffffffff
	`
	var int01 int =50
	var double float64 = float64(int01)
	var string01 string
	string01 = fmt.Sprintf("%d",int01)
	fmt.Println(string01)
	fmt.Println(string01)
	fmt.Println(double)
	fmt.Println(str)
	fmt.Println("hello ..")

	var num int =20
	var num1 *int
	num1 = &num
	*num1 = 40

	var nbc int =20
	var nbb *int = &nbc
	*nbb =90
	fmt.Println(*nbb)
	fmt.Println(*num1)


	a := 20
	b := 10
	a = a+b
	b = a-b
	a = a-b
	fmt.Println(a,b)
	if a < b{
		fmt.Println("a<b")
	}else{
		fmt.Println("b>a")
	}

	if c:=20;c<18{
		fmt.Println("c")
	}else{
		fmt.Println("c不等于18")
	}

	injj := 20
	incc := 30
	inbb := 40
	if indd:=30; indd< injj{

	}else if incc > inbb{

	}else if inbb < injj{

	}else{
		fmt.Println("条件都不成立")
	}

	var number int32 = 20
	var number1 int32 = 30
	if (number+number1)>=50{
		fmt.Println("大于等于50")
	}
	var double1 float64 = 20.0
	var double2 float64 = 10.0

	if double1>10 && double2<20{
		fmt.Println(" double1>10   double2<20")
	}

	var able int32 = 30
	var able1 int32 = 60

	if (able+able1)%3==0&&(able+able1)%5==0{
		fmt.Println("可以被整除")
	}

	var bae int16 = 2016
	if (bae%4 ==0 &&bae%100 !=0)||(bae%400 ==0) {
		fmt.Println("是闰年")
	}

	var yuexiaopeng int= 80
	if yuexiaopeng ==100{

	}else if (yuexiaopeng>=80) && (yuexiaopeng<=99){

		fmt.Println("80,99都没有")
	}else if (yuexiaopeng>60)&&(yuexiaopeng<80){

		fmt.Println("60,80都没有")
	}else{
		fmt.Println("什么都没有")
	}
}