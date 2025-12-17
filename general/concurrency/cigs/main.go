package main

import (
	"fmt"
	"sync"
)

type cigarette struct {
	id         int
	hasPaper   bool
	hasTobacco bool
}
type smoker struct {
	id               int
	smokedCigarettes int
}

func (s smoker) String() string {
	return fmt.Sprintf("Smoker %d has smoked %d cigarettes", s.id, s.smokedCigarettes)
}

func (s *smoker) smoke(cigarettes chan cigarette, wg *sync.WaitGroup) {
	defer wg.Done()

	for range cigarettes {
		s.smokedCigarettes++
	}
}

func genericFactory(genericChan chan int, loopSize int) {
	for i := 0; i < loopSize; i++ {
		genericChan <- i
	}
	close(genericChan)
	fmt.Println("generic factory has been closed")
}

func coordinateProduction(tobaccoCh chan int, paperCh chan int, cigaretteCh chan cigarette) {
	cigaretteIndex := 0
	for {
		_, tobaccoOk := <-tobaccoCh
		_, paperOk := <-paperCh

		if !tobaccoOk && !paperOk {
			break
		}
		cig := cigarette{id: cigaretteIndex}
		cig.hasTobacco = true
		cig.hasPaper = true
		cigaretteIndex++
		cigaretteCh <- cig
	}
	fmt.Println("closing coordinator")
	close(cigaretteCh)
}

func main() {
	tobacco := make(chan int)
	paper := make(chan int)
	cigarettes := make(chan cigarette)
	wg := &sync.WaitGroup{}
	productionGoal := 10000

	go genericFactory(tobacco, productionGoal)
	go genericFactory(paper, productionGoal)

	go coordinateProduction(tobacco, paper, cigarettes)

	smokerCount := 5
	smokers := make([]*smoker, 0, smokerCount)

	for i := 0; i < smokerCount; i++ {
		wg.Add(1)
		s := smoker{id: i}
		smokers = append(smokers, &s)
		go s.smoke(cigarettes, wg)
	}

	wg.Wait()

	smokedTotal := 0
	for s := range smokers {
		smokedTotal += smokers[s].smokedCigarettes
		fmt.Println(smokers[s])
	}
	fmt.Printf("A total of %d has been smoked\n", smokedTotal)

}
