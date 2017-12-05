package world

import (
	"compress/gzip"
	"fmt"
	"github.com/MJKWoolnough/minecraft/nbt"
	"os"
)

func Load(file string) (rt nbt.Tag, err error) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	gzipReader, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	d := nbt.NewDecoder(gzipReader)
	fmt.Println(d)
	rt, err = d.Decode()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}
