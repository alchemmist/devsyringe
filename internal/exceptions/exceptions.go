package exceptions

import "fmt"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Print(e error) {
	if e != nil {
		fmt.Printf("%v.\n", e)
	}
}
