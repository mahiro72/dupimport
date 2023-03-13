package dupimport

import (
	"fmt"
	fmt2 "fmt" // want "WARNING: fmt is duplicated import"
)

func main() {
	fmt.Println("hello")
	fmt2.Println("hello2")
}
