package main

func validTobogganPws(pws []pwTest) int {
	validCounter := 0
	for _, pw := range pws {
		if len(pw.pw) < pw.max || len(pw.pw) < pw.min {
			continue
		}

		test1 := string([]rune(pw.pw)[pw.min-1]) == pw.char
		test2 := string([]rune(pw.pw)[pw.max-1]) == pw.char

		if test1 != test2 {
			validCounter++
		}
	}
	return validCounter
}
