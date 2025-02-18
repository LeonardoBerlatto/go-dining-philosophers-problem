package dining

import (
	"sync"
	"testing"
	"time"
)

func philosopherWithMockForks(name string) *Philosopher {
	leftFork := &Fork{ID: 1}
	rightFork := &Fork{ID: 2}
	return NewPhilosopher(name, leftFork, rightFork)
}

func TestPhilosopherFinishesDining(t *testing.T) {
	p := philosopherWithMockForks("Plato")

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go p.Dine(wg)

	wg.Wait()

	if !p.LeftFork.TryLock() {
		t.Errorf("Philosopher should not be holding the left fork")
	}

	if !p.RightFork.TryLock() {
		t.Errorf("Philosopher should not be holding the right fork")
	}
}

func TestPhilosopherThinksAndEats(t *testing.T) {
	p := philosopherWithMockForks("Socrates")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go p.Dine(wg)

	time.Sleep(10 * time.Millisecond)
	if p.LeftFork.TryLock() {
		t.Errorf("Philosopher should be holding the left fork")
	}
	if p.RightFork.TryLock() {
		t.Errorf("Philosopher should be holding the right fork")
	}

	wg.Wait()
}
