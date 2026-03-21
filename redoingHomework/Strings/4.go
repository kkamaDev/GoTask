// Berilgan gap nechta so’zdan iborat ekanligini topuvchi dastur tuzing

package main

import (
	"fmt"
	"strings"
)

func main() {

	var soz string = "kamron uyga vazifani bajardi "

	splitsoz := strings.Split(soz, " ")
	fmt.Println(len(splitsoz))
}
