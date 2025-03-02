package race

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type RPCWithMutex struct {
	result int
	done   chan struct{}
	mu     sync.Mutex
}

func (rpc *RPCWithMutex) compute() {
	time.Sleep(time.Second) // strenuous computation intensifies
	rpc.result = 42
	close(rpc.done)
}

func (RPCWithMutex) version() int {
	return 1 // never going to need to change this
}

func TestRaceRPCWithMutex(t *testing.T) {
	rpc := &RPCWithMutex{done: make(chan struct{})}

	go rpc.compute()         // kick off computation in the background
	version := rpc.version() // grab some other information while we're waiting
	<-rpc.done               // wait for computation to finish

	rpc.mu.Lock()
	result := rpc.result
	rpc.mu.Unlock()

	fmt.Printf("RPC computation complete, result: %d, version: %d\n", result, version)
}
