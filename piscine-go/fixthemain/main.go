package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(s string) {
	myRune := []rune(s)
	for _, r := range myRune {
		z01.PrintRune(r)
	}
	z01.PrintRune(10)
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.state = true
	return true
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = false
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	return ptrDoor.state
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state
}

func main() {
	var door Door
	OpenDoor(&door)
	IsDoorClose(&door)
	IsDoorOpen(&door)
	CloseDoor(&door)
}
