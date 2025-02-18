package dining

import (
	logger "dining-philosophers/pkg/log"
	"math/rand"
	"sync"
	"time"
)

type Philosopher struct {
	Name      string
	LeftFork  *Fork
	RightFork *Fork
}

const TimesToEat = 3

var log = logger.GetLogger()

func NewPhilosopher(name string, left, right *Fork) *Philosopher {
	return &Philosopher{Name: name, LeftFork: left, RightFork: right}
}

func (p *Philosopher) Dine(wg *sync.WaitGroup) {
	defer wg.Done()
	for range TimesToEat {
		p.eat()
		p.think()
	}
	log.Info("Philosopher " + p.Name + " has finished dining")
}

func (p *Philosopher) think() {
	log.Info("Philosopher " + p.Name + " is thinking")
	time.Sleep(getThinkingTime())
}

func (p *Philosopher) eat() {
	if p.LeftFork.ID > p.RightFork.ID {
		p.RightFork.Lock()
		log.Info("Philosopher " + p.Name + " has right fork")
		p.LeftFork.Lock()
		log.Info("Philosopher " + p.Name + " has left fork")
	} else {
		p.LeftFork.Lock()
		log.Info("Philosopher " + p.Name + " has left fork")
		p.RightFork.Lock()
		log.Info("Philosopher " + p.Name + " has right fork")
	}

	log.Info("Philosopher " + p.Name + " is eating")
	time.Sleep(getEatingTime())

	p.LeftFork.Unlock()
	p.RightFork.Unlock()
}

func getEatingTime() time.Duration {
	return time.Duration(rand.Intn(1000)+1) * time.Millisecond
}

func getThinkingTime() time.Duration {
	return time.Duration(rand.Intn(1000)+1) * time.Millisecond
}
