package main

import (
	"fmt"
	"strings"
)

// returns 2 strings
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	//checking the length of the initial slice
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	//in case the length is not greater than 1, then it will get an underscore( _ ) back as the 2nd name
	return initials[0], "_"
}

func main() {
	fn1, sn1 := getInitials("cloud strife")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("tifa lockhart")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("aerith")
	fmt.Println(fn3, sn3)

	fn4, sn4 := getInitials("barret")
	fmt.Println(fn4, sn4)
}
