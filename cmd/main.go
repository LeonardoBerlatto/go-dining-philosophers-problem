package main

import (
	"dining-philosophers/internal/fork"
	"dining-philosophers/internal/philosopher"
	logger "dining-philosophers/pkg/log"
	"sync"
)

var log = logger.GetLogger()

var philosophersNames = []string{"Aristotle", "Socrates", "Plato", "Heraclitus", "Pythagoras"}

func main() {
	log.Info("Starting the dining philosophers example")
	dine()
	log.Info("The dinner is over!")
}

func dine() {
	var wg sync.WaitGroup
	forks := make([]*fork.Fork, len(philosophersNames))

	for i := 0; i < 5; i++ {
		forks[i] = fork.NewFork(i)
	}

	philosophers := make([]*philosopher.Philosopher, len(philosophersNames))
	for i := 0; i < 5; i++ {
		philosophers[i] = philosopher.NewPhilosopher(
			philosophersNames[i],
			forks[i],
			forks[(i+1)%5],
		)
		wg.Add(1)
		go philosophers[i].Dine(&wg)
	}

	wg.Wait()
}
