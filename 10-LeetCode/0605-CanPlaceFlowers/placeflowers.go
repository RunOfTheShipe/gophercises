package placeflowers

func canPlaceFlowers(flowerbed []int, n int) bool {

	var last = -1
	var next = 1
	var length = len(flowerbed)

	// base cases
	if n == 0 {
		return true
	} else if length == 0 {
		return n == 0
	} else if length == 1 {
		return flowerbed[0] == 0 && n-1 == 0
	}

	for current := 0; current < length; current++ {
		if current == 0 {
			// first item in the array
			if flowerbed[current] == 0 && flowerbed[next] == 0 {
				flowerbed[current] = 1
				n--
			}
		} else if next == length {
			// end of the array
			if flowerbed[last] == 0 && flowerbed[current] == 0 {
				flowerbed[current] = 1
				n--
			}
		} else {
			// middle of the array somewhere
			if flowerbed[last] == 0 && flowerbed[current] == 0 && flowerbed[next] == 0 {
				flowerbed[current] = 1
				n--
			}
		}

		// if all the flowers have been placed, break out early
		if n == 0 {
			break
		}

		last++
		next++
	}

	return n == 0
}
