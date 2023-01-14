/*
Implement the dining philosopher’s problem with the following constraints
1. There should be 5 philosophers sharing chopsticks, with one chopstick between each
adjacent pair of philosophers.
2. Each philosopher should eat only 3 times
3. The philosophers pick up the chopsticks in any order, not lowest-numbered first
4. In order to eat, a philosopher must get permission from a host which executes in its own
goroutine.
5. The host allows no more than 2 philosophers to eat concurrently.
6. Each philosopher is numbered, 1 through 5.
7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting
to eat <number>” on a line by itself, where <number> is the number of the philosopher.
8. When a philosopher finishes eating (before it has released its locks) it prints “finishing
eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

package main

import (
	"fmt"
	"sync"
)

const numberOfPhilosphers = 5

type chopstick struct {
	sync.Mutex
}
type philosopher struct {
	number         int
	leftChopstick  *chopstick
	rightChopstick *chopstick
}

var host = make(chan struct{}, 2)

func (p philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		host <- struct{}{}
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		fmt.Printf("starting to eat %d\n", p.number)
		fmt.Printf("finishing eating %d\n", p.number)
		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()
		<-host
	}
}

func main() {
	var wg sync.WaitGroup

	chopsticks := make([]*chopstick, numberOfPhilosphers)

	for i := 0; i < 5; i++ {
		chopsticks[i] = new(chopstick)
	}

	philosophers := make([]*philosopher, numberOfPhilosphers)
	for i := 0; i < numberOfPhilosphers; i++ {
		philosophers[i] = &philosopher{
			number:         i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%numberOfPhilosphers],
		}
	}
	wg.Add(numberOfPhilosphers)

	for i := 0; i < numberOfPhilosphers; i++ {
		go philosophers[i].eat(&wg)
	}
	wg.Wait()
}
