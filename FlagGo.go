 package main
 import "fmt"
 import "flag"
func main(){
	name:=flag.String("name","wang","你的姓")
	flag.Parse()
	fmt.Printf("%v",*name)
}
