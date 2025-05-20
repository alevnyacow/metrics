package synchronization

type Semaphore struct {
	channel chan struct{}
}

func NewSemaphore(maxRequests uint) *Semaphore {
	return &Semaphore{
		channel: make(chan struct{}, maxRequests),
	}
}

func (semaphore *Semaphore) Request() {
	semaphore.channel <- struct{}{}
}

func (semaphore *Semaphore) Free() {
	<-semaphore.channel
}
