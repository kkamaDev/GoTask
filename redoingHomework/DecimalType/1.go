// To’gri to’rtburchakni a va b tomonlari berilgan
// uning yuzini va perimetrini topuvchi dastur tuzing!
//  (s= a * b, p = 2 * (a + b))

package main

import "fmt"

func main() {

	var (
		a int
		b int
	)
	fmt.Scan(&a)
	fmt.Scan(&b)

	yuza := a * b
	peremetr := 2 * (a + b)
	fmt.Printf("to'rtburchak yuzasi %d peremetri: %d\n", yuza, peremetr)
}
