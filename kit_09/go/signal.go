package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Println("handleSignal() Caught:", signal)
}

func exportMain() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINFO)
/*
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("caught:", sig)
			case syscall.SIGINFO:
				handleSignal(sig)
			}
		}
	}()
*/
	for true {
		fmt.Println(".")
		time.Sleep(20*time.Second)
	}
}
