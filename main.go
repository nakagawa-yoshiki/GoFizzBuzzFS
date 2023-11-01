package main

import (
	"log"

	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
)

func main() {
	const mountDir = "./mnt"
	const debug = false

	server, err := fs.Mount(mountDir, &FizzBuzzRoot{}, &fs.Options{
		MountOptions: fuse.MountOptions{Debug: debug},
	})
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}
	server.Wait()
}
