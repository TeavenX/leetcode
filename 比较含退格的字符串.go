package main

import "fmt"

func main() {
	s, t := "bxj##tw", "bxo#j##tw"
	s, t = "ab#c", "ad#c"
	//s, t = "ab##", "c#d#"
	fmt.Println(backspaceCompare(s, t))
}
func backspaceCompareError(s string, t string) bool {
	ns, nt := len(s), len(t)
	ps, pt := ns-1, nt-1
	backspaceS, backspaceT := 0, 0
	for ps >= 0 && pt >= 0 {
		for ps >= 0 && s[ps] == '#' {
			backspaceS++
			ps--
		}
		for pt >= 0 && t[pt] == '#' {
			backspaceT++
			pt--
		}
		for ps >= 0 && backspaceS >= 0 {
			if s[ps] == '#' {
				backspaceS++
			} else {
				backspaceS--
			}
			ps--
		}
		for pt >= 0 && backspaceT >= 0 {
			if t[pt] == '#' {
				backspaceT++
			} else {
				backspaceT--
			}
			pt--
		}
		if ps >= 0 && pt >= 0 {
			if s[ps] == t[pt] {
				ps--
				pt--
			} else {
				return false
			}
		}
	}
	if ps >= 0 || pt >= 0 {
		return false
	}
	return true
}

func backspaceCompare(s, t string) bool {
	skipS, skipT := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}

func backspaceCompare(s string, t string) bool {
	ps, pt := len(s)-1, len(t)-1
	for ps >= 0 || pt >= 0 {
		psb, ptb := 0, 0
		for ps >= 0 && s[ps] == '#' {
			psb++
			ps--
		}
		for ps >= 0 && psb > 0 {
			ps--
			psb--
			for ps >= 0 && s[ps] == '#' {
				psb++
				ps--
			}
		}
		for pt >= 0 && t[pt] == '#' {
			ptb++
			pt--
		}
		for pt >= 0 && ptb > 0 {
			ptb--
			pt--
			for pt >= 0 && t[pt] == '#' {
				ptb++
				pt--
			}
		}
		if ps >= 0 && pt >= 0 {
			if s[ps] != t[pt] {
				return false
			} else {
				ps--
				pt--
			}
		} else if ps >= 0 || pt >= 0 {
			return false
		}
	}
	if (ps < 0 && pt < 0) || s[:ps] == t[:pt] {
		return true
	}
	return false
}
