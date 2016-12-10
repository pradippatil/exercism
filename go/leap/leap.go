package leap

const testVersion = 2

func IsLeapYear(year int) bool {
	if (year%100 != 0 && year%4 == 0) || year%400 == 0 {
		return true
	}
	return false
}
