package main

import "fmt"

func main() {
	fmt.Println("Any stroke i like, but i message you that i say hello")
	kue()
	fmt.Println("another 1")

	for i := 4; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("looped i", i)
		}
	}

	barx()
}

func kue() {
	fmt.Println("this message started from kue")
}

func barx() {
	fmt.Println("and so we exited")
}

//conrol flow
// (1) sequence
// (2) loop; iteractive
// (3) control flow
