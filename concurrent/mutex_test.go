package concurrent

import (
	"sync"
	"testing"
)

var (
	mutex   sync.Mutex
	balance int
	ch1     = make(chan struct{}, 1)
)

func deposit1(amount int) {
	ch1 <- struct{}{}
	balance += amount
	<-ch1
}

func balance1() int {
	ch1 <- struct{}{} // acquire token
	b := balance
	<-ch1 // release token
	return b
}

func deposit2(amount int) {
	mutex.Lock()
	balance += amount
	mutex.Unlock()
}

func balance2() int {
	mutex.Lock()
	b := balance
	mutex.Unlock()
	return b
}

func TestMutex1(t *testing.T) {
	ch1 <- struct{}{}

}
