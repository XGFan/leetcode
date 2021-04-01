package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	return len(s1)+len(s2) == len(s3) && isInterleaveIter(s1, s2, s3)
}

func isInterleaveIter(s1 string, s2 string, s3 string) bool {
	if len(s3) == 0 {
		return true
	}
	if len(s1) == 0 {
		return s2 == s3
	}
	if len(s2) == 0 {
		return s1 == s3
	}
	if s3[0] == s1[0] && s3[0] == s2[0] {
		return isInterleaveIter(s1[1:], s2, s3[1:]) || isInterleaveIter(s1, s2[1:], s3[1:])
	}
	if s3[0] == s1[0] {
		return isInterleaveIter(s1[1:], s2, s3[1:])
	}
	if s3[0] == s2[0] {
		return isInterleaveIter(s1, s2[1:], s3[1:])
	}
	return false
}

func main() {

}
