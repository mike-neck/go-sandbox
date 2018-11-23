package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
	"syscall"
)

func main() {
	home := os.Getenv("HOME")
	runApp(home, os.Stdin, os.Stdout, os.Stderr)
}

func runApp(home string, src io.Reader, out io.Writer, eo io.Writer) {
	path := fmt.Sprintf("%s/.sdkman/candidates/java/11.0.0-open/bin/jshell", home)
	command := exec.Command(path)

	stdin, _ := command.StdinPipe()
	stdout, _ := command.StdoutPipe()
	stderr, _ := command.StderrPipe()

	err := command.Start()
	if err != nil {
		log.Fatalln("failed to start jshell", err)
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer func() {
			log.Println("finishing(stdin)")
		}()
		defer wg.Done()
		defer stdin.Close()

		_, err := io.Copy(stdin, src)
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.EPIPE {
			// ignore EPIPE
		} else if err != nil {
			log.Println("error to copy stdin", err)
		}
	}()

	go func() {
		defer func() {
			log.Println("finishing(stdout)")
		}()
		defer wg.Done()
		defer stdout.Close()

		_, err := io.Copy(out, stdout)
		if err != nil {
			log.Println("error on copy stdout", err)
			return
		}
	}()

	go func(w io.Writer) {
		defer func() {
			log.Println("finishing(stderr)")
		}()
		defer wg.Done()
		defer stderr.Close()

		_, err := io.Copy(w, stderr)
		if err != nil {
			log.Println("error on copy stderr", err)
			return
		}
	}(eo)

	wg.Wait()
}
