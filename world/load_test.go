package world

import (
	"fmt"
	"github.com/tiecoon/minecraft"
	"testing"
)

func TestLoad(t *testing.T) {
	path, err := minecraft.NewFilePath("./nw")
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	test2, err := minecraft.NewLevel(path)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	fmt.Println(test2)
	fmt.Println("chunk")
	test3, err := test2.GetChunk(0, 0, false)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	fmt.Println(test3)
	var maxheight int32
	fmt.Println("test")
	for i := 0; i < 16; i++ {
		if test3.GetHeight(int32(i), 0) > maxheight {
			maxheight = test3.GetHeight(int32(i), 1)
		}
	}
	fmt.Println("test")
	fmt.Println(maxheight)
	mapslice := make([][]uint16, 16)
	for i := range mapslice {
		mapslice[i] = make([]uint16, maxheight)
	}
	fmt.Println("test")
	fmt.Println(len(mapslice[0]))
	for i := 0; i < 16; i++ {
		for j := 0; int32(j) < maxheight; j++ {
			block, _ := test2.GetBlock(int32(i), int32(j), 1)
			fmt.Printf("%d %d\n", i, j)
			mapslice[i][j] = block.ID
		}
	}
	fmt.Println("test3")
	fmt.Println(mapslice)
	for i := maxheight - 1; i > -1; i-- {
		for j := range mapslice {
			fmt.Printf("%2d ", mapslice[j][i])
		}
		fmt.Println()
	}
	fmt.Println("test4")

	fmt.Println(test3.GetHeight(0, 0))

	test4, err := test2.GetBlock(0, 0, 0)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	fmt.Println(test4)
}
