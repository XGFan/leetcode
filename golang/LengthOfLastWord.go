package main

import "fmt"

func lengthOfLastWord(s string) int {
	lastChatIndex := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if lastChatIndex != -1 {
				return lastChatIndex - i
			}
		} else if lastChatIndex == -1 {
			lastChatIndex = i
		}
	}
	return lastChatIndex + 1
}

func main() {
	fmt.Println(lengthOfLastWord("a ") == 1)
	fmt.Println(lengthOfLastWord(" ") == 0)
	fmt.Println(lengthOfLastWord("hello  world") == 5)
}
