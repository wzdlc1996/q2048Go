package main

import (
	"fmt"
	"math/rand"
	"time"
)

type tile struct {
	cont int
}

func (t *tile) getContent() int {
	return int(t.cont)
}

func (t *tile) getRendered() string {
	return fmt.Sprint(t.cont)
}

func (t *tile) init() {
	t.cont = 0
}

func (t *tile) isEmpty() bool {
	return t.cont == 0
}

func (t *tile) randIni() {
	rand.Seed(time.Now().UnixNano())
	if rand.Float32() < 0.5 {
		t.cont = 1
	} else {
		t.cont = 2
	}
}

func (t *tile) isSame(tt *tile) bool {
	return tt.cont == t.cont
}

func (t *tile) merge(tt *tile) bool {
	if t.isEmpty() || t.isSame(tt) {
		if tt.isEmpty() {
			return false
		}
		t.cont += tt.cont
		tt.init()
		return true
	}
	return false
}
