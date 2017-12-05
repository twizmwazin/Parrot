package world

import (
	"fmt"
	"github.com/tiecoon/minecraft"
	"testing"
)

func TestLoad(t *testing.T) {
	//load world files directory
	path, err := minecraft.NewFilePath("./nw")
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	//create new level
	test2, err := minecraft.NewLevel(path)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	fmt.Println(test2)
	//get chunk from level @ 0,0
	fmt.Println("chunk")
	test3, err := test2.GetChunk(0, 0, false)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	fmt.Println(test3)
	var maxheight int32
	fmt.Println("test")
	//get max height for slice print
	for i := 0; i < 16; i++ {
		if test3.GetHeight(int32(i), 0) > maxheight {
			maxheight = test3.GetHeight(int32(i), 1)
		}
	}
	fmt.Println("test")
	fmt.Println(maxheight)
	mapslice := make([][]uint16, 16)
	//makes array mapslece
	for i := range mapslice {
		mapslice[i] = make([]uint16, maxheight)
	}
	fmt.Println("test")
	fmt.Println(len(mapslice[0]))
	//check cordinates and load in block data
	for i := 0; i < 16; i++ {
		for j := 0; int32(j) < maxheight; j++ {
			block, _ := test2.GetBlock(int32(i), int32(j), 1)
			fmt.Printf("%d %d\n", i, j)
			mapslice[i][j] = block.ID
		}
	}
	fmt.Println("test3")
	fmt.Println(mapslice)
	//print out slice
	for i := maxheight - 1; i > -1; i-- {
		for j := range mapslice {
			fmt.Printf("%2d ", mapslice[j][i])
		}
		fmt.Println()
	}
	fmt.Println("test4")

	//additional debug checks
	fmt.Println(test3.GetHeight(0, 0))
	test4, err := test2.GetBlock(0, 0, 0)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	fmt.Println(test4)
}
