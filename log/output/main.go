package main

import "log"

func main() {
	log.SetFlags(log.Llongfile)

	// Output writes the output for a logging event. The string s contains
	// the text to print after the prefix specified by the flags of the
	// Logger. A newline is appended if the last character of s is not
	// already a newline. Calldepth is the count of the number of
	// frames to skip when computing the file name and line number
	// if Llongfile or Lshortfile is set; a value of 1 will print the details
	// for the caller of Output.
	log.Output(0, "log.Output - flag: Llongfile, calldepth: 0\n\n")
	log.Output(1, "log.Output - flag: Llongfile, calldepth: 1\n\n")
	log.Output(2, "log.Output - flag: Llongfile, calldepth: 2\n\n")
	log.Output(3, "log.Output - flag: Llongfile, calldepth: 3\n\n")

	log.SetFlags(log.Lshortfile)

	// We can make calldepth as deep as needed if log.Output is wrapped
	// in x amount of function calls etc. The greater the abstraction before
	// calling log.Output, the greater the call depth can be
	log.Output(0, "log.Output - flag: Lshortfile, calldepth: 0\n\n")
	log.Output(1, "log.Output - flag: Lshortfile, calldepth: 1\n\n")
	log.Output(2, "log.Output - flag: Lshortfile, calldepth: 2\n\n")
	log.Output(3, "log.Output - flag: Lshortfile, calldepth: 3")
}
