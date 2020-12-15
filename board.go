package main

import (
	"math/rand"
	"time"
)

const (
	wid int = 4
	hei int = 4
)

type board struct {
	body  [wid][hei]tile
	avail map[int]bool
	index map[int][2]int
}

func (b *board) init() {
	i := 0
	for ix := 0; ix < wid; ix++ {
		for iy := 0; iy < hei; iy++ {
			b.body[ix][iy].init()
			b.avail[i] = true
			b.index[i] = [2]int{ix, iy}
			i++
		}
	}
}

func (b *board) updateAvail() {
	i := 0
	for ix := 0; ix < wid; ix++ {
		for iy := 0; iy < hei; iy++ {
			if b.body[ix][iy].isEmpty() {
				b.avail[i] = true
			} else {
				b.avail[i] = false
			}
			i++
		}
	}
}

func (b *board) geti(n int) *tile {
	ind := b.index[n]
	return &b.body[ind[0]][ind[1]]
}

func (b *board) randSet(n int) {
	rand.Seed(time.Now().UnixNano())
	temp := [wid * hei]int{0}
	i := 0
	for k, v := range b.avail {
		if v {
			temp[i] = k
			i++
		} else {
			continue
		}
	}
	for j := range rand.Perm(i)[:n] {
		site := temp[j]
		b.avail[site] = false
		b.geti(site).randIni()
	}
}

func (b *board) tilingLeft() {
	for iy := 0; iy < wid; iy++ {

	}
}
