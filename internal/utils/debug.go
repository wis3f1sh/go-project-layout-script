package utils

import (
	"fmt"
	"io"
)

func Debug(newline bool, a ...any) {
	if newline {
		fmt.Println(a...)
	} else {
		fmt.Print(a...)
	}
}
func Debugf(format string, a ...any) { fmt.Printf(format, a...) }
func Fdebug(w io.Writer, newline bool, a ...any) {
	if newline {
		fmt.Fprintln(w, a...)
	} else {
		fmt.Fprint(w, a...)
	}
}
func Fdebugf(w io.Writer, format string, a ...any) { fmt.Fprintf(w, format, a...) }
