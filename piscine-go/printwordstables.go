package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(table []string) {
	for _, i := range table {
		for _, j := range i {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
