package task2

func IsPositive(a int) string {
	if a > 0 {
		return "Positive"
	} else if a < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}
