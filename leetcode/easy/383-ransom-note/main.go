package main

func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteMaps := makeHash(ransomNote)
	magazineMaps := makeHash(magazine)

	for k, v := range ransomNoteMaps {
		vv, ok := magazineMaps[k]
		if !ok || v > vv {
			return false
		}
	}

	return true
}

func makeHash(str string) map[rune]int {
	runes := []rune(str)

	maps := make(map[rune]int)
	for _, r := range runes {
		_, ok := maps[r]
		if ok {
			maps[r]++
		} else {
			maps[r] = 1
		}
	}

	return maps
}
