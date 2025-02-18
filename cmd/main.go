package main

import (
	"dining-philosophers/internal/dining"
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
	forks := make([]*dining.Fork, len(philosophersNames))

	for i := 0; i < 5; i++ {
		forks[i] = dining.NewFork(i)
	}

	philosophers := make([]*dining.Philosopher, len(philosophersNames))
	for i := 0; i < 5; i++ {
		philosophers[i] = dining.NewPhilosopher(
			philosophersNames[i],
			forks[i],
			forks[(i+1)%5],
		)
		wg.Add(1)
		go philosophers[i].Dine(&wg)
	}

	wg.Wait()
}
