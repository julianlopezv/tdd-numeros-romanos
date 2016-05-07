package roman

func RomanToDecimal(roman string) int {
    count := 0
	lastval := 0
	romans := len(roman)
	for i := romans - 1; i >= 0; i-- {
		if roman[i] == 'I'{
			if lastval > 1{
				lastval = -1
			} else {
				lastval = 1
			}
		} else if roman[i] == 'V' {
			lastval = 5
		} else if roman[i] == 'X' {
			if lastval > 10{
				lastval = -10
			} else {
				lastval = 10
			}
		} else if roman[i] == 'L' {
			lastval = 50
		} else if roman[i] == 'C' {
			if lastval > 100{
				lastval = -100
			} else {
				lastval = 100
			}
		}
		count += lastval
    }
	return count
}