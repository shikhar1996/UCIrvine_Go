package main

import (
	"fmt"
	"sync"
)

type chopS struct {
	sync.Mutex
}
type philosopher struct {
	lChop, rChop *chopS
	id           int
	countEat     int
}

func initialize() ([]*chopS, []*philosopher) {

	cSticks := make([]*chopS, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(chopS)
	}
	fmt.Println(cSticks)

	philosophers := make([]*philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &philosopher{lChop: cSticks[i%5], rChop: cSticks[(i+1)%5], id: i + 1}
	}

	return cSticks, philosophers
}

func getPermission(permission chan bool, request chan bool) {

	i := 1
	for j := range request {
		i++
		// fmt.Println("Currently Eating ", len(request), i)
		permission <- j
	}
}

func (p philosopher) eat(permission chan bool, request chan bool, wg *sync.WaitGroup) {

	request <- true
	permissionValue := <-permission
	if permissionValue {
		if p.id%2 == 0 {
			p.lChop.Lock()
			p.rChop.Lock()
			fmt.Println("Starting to eat ", p.id)
			p.rChop.Unlock()
			p.lChop.Unlock()
			fmt.Println("Finishing eating ", p.id)
			wg.Done()
		} else {
			p.rChop.Lock()
			p.lChop.Lock()
			fmt.Println("Starting to eat ", p.id)
			p.lChop.Unlock()
			p.rChop.Unlock()
			fmt.Println("Finishing eating ", p.id)
			wg.Done()
		}
	}

}
func main() {
	var wg sync.WaitGroup
	permission := make(chan bool, 2)
	request := make(chan bool, 2)
	_, p := initialize()
	fmt.Println(p)

	wg.Add(15)
	go getPermission(permission, request)

	for i := 0; i < 15; i++ {
		j := i % 5
		go p[j].eat(permission, request, &wg)
	}
	wg.Wait()
	fmt.Println("Execution successfully completed.")
}
