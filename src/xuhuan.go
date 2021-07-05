package main

import "fmt"

func sum(n int )int{

	var sum1 int
	var  i int
	for i=1;i<=n;i++{
		sum1 =sum1+i
	}
	return sum1
}
func main()  {
	var n int
	fmt.Print("请输入n的值\n")
	fmt.Scan(&n)
	fmt.Print(sum(n))
}
