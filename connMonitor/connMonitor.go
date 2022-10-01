package connMonitor

import (
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
)

type ConnectionWatcher struct {
	n int64
}

func (cw *ConnectionWatcher) OnStateChange(conn net.Conn, state http.ConnState) {
	switch state {
	case http.StateNew:
		fmt.Println("There is a new connection!")
		cw.Add(1)
	case http.StateHijacked, http.StateClosed:
		fmt.Println("A connection has been closed!")
		cw.Add(-1)
	}
}

// We need atomic counting and adding because of possible multiple connections at the same time changing these values.

func (cw *ConnectionWatcher) Count() int {
	return int(atomic.LoadInt64(&cw.n))
}

func (cw *ConnectionWatcher) Add(c int64) {
	atomic.AddInt64(&cw.n, c)
	fmt.Println("Number of active connections: ", cw.n)
}
