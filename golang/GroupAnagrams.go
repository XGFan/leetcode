package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	cache := make([][2]int, 0)
	for _, s := range strs {
		test(s, &result, &cache)
	}
	return result
}

func test(x string, y *[][]string, cache *[][2]int) {
	equal := false
	var used, count int
	for _, c := range x {
		used |= 1 << (c - 96)
		count += int(c - 96)
	}
	for i := range *y {
		if len(x) == len((*y)[i][0]) {
			ints := (*cache)[i]
			if used == ints[0] && count == ints[1] {
				equal = true
				(*y)[i] = append((*y)[i], x)
				break
			}
		}
	}
	if !equal {
		*y = append(*y, []string{x})
		*cache = append(*cache, [...]int{used, count})
	}
}

func printBigArray(f [][]string) {
	for _, a := range f {
		if len(a) > 1 {
			sort.Strings(a)
			fmt.Println(a)
		}
	}
}

type Wrapper [][]string

func (p Wrapper) Len() int {
	return len(p)
}

func (p Wrapper) Less(i, j int) bool {
	return p[i][0] <= p[j][0]
}

func (p Wrapper) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	//fmt.Printf("%v\n", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

	inputJson, _ := os.OpenFile("go/Input.json", os.O_RDONLY, 0644)
	input := make([]string, 0)
	json.NewDecoder(inputJson).Decode(&input)

	anagrams := groupAnagrams(input)
	for _, v := range anagrams {
		sort.Strings(v)
	}
	anagramsWrapper := Wrapper(anagrams)
	sort.Sort(anagramsWrapper)

	expectResultJson, _ := os.OpenFile("go/ExpectResult.json", os.O_RDONLY, 0644)
	expectResult := make([][]string, 0)
	json.NewDecoder(expectResultJson).Decode(&expectResult)
	for _, v := range expectResult {
		sort.Strings(v)
	}
	expectResultWrapper := Wrapper(expectResult)
	sort.Sort(expectResultWrapper)

	//printBigArray(expectResultWrapper)
	//fmt.Println()
	//printBigArray(anagramsWrapper)
	fmt.Println(hash("angling"))
	fmt.Println(hash("gillian"))
	fmt.Println(hash("passport"))
	fmt.Println(hash("taproots"))
	fmt.Println(hash("pettiness"))
	fmt.Println(hash("snippiest"))
}

//[angling gillian]
//[passport taproots]
//[pettiness snippiest]
func hash(s string) (int, int) {
	var used, count int
	for _, c := range s {
		if used&(1<<(c-96)) > 0 {
			count += int(c-96) * int(c-96)
		} else {
			count += int(c - 96)
		}
		used |= 1 << (c - 96)
	}
	return used, count
}
