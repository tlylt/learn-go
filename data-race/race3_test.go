package race

import (
	"fmt"
	"testing"
	"time"
)

type RPCNoWait struct {
	result int
	done   chan struct{}
}

func (rpc *RPCNoWait) compute() {
	time.Sleep(time.Second) // strenuous computation intensifies
	rpc.result = 42
	close(rpc.done)
}

func (RPCNoWait) version() int {
	return 1 // never going to need to change this
}

func TestRaceRPCNoWait(t *testing.T) {
	rpc := &RPCNoWait{done: make(chan struct{})}

	go rpc.compute() // kick off computation in the background
	// version := rpc.version() // grab some other information while we're waiting
	// <-rpc.done               // wait for computation to finish

	result := rpc.result

	fmt.Printf("RPC computation complete, result: %d\n", result)
}
