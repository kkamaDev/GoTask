// Berilgan 4 xonali sonning juft o’rindagi raqamlari yi’gindisidan toq o’rindagi raqamlari yig’indising ayrimasini topuvchi dastur tuzing

package main

import "fmt"

func main() {

	natija := Fournum(1345)
	fmt.Println(natija)

}

func Fournum(son int) int {

	t1 := son / 1000
	j1 := son / 100 % 10
	t2 := son / 10 % 10
	j2 := son % 10
	result := (j1 + j2) - (t1 + t2)
	return result

}
