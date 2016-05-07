package roman

func RomanToDecimal(roman string) int {
    count := 0
	lastval := 0
	romans := len(roman)
	for i := 0; i < romans; i++ {
		if roman[i] == 'I'{
			count += 1
			lastval = 2
		} else if roman[i] == 'V' {
			count += 5 - lastval
			lastval = 0
		} else if roman[i] == 'X' {
			count += 10
			lastval = 20
		} else if roman[i] == 'L' {
			count += 50 - lastval
			lastval = 0
		}
    }
	return count
}