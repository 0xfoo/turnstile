package main

import (
	"errors"
	"fmt"
)

type Action int

const (
	coin Action = iota
	push Action = iota
)

type TurnstileState interface {
	Handle() (TurnstileState, error)
}

type TurnstileContext struct {
	action       Action
	currentState TurnstileState
}

type UnlockedState struct {
	context TurnstileContext
}

type LockedState struct {
	context TurnstileContext
}

func (us *UnlockedState) Handle() (TurnstileState, error) {
	return handle(us.context)
}

func (ls *LockedState) Handle() (TurnstileState, error) {
	return handle(ls.context)
}

func handle(context TurnstileContext) (TurnstileState, error) {
	switch context.action {
	case coin:
		return &UnlockedState{context}, nil
	case push:
		return &LockedState{context}, nil
	default:
		return nil, errors.New("unexpected action")
	}
}

func main() {
	fmt.Printf("Turnstile\n")
}
