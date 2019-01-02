package log

import (
	"log"
	"os"
)

// Default logger is golang standard logger.
var DefaultLogger Logger = log.New(os.Stderr, "", log.LstdFlags)

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	DefaultLogger.Print(v...)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	DefaultLogger.Printf(format, v...)
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	DefaultLogger.Println(v...)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	DefaultLogger.Fatal(v...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	DefaultLogger.Fatalf(format, v...)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	DefaultLogger.Fatalln(v...)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	DefaultLogger.Panic(v...)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	DefaultLogger.Panicf(format, v...)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	DefaultLogger.Panicln(v...)
}
