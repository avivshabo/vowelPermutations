package main

import (
	"fmt"
)

/*
Given an integer n, your task is to count how many strings of length n can be formed under the following rules:

Each character is a lower case vowel ('a', 'e', 'i', 'o', 'u')
Each vowel 'a' may only be followed by an 'e'.
Each vowel 'e' may only be followed by an 'a' or an 'i'.
Each vowel 'i' may not be followed by another 'i'.
Each vowel 'o' may only be followed by an 'i' or a 'u'.
Each vowel 'u' may only be followed by an 'a'.
Since the answer may be too large, return it modulo 10^9 + 7.



Example 1:

Input: n = 1
Output: 5
Explanation: All possible strings are: "a", "e", "i" , "o" and "u".
Example 2:

Input: n = 2
Output: 10
Explanation: All possible strings are: "ae", "ea", "ei", "ia", "ie", "io", "iu", "oi", "ou" and "ua".
Example 3:

Input: n = 5
Output: 68


Constraints:

1 <= n <= 2 * 10^4
*/

//define the allowed next characters for each vowel
var rules = map[byte][]byte{
	'a': {'e'},
	'e': {'a', 'i'},
	'i': {'a', 'e', 'o', 'u'},
	'o': {'i', 'u'},
	'u': {'a'},
}

var iterations = 0

func countVowelPermutation(n int) int {
	//The first letter can be any vowel
	all_vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	return calcPerms(all_vowels, n)
}

//calcPerms is a recursive function that receives a bytes slice (vowels) and an int (n)
//The function will return the number of combinations that an n characters string can be created
//following the vowels rules defined in the question

func calcPerms(vowels []byte, n int) int {
	perms := 0
	iterations++
	if n == 1 {
		//if n is 1, then we only need to check how many options we have for the next letter.
		perms = len(vowels)
	} else {
		//iterate over every char in the received list of vowels
		for _, char := range vowels {
			//run the function recursively on the vowels allowed after the current char, reduce required length by 1.
			perms += calcPerms(rules[char], n-1)
		}
	}
	return perms
}

func main() {
	n := 7
	fmt.Println("Started with N=", n)
	perms := countVowelPermutation(n)
	fmt.Println("Permutations: ", perms, " Iterations: ", iterations)
}
