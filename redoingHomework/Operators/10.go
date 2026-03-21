// To’rt xonali a soni berilgan! Jumlani rostilkka tekshiring!
//  A sonining juft o’rindagi raqamlari yigindisi
//   va toq o’rindiadi ramlari yig’indisining ayrimasi 0 ga teng
//   va a soning xar-bir raqami takrorlanmas

package main

import "fmt"

func main() {
	var a int = 1254

	sum := (a/1000+a/10%10)-(a/100%10+a%10) == 0

	takrorlanmas := a/1000 != a/10%10 && a/1000 != a/100%10 && a/100 != a%10

	result := sum && takrorlanmas

	fmt.Println(result)

}
