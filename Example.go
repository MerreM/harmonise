// Example
package harmonise

import (
	"fmt"
)

func main() {
	n1 := Create(1)
	n2 := Create(2)
	n1.Join(&n1)
	n2.Join(&n1)
	fmt.Println("Done")
}
