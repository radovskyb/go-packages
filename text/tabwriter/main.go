package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func main() {
	// Create a new tabwriter.Writer object
	tw := new(tabwriter.Writer)

	// Initialize `tw` and set it to write to os.Stdout and to format
	// in tab-separated columns with a tab stop of 8 and set it's tab
	// character to `\t`
	//
	// A Writer must be initialized with a call to Init. The first parameter (output)
	// specifies the filter output. The remaining parameters control the formatting:
	//
	//	minwidth	minimal cell width including any padding
	//	tabwidth	width of tab characters (equivalent number of spaces)
	//	padding		padding added to a cell before computing its width
	//	padchar		ASCII char used for padding
	//			if padchar == '\t', the Writer will assume that the
	//			width of a '\t' in the formatted output is tabwidth,
	//			and cells are left-aligned independent of align_left
	//			(for correct-looking results, tabwidth must correspond
	//			to the tab width in the viewer displaying the result)
	//	flags		formatting control
	//
	// func (b *Writer) Init(output io.Writer, minwidth, tabwidth,
	//			padding int, padchar byte, flags uint) *Writer
	tw.Init(os.Stdout, 0, 8, 0, '\t', 0)

	// Write a few tab separated values to `tw` using `tw.Write`
	//
	// Write writes buf to the writer b.
	// The only errors returned are ones encountered
	// while writing to the underlying output stream.
	_, err := tw.Write([]byte("a\tb\tc\td\t.\n"))
	if err != nil {
		log.Fatalln(err)
	}

	// Write another few tab separated values to `tw`, but this time write
	// to `tw` by using `fmt.Fprintln` instead of `tw.Write`
	fmt.Fprintln(tw, "123\t12345\t1234567\t123456789\t.")

	// Flush `tw`
	//
	// Flush should be called after the last call to Write to ensure
	// that any data buffered in the Writer is written to output. Any
	// incomplete escape sequence at the end is considered
	// complete for formatting purposes.
	if err := tw.Flush(); err != nil {
		log.Fatalln(err)
	}

	// Re-initialize `tw` and set it to write to os.Stdout again but this time
	// set it to right-align in space-separated columns of minimal width 5 by
	// setting the flags parameter to tabwriter.AlignRight and at least one blank
	// of padding (so wider column entries do not touch each other).
	tw.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.AlignRight)

	// Write some values to `tw` twice again like above
	fmt.Fprintln(tw, "\na\tb\tc\td\t.")
	fmt.Fprintln(tw, "123\t12345\t1234567\t123456789\t.")

	// Once again call `Flush` on `tw`
	if err := tw.Flush(); err != nil {
		log.Fatalln(err)
	}
}
