package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
)

var (
	doors = [3]bool{false, false, false}

	myChoiceWin    int
	otherChoiceWin int
)

func RandomIndex() uint32 {
	b := make([]byte, 4)
	rand.Read(b)

	out := uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])

	if out >= 0 && out < 1431655765 {
		return 0
	} else if out >= 1431655765 && out < 2863311531 {
		return 1
	} else {
		return 2
	}
}

func main() {
	var recursion int

	if len(os.Args) >= 2 {
		recursion, _ = strconv.Atoi(os.Args[1])
	} else {
		recursion = 100000
	}

	for i := 0; i < recursion; i++ {
		doors = [3]bool{false, false, false}
		doors[RandomIndex()] = true

		myChoice := RandomIndex()

		if doors[myChoice] == true {
			myChoiceWin++
		} else {
			otherChoiceWin++
		}

	}
	fmt.Printf("Ran: %v times\n", recursion)

	fmt.Println()

	fmt.Printf("My choice wins: %v\n", myChoiceWin)
	fmt.Printf("Other choice wins: %v\n", otherChoiceWin)

	fmt.Println()

	fmt.Printf("My choice win rate: %v\n", float64(myChoiceWin)/float64(myChoiceWin+otherChoiceWin))
	fmt.Printf("Other choice win rate: %v\n", float64(otherChoiceWin)/float64(myChoiceWin+otherChoiceWin))
}
