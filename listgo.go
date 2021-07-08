
package main
import (
	"fmt"
	"container/list"
)
func main(){
	fmt.Print("请输入值:")
	var input int
	tmplist := list.New()
	for i:=1;i<=10;i++{
	fmt.Scan(&input)
	tmplist.PushBack(input)
}
	
	first :=tmplist.PushFront(0)
	tmplist.Remove(first)

	for l:=tmplist.Front() ; l !=nil; l=l.Next(){

	fmt.Print(l.Value," ")
	}
}
