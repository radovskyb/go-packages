package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	// Start a new process and set its stdout and stderr attributes
	// to os.Stdout and os.Stderr but leave it's stdin to nil since no
	// input will be sent to it
	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{nil, os.Stdout, os.Stderr}
	proc, err := os.StartProcess("/bin/ls", []string{"-la"}, &procAttr)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print the proc's pid to compare it to the procstates pid later on
	fmt.Println(proc.Pid)

	// Wait for the process to exit and return a processes state based on
	// the process that is being waited on
	//
	// Wait waits for the Process to exit, and then returns a ProcessState
	// describing its status and an error, if any. Wait releases any resources
	// associated with the Process. On most operating systems, the Process
	// must be a child of the current process or an error will be returned.
	procState, err := proc.Wait()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print whether the process has exited or not
	fmt.Println(procState.Exited()) // true

	// Print the pid of the exited process
	fmt.Println(procState.Pid())

	// Print the proccess's String function
	fmt.Println(procState.String()) // exit status 0

	// Print whether the process exited successfully
	//
	// Success reports whether the program exited successfully, such as
	// with exit status 0 on Unix.
	fmt.Println(procState.Success()) // true

	// Print the exit information about the process
	//
	// Sys returns system-dependent exit information about the process.
	// Convert it to the appropriate underlying type, such as
	// syscall.WaitStatus on Unix, to access its contents.
	fmt.Println(procState.Sys()) // 0

	// Print the system CPU time of the exited process and it's children
	//
	// SystemTime returns the system CPU time of the exited process
	// and its children.
	fmt.Println(procState.SystemTime().Nanoseconds())

	// Now print the user CPU time of the exited process and it's children
	//
	// UserTime returns the user CPU time of the exited
	// process and its children.
	fmt.Println(procState.UserTime().Nanoseconds())

	// Return the processes resource usage information and then
	// convert it to this computer's resource usage info type using
	// syscall.Rusage
	usage := procState.SysUsage().(*syscall.Rusage)

	// Print a blank line before the rest of the prints below
	fmt.Println()

	// Print all of the usage's values
	fmt.Println(usage.Idrss)
	fmt.Println(usage.Inblock)
	fmt.Println(usage.Isrss)
	fmt.Println(usage.Ixrss)
	fmt.Println(usage.Majflt)
	fmt.Println(usage.Maxrss)
	fmt.Println(usage.Minflt)
	fmt.Println(usage.Msgrcv)
	fmt.Println(usage.Msgsnd)
	fmt.Println(usage.Nivcsw)
	fmt.Println(usage.Nsignals)
	fmt.Println(usage.Nswap)
	fmt.Println(usage.Nvcsw)
	fmt.Println(usage.Oublock)
	fmt.Println(usage.Stime.Nano())
	fmt.Println(usage.Utime.Nano())
}
