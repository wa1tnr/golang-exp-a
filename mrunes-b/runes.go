/* runes.go */
/* Wed  2 Oct 13:20:35 UTC 2024 */
/* was: */
/*   Wed  2 Oct 12:19:21 UTC 2024 */
/*   Tue  1 Oct 13:14:22 UTC 2024 */
/* src: google 'go rune literal' with uncredited 'AI Overview' */

/* golang/mrunes-b */
/* variant: fancy printing */

package main

import "fmt"

func main() {
	var r1 rune = 'a'
	var r2 rune = 'π'
	var r3 rune = '\r'
	var r4 rune = '\n'
	var r5 rune = '\u0041'

	fmt.Printf(
		"\n  runes.go:\n    %c\n    %c\n       (newline)%c%c    %c\n",
		r1, r2, r3, r4, r5,
	)

	fmt.Println("       .. and now, call this() in sequence:\n")
	this()
}

/* end. */
