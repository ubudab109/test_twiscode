package main

import (
	"fmt"
	"strings"
)

func main() {
	BinaryToDecimal("1001")
	DecimalToBinary(9)
	fmt.Println(longestPalindrome("aku suka makan nasi"))
}

/*
 Palindrome
*/
func longestPalindrome(s string) string {
	stringMerge := strings.Join(strings.Fields(s), "")
	stringLength := len(stringMerge)

	if stringMerge == "" || stringLength == 0 {
		return stringMerge
	}

	maxLength := 1
	low, high, start := 0, 0, 0

	for i := 1; i < stringLength; i++ {
		// check for even length palindrome substring
		low = i - 1
		high = i

		for low >= 0 && high < stringLength && stringMerge[low] == stringMerge[high] {
			if high-low+1 > maxLength {
				start = low
				maxLength = high - low + 1
			}
			low--
			high++
		}

		// check for odd length palindrome substring
		low = i - 1
		high = i + 1

		for low >= 0 && high < stringLength && stringMerge[low] == stringMerge[high] {
			if high-low+1 > maxLength {
				start = low
				maxLength = high - low + 1
			}
			low--
			high++
		}
	}

	return stringMerge[start : start+maxLength]
}

/*
* Decimal To Binary
 */
func DecimalToBinary(num int) {
	var binary []int // temp binary

	for num != 0 { // looping if input num is not 0
		binary = append(binary, num%2) // append to binary with num mod 2
		num = num / 2                  // calculcate num with divide by 2
	}
	fmt.Print()
	fmt.Println("Input Decimal: ", num)
	if len(binary) == 0 { // check if binary length is not 0
		fmt.Print("0")
	} else {
		fmt.Println("Binary : ")
		for i := len(binary) - 1; i >= 0; i-- { // looping lenght binary with substract 1 and looping until index is more than or equal to 0
			fmt.Print(binary[i]) // get binary
		}
		fmt.Println()
	}

}

/*
* Binary To Decimal
 */
func BinaryToDecimal(number string) {
	// temporary result
	var result int64 = 0

	//temp bit
	var bit int = 0

	// save length of input number binary
	var n int = len(number) - 1

	fmt.Println("Input Binary : ", number)

	// looping input binary string
	for n >= 0 {
		if number[n] == '1' { // check if binary no '1'
			result += (1 << (bit)) // get value of binary
		}
		n = n - 1
		bit++
	}

	fmt.Println("Decimal : ", result)
}
