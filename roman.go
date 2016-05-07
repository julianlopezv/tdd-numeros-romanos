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
			if i < romans - 1 && roman[i+1] == 'L'{
				count -= 10
			} else {
				count += 10
			}	
		} else if roman[i] == 'L' {
			count += 50
		}
    }
	return count
}