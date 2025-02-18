package dining

import "sync"

type Fork struct {
	ID int
	sync.Mutex
}

func NewFork(id int) *Fork {
	return &Fork{ID: id}
}
