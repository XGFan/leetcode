package main

import "fmt"

func fullJustify(words []string, maxWidth int) []string {
	count := -1
	f := 0
	strings := make([]string, 0)
	for i, v := range words {
		//fmt.Println("count: ", count)
		if count+len(v)+1 <= maxWidth {
			count += 1 + len(v)
		} else {
			bytes := make([]byte, 0, maxWidth)
			var base, ext int
			if i-f-1 == 0 {
				base = maxWidth - count
				bytes = append(bytes, words[i-1]...)
				for k := 0; k < base; k++ {
					bytes = append(bytes, ' ')
				}
			} else {
				base = (maxWidth - count) / (i - f - 1)
				ext = (maxWidth - count) % (i - f - 1)
				for j := 0; j < i-f-1; j++ {
					bytes = append(bytes, words[f+j]...)
					//base only
					for k := 0; k <= base; k++ {
						bytes = append(bytes, ' ')
					}
					if j < ext {
						bytes = append(bytes, ' ')
					}
				}
				bytes = append(bytes, words[i-1]...)
			}
			strings = append(strings, string(bytes))
			//fmt.Println(maxWidth-count, i-f-1, base, ext)
			//fmt.Println(string(bytes))
			count = len(v)
			f = i
		}
	}
	bytes := make([]byte, 0, maxWidth)
	for i := f; i < len(words)-1; i++ {
		bytes = append(bytes, words[i]...)
		bytes = append(bytes, ' ')
	}
	bytes = append(bytes, words[len(words)-1]...)
	for len(bytes) < maxWidth {
		bytes = append(bytes, ' ')
	}
	strings = append(strings, string(bytes))
	//fmt.Println(string(bytes))
	return strings
}

func main() {
	for _, s := range fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16) {
		fmt.Println(s)
	}
	for _, s := range fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16) {
		fmt.Println(s)
	}
}

//This    is   an
//example  of text
//justification.
