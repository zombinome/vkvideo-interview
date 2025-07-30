package logging

import "fmt"

type ConsoleLogger struct {
}

func (cl *ConsoleLogger) Error(message string, err error) {
	fmt.Print(message + ": " + err.Error())
}
