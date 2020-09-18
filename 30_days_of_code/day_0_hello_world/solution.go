/**
 * Author: Zarandi David (Azuwey)
 * Date: 09/18/2020
 * Source: HackerRank - https://www.hackerrank.com/challenges/30-hello-world
 **/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello, World.")
	fmt.Println(input)
}
