/*Package main is a really dope package, yo. For realsies tho:

WHY THIS PACKAGE IS DOPE

Cause it is, duh.

*/
package main

import "fmt"

// Meow is an important test identifer for kitties.
//
// Loudness
//
// Loudness of a meow may determine how we categorize a kitty:
//
// 	m0 := Meow{Loudness:5} // medium loudness kitty
//  m1 := Meow{Loudness:1} // low loudness kitty
type Meow struct {
	Loudness uint   // how loud is this meow?
	Sounds   string // how would you spell the sound of this meow?
}

func main() {
	m := Meow{}
	fmt.Printf("%+v", m)
}
