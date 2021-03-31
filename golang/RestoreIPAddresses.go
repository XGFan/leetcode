package main

func restoreIpAddresses(s string) []string {
	strings := make([]string, 0)
	makeUp(&strings, []byte{}, s, 4)
	return strings
}

func makeUp(ips *[]string, prefix []byte, s string, p int) {
	if p < 0 || len(s) > 3*p || len(s) < p {
		return
	}
	if p == 0 {
		if len(s) == 0 {
			*ips = append(*ips, string(prefix))
		}
		return
	}
	if len(prefix) != 0 {
		prefix = append(prefix, '.')
	}
	switch s[0] {
	case '0':
		makeUp(ips, append(prefix, s[0]), s[1:], p-1)
	case '1':
		makeUp(ips, append(prefix, s[0]), s[1:], p-1)
		if len(s) >= 2 {
			makeUp(ips, append(prefix, s[:2]...), s[2:], p-1)
		}
		if len(s) >= 3 {
			makeUp(ips, append(prefix, s[:3]...), s[3:], p-1)
		}
	case '2':
		makeUp(ips, append(prefix, s[0]), s[1:], p-1)
		if len(s) >= 2 {
			makeUp(ips, append(prefix, s[:2]...), s[2:], p-1)
		}
		if len(s) >= 3 && (s[1] < '5' || (s[1] == '5' && s[2] <= '5')) {
			makeUp(ips, append(prefix, s[:3]...), s[3:], p-1)
		}
	default:
		makeUp(ips, append(prefix, s[0]), s[1:], p-1)
		if len(s) >= 2 {
			makeUp(ips, append(prefix, s[:2]...), s[2:], p-1)
		}
	}
}

func main() {
	//restoreIpAddresses("0000")
	restoreIpAddresses("25525511135")
}
