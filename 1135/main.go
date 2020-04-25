package main

import "fmt"

func main() {
	words := []string{"dyiclysmffuhibgfvapygkorkqllqlvokosagyelotobicwcmebnpznjbirzrzsrtzjxhsfpiwyfhzyonmuabtlwin", "ndqeyhhcquplmznwslewjzuyfgklssvkqxmqjpwhrshycmvrb", "ulrrbpspyudncdlbkxkrqpivfftrggemkpyjl", "boygirdlggnh", "xmqohbyqwagkjzpyawsydmdaattthmuvjbzwpyopyafphx", "nulvimegcsiwvhwuiyednoxpugfeimnnyeoczuzxgxbqjvegcxeqnjbwnbvowastqhojepisusvsidhqmszbrnynkyop", "hiefuovybkpgzygprmndrkyspoiyapdwkxebgsmodhzpx", "juldqdzeskpffaoqcyyxiqqowsalqumddcufhouhrskozhlmobiwzxnhdkidr", "lnnvsdcrvzfmrvurucrzlfyigcycffpiuoo", "oxgaskztzroxuntiwlfyufddl", "tfspedteabxatkaypitjfkhkkigdwdkctqbczcugripkgcyfezpuklfqfcsccboarbfbjfrkxp", "qnagrpfzlyrouolqquytwnwnsqnmuzphne", "eeilfdaookieawrrbvtnqfzcricvhpiv", "sisvsjzyrbdsjcwwygdnxcjhzhsxhpceqz", "yhouqhjevqxtecomahbwoptzlkyvjexhzcbccusbjjdgcfzlkoqwiwue", "hwxxighzvceaplsycajkhynkhzkwkouszwaiuzqcleyflqrxgjsvlegvupzqijbornbfwpefhxekgpuvgiyeudhncv", "cpwcjwgbcquirnsazumgjjcltitmeyfaudbnbqhflvecjsupjmgwfbjo", "teyygdmmyadppuopvqdodaczob", "qaeowuwqsqffvibrtxnjnzvzuuonrkwpysyxvkijemmpdmtnqxwekbpfzs", "qqxpxpmemkldghbmbyxpkwgkaykaerhmwwjonrhcsubchs"}
	chars := "usdruypficfbpfbivlrhutcgvyjenlxzeovdyjtgvvfdjzcmikjraspdfp"
	//words := []string{"cat", "bt", "hat", "tree"}
	//chars := "atach"
	fmt.Println(countCharacters(words, chars))
}

func countCharacters(words []string, chars string) int {
	hashmapuh := make(map[rune]int)
	counter := 0
	c := 0
	for _, char := range chars {
		hashmapuh[char]++
	}
	for _, word := range words {
		c = 0
		hm := copyMap(hashmapuh)
		for _, v := range word {
			if val, ok := hm[v]; ok {
				if val > 0 {
					hm[v]--
					c++
				} else {
					c = 0
					break
				}
			}
		}

		if c == len(word) {
			counter += c
		}
	}
	return counter

}

func copyMap(c map[rune]int) map[rune]int {
	t := make(map[rune]int)
	for k, v := range c {
		t[k] = v
	}
	return t
}

func countCharacters2(words []string, chars string) int {
	alphabetCount := 26
	index := make([]byte, alphabetCount)
	for i := 0; i < len(chars); i++ {
		index[chars[i]-'a'] += 1
	}
	count := 0
	for _, word := range words {
		tmpIndex := make([]byte, alphabetCount)
		for i := 0; i < alphabetCount; i++ {
			tmpIndex[i] = index[i]
		}
		found, n := true, len(word)
		for i := 0; i < n; i++ {
			if tmpIndex[word[i]-'a'] <= 0 {
				found = false
				break
			}
			tmpIndex[word[i]-'a'] -= 1
		}
		if found {
			count += n
		}
	}
	return count
}
