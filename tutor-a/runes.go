/* runes.go */
/* Fri  4 Oct 15:41:35 UTC 2024 */
/* was: */
/*   Thu  3 Oct 12:03:40 UTC 2024 */
/*   Wed  2 Oct 13:20:35 UTC 2024 */
/*   Tue  1 Oct 13:14:22 UTC 2024 */
/* src: google 'go rune literal' with uncredited 'AI Overview' */

/* golang/mrunes-c */
/* variant: fancy printing */

package main

import "fmt"

func mainxrescinded() {
	var r1 rune = 'a'
	var r2 rune = 'π'
	var r3 rune = '\n'
	var r4 rune = '\u0041'

	fmt.Printf("\n  runes.go:\n    ")
	fmt.Printf("%c\n    ", r1)
	fmt.Printf("%c\n       ", r2)
	fmt.Printf("(newline)%c    ", r3)
	fmt.Printf("%c", r4)
	newline()
	// fmt.Println("       - - - - -   f e n c e   - - - - -")
	newline()

	fmt.Println("       .. and now, call this() in sequence:\n")
	thisPrintSample()
}

/* end. */
