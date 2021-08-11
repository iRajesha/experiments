package myinterface

import "fmt"

type Writer interface {
	write(data []byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) write(data []byte) (int, error) {
	n, err := fmt.Printf("Writing %v", string(data))
	return n, err
}

func DemonstrateInterfaceUsage() {
	var wr Writer = ConsoleWriter{}
	wr.write([]byte("Good morning Bangalore\n"))
}
