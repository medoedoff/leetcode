package easy

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	myNum := 5

	if num > myNum {
		return -1
	} else if num < myNum {
		return 1
	} else {
		return 0
	}
}

func guessNumber(n int) int {
	low := 1
	high := n
	result := binarySearch(low, high)
	return result
}

func binarySearch(low int, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	result := guess(mid)
	if result == 0 {
		return mid
	} else if result < 0 {
		high = mid - 1
	} else {
		low = mid + 1
	}
	return binarySearch(low, high)
}
