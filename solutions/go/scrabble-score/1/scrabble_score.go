package scrabble

import "strings"

func Score(word string) int {
	points := 0
    one := "aeioulnrst"
    two := "dg"
    three := "bcmp"
    four := "fhvwy"
    five := "k"
    eight := "jx"
    ten := "qz"
	for i := range word {
        word = strings.ToLower(word)
    	if strings.Contains(one, string(word[i])) {
            points++
        } else if strings.Contains(two, string(word[i])) {
            points += 2
        }else if strings.Contains(three, string(word[i])) {
            points += 3
        }else if strings.Contains(four, string(word[i])) {
            points += 4
        }else if strings.Contains(five, string(word[i])) {
            points += 5
        }else if strings.Contains(eight, string(word[i])) {
            points += 8
        }else if strings.Contains(ten, string(word[i])) {
            points += 10
        }
    }
    return points
}
