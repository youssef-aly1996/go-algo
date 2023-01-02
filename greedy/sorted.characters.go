package greedy

import (
	"fmt"
	"sort"
)

type charFreqs struct {
	char string
	frequency int
}

func SortCharactersFrequency(msg string) []charFreqs {
	var mapMsg = make(map[byte]int)
	var sortedFrequency []charFreqs
	for i := 0; i < len(msg); i++ {
		mapMsg[msg[i]]++
	}
	for k, v := range mapMsg {
		cf := charFreqs {string(k), v}
		sortedFrequency = append(sortedFrequency, cf)
	}
	sort.SliceStable(sortedFrequency, func(i, j int) bool {
		return sortedFrequency[i].frequency < sortedFrequency[j].frequency
	})

	for _, v := range sortedFrequency {
		fmt.Print(v.char)
	}

	return sortedFrequency
}


func CharactersFrequency(msg string)  {
	asciiCodes := make([]int, 127)
	msgRunes := []rune(msg)
	
	for i:=0; i< len(msgRunes); i++ {
		current_code := msgRunes[i]
		asciiCodes[current_code]++
	}
	for i, v := range asciiCodes {
		if v > 0 {
			c := string(rune(i))
			fmt.Print(c + " ")
			fmt.Println(v)
		}
	}
}