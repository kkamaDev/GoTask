// Berilgan 4 xonali sonning raqamlarini teskari tartibda yozish orqali xosil bo’ladigan sonni toping!

package main

import "fmt"

func main() {

	var numschage = 1234

	minglar := numschage / 1000
	yuzlar := numschage / 100 % 10
	onlar := numschage / 10 % 10
	birlar := numschage % 10

	result := birlar*1000 + onlar*100 + yuzlar*10 + minglar

	fmt.Println(result)

}
