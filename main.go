package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// func main() {
// 	ch := make(chan os.Signal, 1)
// 	signal.Notify(ch, syscall.SIGPIPE)
// 	go func() {
// 		<-ch
// 		fmt.Println("got pipe")
// 		os.Exit(0)
// 	}()

// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// }

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGPIPE)
	go func() {
		<-ch
		fmt.Fprintf(os.Stderr, "got SIGPIPE\n")
		os.Exit(0)
	}()

	// Give the child time to terminate:
	time.Sleep(1 * time.Second)
	_, err := fmt.Printf("Hello, world!\n")
	if err == nil {
		fmt.Fprintf(os.Stderr, "Printf succeeded\n")
	} else {
		fmt.Fprintf(os.Stderr, "Printf failed with %v\n", err)
	}
	// Allow time for any possible signals to be handled:
	time.Sleep(1 * time.Second)
	fmt.Fprintf(os.Stderr, "Bye!\n")
}
