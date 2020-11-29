package main
import (
	"fmt"
	"strings"
)
func main(){
	fmt.Println("You find yourself in a dimly lit cavren.")
	var command="walk outside"
	var exit=strings.Contains(command, "walk outside")
	fmt.Println("You leave the cave:",exit)

	//qc 3.1
	//1
	var wearShades=true
	//2
	
	wearShades=strings.Contains(command,"read")
}