package main

import (
	"fmt"
)

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (s ServerState) getName() string {
	return stateName[s]
}

func transtions(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
func main() {
	ns := transtions(StateIdle)
	fmt.Println(ns.getName())

	ns2 := transtions(ns)
	fmt.Println(ns2.getName())
}

//connected
//idle
