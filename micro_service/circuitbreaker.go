package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/cclehui/study_golang/micro_service/circuit"
)

var count int = 0

func main() {

	cb := circuit.NewThresholdBreaker(5)

	for {
		if cb.Ready() {
			// Breaker is not tripped, proceed
			err := doSomething()
			if err != nil {
				cb.Fail() // This will trip the breaker once it's failed 10 times
				continue
			}
			cb.Success()
		} else {
			// Breaker is in a tripped state.
			fmt.Println("Breaker is in a tripped stat")
		}

		time.Sleep(time.Second * 1)

	}

}

func doSomething() error {
	count++
	fmt.Println("doSomething, ", count)

	if count < 6 {
		return errors.New("error xxxx")
	}

	if count > 8 {
		return errors.New("error xxxx")

	}

	return nil

}
