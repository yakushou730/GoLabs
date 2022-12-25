package main

func backspaceCompare(s string, t string) bool {
	return runString(s) == runString(t)
}

func runString(s string) string {
	var res []rune
	for _, ch := range s {
		if len(res) == 0 {
			if ch != '#' {
				res = append(res, ch)
			}
			continue
		}

		if ch == '#' {
			res = res[:len(res)-1]
		} else {
			res = append(res, ch)
		}
	}

	return string(res)
}
