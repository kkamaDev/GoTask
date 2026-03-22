//  /**  Slicening har bir elementi nechtadadan takrorlanganligini topuvchi dasturtuzing
//  [1, 1, 1, 2, 1, 1, 1, 2, 2, 2, 2, 3, 4, 5, 6, 6]

// [1:6 2:5 3:1 4:1 5:1 6:2]
// /* *

package main

import "fmt"

func main() {
	var nums []int = []int{1, 1, 1, 1, 2, 4, 4, 5, 5, 9, 9, 8, 3, 3, 14, 15}

	result := Retake(nums)
	fmt.Println(result)

}

func Retake(nums []int) map[int]int {
	mapp := map[int]int{}

	for _, key := range nums {

		mapp[key] += 1

	}

	return mapp
}
