/*Ryan Wendling
10/ 06/ 2017
dotProduct.RW.go
The purpose of this program is to take into two matrix files and perform dot notation to print out the new combined matrix */

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

// Some error handling.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

/* ReadInts reads whitespace-separated ints from a line of one of our input matrix strings. The found ints will be appended to an int slice
where all the slices will eventually be appended to our new 2d slice matrix. The 2d slice matrix will then be returned. */
func ReadInts(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var resultTwo [][]int

	for scanner.Scan() {
		var result []int
		line := scanner.Text()
		//fmt.Println("this line  " + string(line))
		// Split the line on spaces.
		parts := strings.Split(line, " ")

		// Loop over the parts from the string.
		for i := range parts {
			//fmt.Println(parts[i])
			x, err := strconv.Atoi(parts[i])
			if err != nil {
				return resultTwo, err
			}
			result = append(result, x)
		}
		//fmt.Println(result)
		resultTwo = append(resultTwo, result)
		//result = append(result, x)
	}
	return resultTwo, scanner.Err()
}

/* This is where we ask the program user for the names of two matrix files. The program will then attempt to open them, convert them into a string,
and call our ReadInts function. After doing the dot notation logic we save our results to a new 2d slice and print out the contents. */
func main() {

	// We take in two seperate matrix files. ex: "matrix1.txt" and "matrix2.txt". note: they will be used in the order they are read.
	fmt.Printf("please enter the name of the first matrix. \n")

	var input string
	fmt.Scanln(&input)

	fmt.Printf("please enter the name of the second matrix. \n")

	var input2 string
	fmt.Scanln(&input2)

	dat, err := ioutil.ReadFile(input)
	check(err)
	//fmt.Print(string(dat))
	s := string(dat)

	dat2, err := ioutil.ReadFile(input2)
	check(err)
	//fmt.Print(string(dat2))
	s2 := string(dat2)

	// Turns our matrix strings into 2d slices
	ints, err := ReadInts(strings.NewReader(s))
	fmt.Printf("Printing first 2d slice from ReadInts function. \n")
	fmt.Println(ints, err)

	ints2, err := ReadInts(strings.NewReader(s2))
	fmt.Printf("Printing second 2d slice from ReadInts function. \n")
	fmt.Println(ints2, err)

	fmt.Println("Matching lengths of the two input matrices are:  ")
	fmt.Println(len(ints[0]))
	fmt.Println(len(ints2))

	// Return error if columns of matrix1 does not match rows of matrix2.
	if len(ints[0]) != len(ints2) {
		fmt.Println("matrices dont match, can't compute dot notation")
		return
	}

	// Our 2d slice that we will fill and print later on in the program.
	var combMatrix [][]int

	// Goes through rows of matrix1
	for i := 0; i < len(ints); i++ {
		var combResult []int

		// This allows the rows of matrix1 to be reused for every column of matrix 2
		for k := 0; k < len(ints2[0]); k++ {
			combVal := 0

			//This should take the row i of matrix1 and combine it with the next column of matrix 2
			for j := 0; j < len(ints[0]); j++ {
				combVal += ints[i][j] * ints2[j][k]

			}
			combResult = append(combResult, combVal)
		}
		combMatrix = append(combMatrix, combResult)
	}
	// Successful results of our program are printed here.
	fmt.Println("Here is our new matrix using dot notation:  ")
	fmt.Println(combMatrix)
}
