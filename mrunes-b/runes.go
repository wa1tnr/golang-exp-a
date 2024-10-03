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

/* one flaw one note:
 * ^M visible at end of '(newline)' utterance as in '(newline)^M'
 * the dot in 'runes.go' (printed) aligns with the ( of  '(newline)'
 */

/* 0x0d 0x0a  cr lf */
/* 13   10 */

func main() {
	var r1 rune = 'a'
	var r2 rune = 'π'
	var r3 rune = '\n'
	var r4 rune = '\u0041'

	fmt.Printf("\n  runes.go:\n    ",)

	fmt.Printf(
                                      "%c\n    ", r1,)
	fmt.Printf(
                                              "%c\n       ", r2,)
	fmt.Printf(
                                                         "(newline)%c    ", r3,)
	fmt.Printf(
                                                                        "%c\n",r4,)
//	fmt.Printf("\n  runes.go:\n    %c\n    %c\n       (newline)%c    %c\n",
	fmt.Printf("\n  runes.go:\n    %c\n    %c\n       (newline)%c    %c\n",
		r1, r2, r3, r4,
	)

	fmt.Println("       .. and now, call this() in sequence:\n")
	this()
}

/* end. */
