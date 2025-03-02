package main

import (
	"fmt"
	"strings"
)

type fsmVar int

const (
	success fsmVar = iota
	failure
)

type fsmState int

const (
	s1 fsmState = iota
	s2
	s3
)

func (f *fsm) StartFSM() {
	fmt.Print("type yes to move to next state:")

	for {
		select {
		case v := <-f.v1:

			switch f.curState {
			case s1:
				switch v {
				case success:
					f.nxtState = s2
				case failure:
					f.nxtState = s3

				}
			case s2:
				switch v {
				case success:
					f.nxtState = s3
				case failure:
					f.nxtState = s1

				}
			case s3:
				switch v {
				case success:
					f.nxtState = s1
				case failure:
					f.nxtState = s2
				}
			}
		}
		f.curState = f.nxtState
		fmt.Println("New state", f.curState)
	}
}

type fsm struct {
	v1       chan fsmVar
	curState fsmState
	nxtState fsmState
}

func main() {
	fmt.Println("Starting ...")

	f := fsm{v1: make(chan fsmVar),
		curState: s1,
		nxtState: s2,
	}

	go f.StartFSM()

	var user string
	for {
		fmt.Scan(&user)
		if strings.Compare(user, "yes") == 0 {
			f.v1 <- success
		} else {
			f.v1 <- failure
		}
	}
}
