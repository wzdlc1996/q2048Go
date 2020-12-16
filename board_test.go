package main

import (
	"testing"
)

func Test_board(t *testing.T) {
	var testb board
	testb.init()
	testb.roughRender()

}

func copyAvail(tar map[int]bool) map[int]bool {
	avail := make(map[int]bool)
	for k, v := range tar {
		avail[k] = v
	}
	return avail
}

func Test_randSet(t *testing.T) {
	var testb board
	testb.init()
	avail0 := copyAvail(testb.avail)
	testb.randSet(5)
	avail1 := copyAvail(testb.avail)
	testb.updateAvail()
	avail2 := copyAvail(testb.avail)
	res := true
	for i := 0; i < testb.size; i++ {
		res = res && avail1[i] == avail2[i]
	}
	if !res {
		t.Errorf("Error in randSet_updateAvail")
	}
	res = true
	for i := 0; i < testb.size; i++ {
		res = res && avail0[i] == avail1[i]
	}
	if res {
		t.Errorf("Error in test itself")
	}
}

func Test_tiling(t *testing.T) {
	var testb board
	testb.init()
	testb.randSet(5)
	testb.tilingAlong('b')
	avail1 := copyAvail(testb.avail)
	testb.updateAvail()
	avail2 := copyAvail(testb.avail)
	res := true
	for i := 0; i < testb.size; i++ {
		res = res && avail1[i] == avail2[i]
	}
	if !res {
		t.Errorf("Error in tiling: updateAvail")
	}
}
