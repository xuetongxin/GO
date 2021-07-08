package main
import "fmt"
func counter(initial int) func () int{
	if initial<0{
		initial=0
	}

	return func() int{
	initial++
	
	return initial;
	}
}

func main(){
	c1:=counter(1)
	fmt.Println(c1())
	fmt.Println(c1())

	c2:=counter(10)
	fmt.Println(c2())
	fmt.Println(c2())
}
