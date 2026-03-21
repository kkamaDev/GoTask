// Berilgan 4 xonali sonning juft o’rindagi raqamlari yi’gindisidan toq
// o’rindagi raqamlari yig’indising ayrimasini topuvchi dastur tuzing

package main

import "fmt"

func main() {

	var nums int = 1234

	t1 := nums / 1000
	j1 := nums / 100 % 10
	t2 := nums / 10 % 10
	j2 := nums % 10

	result := (j1 + j2) - (t1 + t2)

	fmt.Println(result)

}
