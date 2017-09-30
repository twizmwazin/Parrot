package world

import (
	"compress/gzip"
	"fmt"
	"github.com/MJKWoolnough/minecraft/nbt"
	"os"
)

func load(file string) (d nbt.Decoder, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	gzipReader, err := gzip.NewReader(f)
	if err != nil {
		return
	}
	d = nbt.NewDecoder(gzipReader)
	rt, err := d.Decode()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(rt)
	return
}
