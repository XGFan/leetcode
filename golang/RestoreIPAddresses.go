package main

func restoreIpAddresses(s string) []string {
	strings := make([]string, 0)
	makeUp(&strings, "", s, 4)
	return strings
}

func makeUp(ips *[]string, prefix, s string, p int) {
	if p < 0 {
		return
	}
	if len(s) > 3*p || len(s) < p {
		return
	}
	if len(s) == 0 {
		if p == 0 {
			*ips = append(*ips, prefix)
			return
		} else {
			return
		}
	}
	if prefix != "" {
		prefix += "."
	}
	switch s[0] {
	case '0':
		makeUp(ips, prefix+"0", s[1:], p-1)
	case '1':
		makeUp(ips, prefix+"1", s[1:], p-1)
		if len(s) >= 2 {
			makeUp(ips, prefix+s[:2], s[2:], p-1)
		}
		if len(s) >= 3 {
			makeUp(ips, prefix+s[:3], s[3:], p-1)
		}
	case '2':
		makeUp(ips, prefix+"2", s[1:], p-1)
		if len(s) >= 2 {
			makeUp(ips, prefix+s[:2], s[2:], p-1)
		}
		if len(s) >= 3 && (s[1] < '5' || (s[1] == '5' && s[2] <= '5')) {
			makeUp(ips, prefix+s[:3], s[3:], p-1)
		}
	default:
		makeUp(ips, prefix+s[:1], s[1:], p-1)
		if len(s) >= 2 {
			makeUp(ips, prefix+s[:2], s[2:], p-1)
		}
	}
}

func main() {
	//restoreIpAddresses("0000")
	restoreIpAddresses("25525511135")
}
