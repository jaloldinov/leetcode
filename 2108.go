package main

import "fmt"

/*
2108. Find First Palindromic String in the Array
Easy
Given an array of strings words, return the first palindromic string in the array. If there is no such string, return an empty string "".

A string is palindromic if it reads the same forward and backward.



Example 1:

Input: words = ["abc","car","ada","racecar","cool"]
Output: "ada"
Explanation: The first string that is palindromic is "ada".
Note that "racecar" is also palindromic, but it is not the first.
Example 2:

Input: words = ["notapalindrome","racecar"]
Output: "racecar"
Explanation: The first and only string that is palindromic is "racecar"
*/

func main() {
	//	Input: words = ["abc","car","ada","racecar","cool"]
	// Output: "ada"
	//["xngla","e","itrn","y","s","pfp","rfd"]

	words := []string{"xngla", "e", "itrn", "y", "s", "pfp", "rfd"}

	fmt.Println(firstPalindrome(words)) // "e"

}
func firstPalindrome(words []string) string {
	for _, word := range words {
		isPalindrome := true
		for i := 0; i < len(word)/2; i++ {
			if word[i] != word[len(word)-i-1] {
				isPalindrome = false
				break
			}
		}
		if isPalindrome {
			return word
		}
	}
	return ""
}

/*
func isPalindrome(word string) bool {
    for i := 0; i < len(word)/2; i++ {
        if word[i] != word[len(word)-i-1] {
            return false
        }
    }
    return true
}

func firstPalindrome(words []string) string {
    for _, word := range words {
        if isPalindrome(word) {
            return word
        }
    }
    return ""
}

*/
