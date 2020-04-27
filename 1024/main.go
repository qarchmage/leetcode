package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("aaa"))
}

func removeDuplicates(S string) string {
	b := []byte(S)
	rem := false
	for {
		for i := 0; i < len(b)-1; {
			fmt.Println(i)
			if b[i] == b[i+1] {
				rem = true
				b = append(b[:i], b[i+2:]...)
				continue
			}
			i++
		}
		if !rem {
			break
		}
		rem = false
	}
	return string(b)
}
