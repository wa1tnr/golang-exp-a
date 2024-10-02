/* extra.go */
/* Wed  2 Oct 12:19:21 UTC 2024 */
/* was: */
/*   Tue  1 Oct 13:14:22 UTC 2024 */

package main

import "fmt"

func this() {
	var r1 rune = 'b'
	var r2 rune = 'π'
	var r3 rune = '\n'
	var r4 rune = '\u0046'

	fmt.Printf("extra.go: %c%c%c%c\n", r1, r2, r3, r4)
}

/* end. */
