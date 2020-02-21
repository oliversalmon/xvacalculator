package xvaCalculator

import (
	"log"
	"os"
	"runtime"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// main is the entry point for the program.
func main() {
	// TODO Implement Application

}
