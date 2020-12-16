package main

import "fmt"

func main() {
	var testb board
	testb.init()

	z := [4][4]int{
		{0, 0, 2, 0},
		{0, 0, 2, 1},
		{0, 0, 2, 2},
		{0, 0, 0, 0}}

	for ix := 0; ix < 4; ix++ {
		for iy := 0; iy < 4; iy++ {
			testb.body[ix][iy].cont = z[ix][iy]
		}
	}

	//testb.randSet(5)

	testb.roughRender()

	t13, _ := testb.getxy([2]int{1, 3})
	t12, _ := testb.getxy([2]int{1, 2})

	fmt.Println(t13.merge(t12))
	fmt.Println(t13.cont)
	fmt.Println(t12.cont)

	fmt.Println("-------------")

	testb.tilingAlong('r')
	testb.roughRender()

}
