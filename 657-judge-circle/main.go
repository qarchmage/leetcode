package main

func main() {

}

func judgeCircle(moves string) bool {
	var h, b int

	for _, v := range moves {
		switch v {
		case 'D':
			h--
		case 'U':
			h++
		case 'L':
			b--
		case 'R':
			b++
		}
	}
	if h == 0 && b == 0 {
		return true
	}
	return false
}
