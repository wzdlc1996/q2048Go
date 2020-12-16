package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	wid int = 4
	hei int = 4
)

type board struct {
	body  [hei][wid]tile
	size  int
	avail map[int]bool
	index map[int][2]int
	rdex  map[[2]int]int
}

func (b *board) init() {
	i := 0
	b.avail = make(map[int]bool)
	b.index = make(map[int][2]int)
	b.rdex = make(map[[2]int]int)
	for iy := 0; iy < hei; iy++ {
		for ix := 0; ix < wid; ix++ {
			b.body[iy][ix].init()
			b.avail[i] = true
			b.index[i] = [2]int{ix, iy}
			b.rdex[[2]int{ix, iy}] = i
			i++
		}
	}
	b.size = i
}

func (b *board) updateAvail() {
	for ix := 0; ix < wid; ix++ {
		for iy := 0; iy < hei; iy++ {
			i := b.rdex[[2]int{ix, iy}]
			if b.body[ix][iy].isEmpty() {
				b.avail[i] = true
			} else {
				b.avail[i] = false
			}
			i++
		}
	}
}

func (b *board) geti(n int) (*tile, bool) {
	ind, ok := b.index[n]
	if !ok {
		return nil, ok
	}
	return &b.body[ind[0]][ind[1]], true
}

func (b *board) getxy(arr [2]int) (*tile, bool) {
	i, ok := b.rdex[arr]
	if !ok {
		return nil, ok
	}
	return b.geti(i)
}

func (b *board) getAvail() []int {
	res := make([]int, 0, b.size)
	for k, v := range b.avail {
		if v {
			res = append(res, k)
		}
	}
	return res
}

func (b *board) randSet(n int) {
	rand.Seed(time.Now().UnixNano())
	temp := b.getAvail()

	var size int
	if n > len(temp) {
		size = len(temp)
	} else {
		size = n
	}
	sele := rand.Perm(len(temp))[:size]
	for j := range sele {

		site := temp[sele[j]]
		b.avail[site] = false
		item, _ := b.geti(site)
		item.randIni()
	}
}

func (b *board) tilingAlong(mode byte) bool {
	res := false
	var starts [][2]int
	var arr [2]int
	switch mode {
	case 'r':
		arr = [2]int{0, -1}
		starts = make([][2]int, hei)
		for i := 0; i < hei; i++ {
			starts[i] = [2]int{i, wid - 1}
		}
	case 'l':
		arr = [2]int{0, 1}
		starts = make([][2]int, hei)
		for i := 0; i < hei; i++ {
			starts[i] = [2]int{i, 0}
		}
	case 'b':
		arr = [2]int{-1, 0}
		starts = make([][2]int, wid)
		for i := 0; i < wid; i++ {
			starts[i] = [2]int{hei - 1, i}
		}
	case 't':
		arr = [2]int{1, 0}
		starts = make([][2]int, wid)
		for i := 0; i < wid; i++ {
			starts[i] = [2]int{0, i}
		}
	}
	for key := range starts {
		nowxy := starts[key]
		nexxy := [2]int{nowxy[0] + arr[0], nowxy[1] + arr[1]}
		for true {
			var nex *tile
			var okex bool
			for true {
				nex, okex = b.getxy(nexxy)
				if !okex {
					break
				}
				if nex.isEmpty() {
					nexxy = [2]int{nexxy[0] + arr[0], nexxy[1] + arr[1]}
					continue
				}
				break
			}
			if !okex {
				break
			}
			now, _ := b.getxy(nowxy)

			// but here. can be now==nex. And now [2, 0, 2, 0] ->r [0, 0, 2, 2]
			done := now.merge(nex)

			if done {
				b.avail[b.rdex[nowxy]] = false
				b.avail[b.rdex[nexxy]] = true
			}
			res = res || done
			nowxy = [2]int{nowxy[0] + arr[0], nowxy[1] + arr[1]}
		}
	}
	return res
}

func (b *board) gameLoop(mode byte) bool {
	return false
}

func (b *board) roughRender() {
	for ix := 0; ix < hei; ix++ {
		for iy := 0; iy < wid; iy++ {
			fmt.Print(b.body[ix][iy].getRendered(), "\t")
		}
		fmt.Println()
	}

}
