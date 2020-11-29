package main
import (
	"fmt"
	
)
func main(){
	fmt.Println("There is a sign near the entrance that reads no minors")
	var age=21
	var minor=age<18
	fmt.Printf("At age %v, am I a minor? %v\n",age,minor)
}