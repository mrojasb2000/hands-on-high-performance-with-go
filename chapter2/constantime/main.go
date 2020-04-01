package main

import "fmt"

func main() {
	words := [3]string{"foo", "bar", "baz"}
	fmt.Println(words[1]) // This reference the string in position 1 in the array, "bar"
}
