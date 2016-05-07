package roman

func RomanToDecimal(roman string) int {
    count := 0
	romans := len(roman)
	for i := 0; i < romans; i++ {
		if roman[i] == 'I'{
			if i < romans - 1 && roman[i+1] != 'I'{
				count -= 1
			} else {
				count += 1
			}
		} else if roman[i] == 'V' {
			count += 5
		} else if roman[i] == 'X' {
			count += 10	
		}
    }
	return count
}