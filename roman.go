package roman

const romans = "IVXLCDM"
var decimals = [...]int{1, 5, 10, 50, 100, 500, 1000}

func RomanToDecimal(roman string) int {
	count := 0
	lastval := 0
	for i := len(roman) - 1; i >= 0; i-- {
		for j := 0; j < 7; j++{
			if roman[i] == romans[j]{
				if lastval > decimals[j]{
					count -= decimals[j]
				} else {
					count += decimals[j]
				}
				lastval = decimals[j]
				break
			}
		}
    }
	return count
}