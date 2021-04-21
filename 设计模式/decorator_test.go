package main

import "testing"

func TestFunc2(t *testing.T) {
	worker := &Worker{}
	poolman := &Poolman{}
	poolman.person = worker
	poolman.role()

	richman := &Richman{}
	richman.person = worker
	richman.role()
}
