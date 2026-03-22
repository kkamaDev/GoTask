// Slicening eng kop takrorlangan elementini toping
//  [1, 1, 1, 2, 1, 1, 1, 2, 2, 2, 2, 3, 4, 5, 6, 6]

package main

import "fmt"

func main() {

	slays := []int{1, 1, 1, 2, 1, 1, 1, 2, 2, 2, 2, 3, 4, 5, 6, 6}

	mapps := takror(slays)
	max, element := engKop(mapps)

	fmt.Println(element, max)

}
func takror(nums []int) map[int]int {
	mapp := map[int]int{}

	for _, key := range nums {

		mapp[key] += 1

	}

	return mapp
}

func engKop(mapp map[int]int) (int, int) {

	koptakrorlanish := 0
	engkopelement := 0

	for kalit, value := range mapp {
		if value > koptakrorlanish {
			koptakrorlanish = value
			engkopelement = kalit
		}
	}
	return koptakrorlanish, engkopelement
}
