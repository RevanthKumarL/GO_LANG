package main
import (
	"fmt"
)

//multiple init() function
func init(){
	fmt.Println("Welcome to init() function")
}
func init(){
	fmt.Println("Hello! init() function")
}


func main() {
	fmt.Println("Welcome to main() function")

}