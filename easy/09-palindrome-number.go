package easy

var x = -121

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	result := rev(x, 0)

	if x == result {
		return true
	} else {
		return false
	}
}

func rev(n int, temp int) int {
	if n == 0 {
		return temp
	}

	temp = (temp * 10) + (n % 10)
	return rev(n/10, temp)
}
