package main
import "fmt"

func main() {
	var command ="go east"
	if command =="go east" {
		fmt.Println("You head further up rhe mountain.")

	}else if command=="go inside"{
		fmt.Println("You enterthe cave where you live out the rest of your life.")

	}else{
		fmt.Println("Didn't quite get that.")
	}

	//Qc 3.3
	var room = "entrance"
	if room == "cave" {
		fmt.Println("You live in a cave.")
	}else if room == "mountain" {
		fmt.Println("You a mountain guy.")
	}else if room=="entrance" {
		fmt.Println("You found it, you'll go to next level.")
	}else {
		fmt.Println("Didn't quite get that.")
	}

}