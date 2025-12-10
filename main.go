package main

import (
	"fmt"

	"github.com/tom-dudley/huffman"
	"github.com/tom-dudley/lz77"
)

const lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut non condimentum felis, ut viverra lorem. Cras laoreet felis non mi egestas, quis vestibulum turpis bibendum. Vivamus ex leo, cursus euismod leo eget, cursus condimentum velit. Sed aliquet, augue in gravida cursus, urna magna lobortis est, in lobortis orci ligula in lorem. Curabitur fringilla consequat mauris in elementum. Interdum et malesuada fames ac ante ipsum primis in faucibus. Suspendisse semper, ligula tempor gravida commodo, nulla mauris mollis metus, a fermentum dolor justo sed arcu. Nam vel varius erat. Sed vehicula aliquet magna. Sed facilisis sem sed ligula tristique mollis."

func Encode(input []byte) []byte {
	fmt.Printf("Input: %d bytes\n", len(input))
	encoded := huffman.Encode(lz77.Encode(input))
	fmt.Printf("Encoded: %d bytes\n", len(encoded))
	return encoded
}

func Decode(input []byte) []byte {
	fmt.Printf("Encoded: %d bytes\n", len(input))
	decoded := lz77.Decode(huffman.Decode(input))
	fmt.Printf("Decoded: %d bytes\n", len(decoded))
	return decoded
}

func main() {
	encoded := Encode([]byte(lorem))
	decoded := Decode(encoded)
	fmt.Println(string(decoded))
}
