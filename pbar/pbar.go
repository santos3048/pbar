package pbar

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"sync"
)

const Char = "—" 

type Bar struct {
	mu      sync.Mutex
	ct      int
	size    int
	message string
	ch      string
	keep    bool
}

func Create(msg string, size int) Bar {
	var outMode uint32
	out := windows.Handle(os.Stdout.Fd())
	windows.GetConsoleMode(out, &outMode)
	outMode |= windows.ENABLE_PROCESSED_OUTPUT | windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	windows.SetConsoleMode(out, outMode)
	return Bar{
		ct:      0,
		size:    size,
		message: msg,
		ch:      "/",
		keep:    true,
	}
}

func (b *Bar) Print() {
	b.mu.Lock()
	b.keep = true
	b.mu.Unlock()
	for b.keep {
		b.mu.Lock()
		pt := "\r[\x1b[36m"
		i := 50 * b.ct / b.size
		for z := 1; z <= i; z++ {
			pt += Char
		}
		pt += "\x1b[0m"
		for z := 1; z <= 50-i; z++ {
			pt += Char
		}
		pt = pt + "] (" + fmt.Sprint(b.ct) + "/" + fmt.Sprint(b.size) + ") " + b.message + " " + b.ch
		switch b.ch {
		case "/":
			b.ch = "—"
		case "—":
			b.ch = "\\"
		case "\\":
			b.ch = "|"
		case "|":
			b.ch = "/"
		}
		b.mu.Unlock()
		fmt.Print(pt)
	}
}

func (b *Bar) Msg(msg string) {
	b.mu.Lock()
	b.message = msg
	b.mu.Unlock()
}

func (b *Bar) Up() {
	b.mu.Lock()
	b.ct += 1
	b.mu.Unlock()
}

func (b *Bar) Finish(message string) {
	b.mu.Lock()
	b.keep = false
	pt := "\r[\x1b[32m"
	for z := 1; z <= 50; z++ {
		pt += Char
	}
	pt += "\x1b[0m] ("
	sz := fmt.Sprint(b.size)
	pt += sz
	pt += "/"
	pt += sz
	pt += ") "
	pt += "\x1b[1m"
	pt += message
	pt += "\x1b[0m"
	fmt.Println(pt)
	b.mu.Unlock()
}

func (b *Bar) Stop() {
	b.mu.Lock()
	b.keep = false
	b.mu.Unlock()
}
