package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	wid      int     = 4
	hei      int     = 4
	new1prob float32 = 0.5
)

// Board is the basic board framework for 2048 game
type Board struct {
	body  [hei][wid]tile
	size  int
	avail map[int]bool
	index map[int][2]int
	rdex  map[[2]int]int
}

// Init initializes the board with empty tiles
func (b *Board) Init() {
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

// UpdateAvail updates Board.avail, which stores the dictionary of tiles availability(isEmpty)
func (b *Board) UpdateAvail() map[int]bool {
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
	return b.avail
}

func (b *Board) geti(n int) (*tile, bool) {
	ind, ok := b.index[n]
	if !ok {
		return nil, ok
	}
	return &b.body[ind[0]][ind[1]], true
}

func (b *Board) getxy(arr [2]int) (*tile, bool) {
	i, ok := b.rdex[arr]
	if !ok {
		return nil, ok
	}
	return b.geti(i)
}

func (b *Board) getAvail() []int {
	res := make([]int, 0, b.size)
	for k, v := range b.avail {
		if v {
			res = append(res, k)
		}
	}
	return res
}

// RandSet initialize the tiles by randIni randomly
func (b *Board) RandSet(n int) {
	rand.Seed(time.Now().UnixNano())
	_ = b.UpdateAvail()
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

func (b *Board) tilingXY(xy1, xy2 [2]int) bool {
	get1, ok1 := b.getxy(xy1)
	get2, ok2 := b.getxy(xy2)
	if !(ok1 && ok2) {
		return false
	}
	done := get1.merge(get2)
	return done
}

func addArr(xy, arr [2]int) [2]int {
	return [2]int{xy[0] + arr[0], xy[1] + arr[1]}
}

func sliceTileNonEmpty(slc []tile) []tile {
	res := make([]tile, 0, cap(slc))
	for i := range slc {
		if !slc[i].isEmpty() {
			res = append(res, slc[i])
		}
	}
	return res
}

func sliceTiling(slc []tile) ([]tile, bool) {
	res := sliceTileNonEmpty(slc)
	resf := false
	for i := 0; i < len(res)-1; i++ {
		resf = res[i].merge(&(res[i+1])) || resf
	}
	ret := sliceTileNonEmpty(res)
	return ret, resf
}

func (b *Board) tilingAlong(mode byte) bool {
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
	default:
		return false
	}
	for key := range starts {
		tileSlc := make([]tile, 0, wid*hei)
		indSlc := make([]int, 0, wid*hei)
		nowxy := starts[key]
		for true {
			nowt, ok := b.getxy(nowxy)
			if !ok {
				break
			}
			tileSlc = append(tileSlc, *nowt)
			indSlc = append(indSlc, b.rdex[nowxy])
			nowxy = addArr(nowxy, arr)
		}
		tileRes, resf := sliceTiling(tileSlc)
		for j := range indSlc {
			ind := indSlc[j]
			ptr, ok := b.geti(ind)
			if !ok {
				break
			}
			if j < len(tileRes) {
				resf = ptr.copy(tileRes[j]) || resf
			} else {
				ptr.init()
			}
		}
		res = res || resf
	}
	return res
}

// GameLoop is the main loop of the 2048 game
func (b *Board) GameLoop(mode byte) bool {
	res := b.tilingAlong(mode)
	if res {
		rand.Seed(time.Now().UnixNano())
		var n int
		check := rand.Float32()
		if check < new1prob {
			n = 1
		} else {
			n = 2
		}
		b.RandSet(n)
	}
	return res
}

// RoughRender returns a string for render the board roughly
func (b *Board) RoughRender() string {
	z := ""
	for ix := 0; ix < hei; ix++ {
		for iy := 0; iy < wid; iy++ {
			z += fmt.Sprint(b.body[ix][iy].getRendered(), "\t")
		}
		z += "\n"
	}
	return z
}
