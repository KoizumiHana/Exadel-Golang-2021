package number

import "math"

func isAutomorphic(number int) bool {
	square := int(math.Pow(float64(number), 2))
	if number < 0 {
		return false
	}
	for number > 0 {
		if number%10 != square%10 {
			return false
		}
		number /= 10
		square /= 10
	}
	return true
}
