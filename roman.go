package roman

func RomanToDecimal(roman string) int {
    count := 0
	flag := 0
	for i := 0; i < len(roman); i++ {
		if roman[i] == 'I' {
			count += 1
			flag = 1
		} else if roman[i] == 'V' {
			count += 5
			if flag == 1 {
				count -= 2
				flag = 0
			}	
		} else if roman[i] == 'X' {
			count += 10
			if flag == 1 {
				count -= 2
				flag = 0
			}	
		}
    }
	return count
}