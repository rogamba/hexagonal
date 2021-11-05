package main

import "fmt"

func main() {
	app := NewApp()

	errs := make(chan error, 2)
	errs = app.RunAsync()
	fmt.Printf("Terminated %s", <-errs)
}
