package main

// https://leetcode.com/problems/ransom-note/
func canConstruct(ransomNote string, magazine string) bool {
	magazineHash := make(map[rune]int)
	for _, v := range magazine {
		magazineHash[v] += 1
	}

	for _, v := range ransomNote {
		count, ok := magazineHash[v]
		if ok && count > 0 {
			magazineHash[v]--
		} else {
			return false
		}
	}

	return true
}

//func canConstruct(ransomNote string, magazine string) bool {
//	ransomNoteMaps := makeHash(ransomNote)
//	magazineMaps := makeHash(magazine)
//
//	for k, v := range ransomNoteMaps {
//		vv, ok := magazineMaps[k]
//		if !ok || v > vv {
//			return false
//		}
//	}
//
//	return true
//}
//
//func makeHash(str string) map[rune]int {
//	runes := []rune(str)
//
//	maps := make(map[rune]int)
//	for _, r := range runes {
//		_, ok := maps[r]
//		if ok {
//			maps[r]++
//		} else {
//			maps[r] = 1
//		}
//	}
//
//	return maps
//}
