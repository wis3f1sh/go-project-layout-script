package utils

import (
	"fmt"
	"io"
)

func ToReport(newline bool, a ...any) {
	if newline {
		fmt.Println(a...)
	} else {
		fmt.Print(a...)
	}
}
func ToReportf(format string, a ...any) { fmt.Printf(format, a...) }
func ToFreport(w io.Writer, newline bool, a ...any) {
	if newline {
		fmt.Fprintln(w, a...)
	} else {
		fmt.Fprint(w, a...)
	}
}
func toFreportf(w io.Writer, format string, a ...any) { fmt.Fprintf(w, format, a...) }
