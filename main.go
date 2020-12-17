package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	keyMap := map[byte]byte{'w': 't', 's': 'b', 'a': 'l', 'd': 'r'}
	var mainB Board
	mainB.Init()
	mainB.RandSet(2)
	for true {
		z := mainB.RoughRender()
		fmt.Println(z)
		inp, _ := reader.ReadString('\n')
		mainB.GameLoop(keyMap[inp[0]])
	}
}
