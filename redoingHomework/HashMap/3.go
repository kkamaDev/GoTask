// Slicening eng kam takrorlangan elementini toping

package main

import "fmt"

func main() {
	slays := []int{1, 1, 2, 3, 3, 3, 4, 4, 4, 4}

	mapps := TakrorDigit(slays)

	kalit, minvalue := engkamTakror(mapps)

	fmt.Printf("eng kam takrorlangan son %d va u %d marta takrorlandi ", kalit, minvalue)

}

func TakrorDigit(arrays []int) map[int]int {

	mapp := map[int]int{}

	for _, qiymat := range arrays {

		mapp[qiymat] += 1
	}

	return mapp

}

func engkamTakror(mapp map[int]int) (int, int) {

	kalit := 0
	engkichikqiymat := 0

	for kalitt, qiymat := range mapp {
		if qiymat > kalit {
			kalitt = qiymat
			engkichikqiymat = kalitt
		}
	}
	return kalit, engkichikqiymat
}
