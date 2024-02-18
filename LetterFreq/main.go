package main

import (
	"fmt"
	"math"
)

func equalFrequency(word string) bool {
	length := len(word)
	wordArray := []byte(word)
	//min, max := math.MaxInt, 0

	freqMap := make(map[byte]int, 26)

	for i := 0; i < length; i++ {
		if _, ok := freqMap[wordArray[i]]; !ok {
			freqMap[wordArray[i]] = 1
		} else {
			freqMap[wordArray[i]]++
		}
	}

	freqOfFreq := make(map[int]int, len(freqMap))
	for _, v := range freqMap {
		if _, ok := freqOfFreq[v]; !ok {
			freqOfFreq[v] = 1
		} else {
			freqOfFreq[v]++
		}
	}
	fmt.Printf("freq of Freq:%v\n", freqOfFreq)

	if len(freqOfFreq) == 1 {
		if _, ok := freqOfFreq[1]; ok {
			return true
		}
		for _, v := range freqOfFreq {
			if v == 1 {
				return true
			}
		}

	} else if len(freqOfFreq) == 2 {
		keys, values := make([]int, 0, len(freqOfFreq)), make([]int, 0, len(freqOfFreq))
		for k, v := range freqOfFreq {
			keys = append(keys, k)
			values = append(values, v)
		}

		if _, ok := freqOfFreq[1]; ok {
			if freqOfFreq[1] == 1 {
				return true
			}
		}

		keyVariance := int(math.Abs(float64(keys[0] - keys[1])))
		if keyVariance == 1 {
			maxKey := 0
			for _, k := range keys {
				maxKey = int(math.Max(float64(maxKey), float64(k)))
			}
			if freqOfFreq[maxKey] == 1 {
				return true
			}
		}
	}

	return false
}

func main() {
	word := "abdefgabccdefg"
	//word = "abc"
	//word = "cbbbcccaaa"
	//word = "aabbcc"
	word = "aabbbb"

	retVal := equalFrequency(word)

	fmt.Printf("Is it possile to equalize \"%v\":%v\n", word, retVal)
}
