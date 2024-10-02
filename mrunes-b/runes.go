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
	var r3 rune = '\u0010' /* h 8 i 9 j 10 */
	var r4 rune = '\u0008' /* h 8 */
	var r5 rune = '\r'
	var r6 rune = '\n'
	var r7 rune = '\u0041'

	fmt.Printf(
		"\n  runes.go:\n    %c\n    %c\n%c%c       (newline)%c%c    %c\n",
		r1, r2, r3, r4, r5, r6, r7,
	)

	fmt.Println("       .. and now, call this() in sequence:\n")
	this()
}

/* end. */
