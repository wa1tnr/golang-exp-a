// (blogger said)
//     "..I also re-watched the presentation Rob Pike gave
// for Google Tech Talks on youtube.com.
// In there he presented the following program, chain.go:
// (Formatted with gofmt as it should)

// reference:
// Google I/O 2012 - Go Concurrency Patterns
//  https://youtu.be/f6kdp27TYZs?t=1587
//    see: Daisy-chain (at time index 26:27)
// rob pike

package main

import (
    "flag";
    "fmt";
)

var ngoroutine = flag.Int("n", 100000, "how may")

func f(left, right chan int)    { left <- 1+<-right }

func main() {
    flag.Parse();
    leftmost := make(chan int);

    var left, right chan int = nil, leftmost;
    for i := 0; i < *ngoroutine; i++ {
	    left, right = right, make(chan int);
	    go f(left, right);
    }
    right <- 0;             // bang!
    x := <-leftmost;        // wait for completion
    fmt.Println(x);         // 100000
}

// In this short program we make a chain of 100000 goroutines
// which are connected to each other.
// Each one adds 1 to the value it gets from its right neighbor.
// We start it off by giving the last one (right) a value of 0.
// Then we wait until they are finished and print it.

// Compile and run
// To compile the above program you

// 8g chain.go 
// 8l -o chain chain.8
// And then run it

// ./chain

// vimrc excerpt:
// autocmd Filetype go command! Fmt %!gofmt

// The last line adds a new command which reformats your Go code.
// Just do <ESC>:Fmt and you’ll end up with properly formatted Go code.
// The coding style rules for Go are simple.
// It's what you get when you run it through gofmt.

// nice: ddcb434d129ba935e834cf934b92c15d8aa9e48a :: Subscribe :: Tags :: Categories :: Galleries

// end.
