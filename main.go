package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"syscall"

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

type FizzBuzzRoot struct {
	fs.Inode
}

var _ = (fs.NodeOnAdder)((*FizzBuzzRoot)(nil))

func (r *FizzBuzzRoot) OnAdd(ctx context.Context) {
	ch := r.NewPersistentInode(
		ctx,
		&FizzBuzzNode{size: math.MaxInt64},
		fs.StableAttr{Mode: syscall.S_IFREG, Ino: 2},
	)
	r.AddChild("fizzbuzz.txt", ch, false)
}

type FizzBuzzNode struct {
	fs.Inode
	size int64
}

var _ = (fs.NodeGetattrer)((*FizzBuzzNode)(nil))
var _ = (fs.NodeOpener)((*FizzBuzzNode)(nil))
var _ = (fs.NodeReader)((*FizzBuzzNode)(nil))

func (n *FizzBuzzNode) Getattr(ctx context.Context, f fs.FileHandle, out *fuse.AttrOut) syscall.Errno {
	out.Size = uint64(n.size)
	return 0
}

func (*FizzBuzzNode) Open(ctx context.Context, openFlags uint32) (fh fs.FileHandle, fuseFlags uint32, errno syscall.Errno) {
	return nil, fuse.FOPEN_DIRECT_IO, 0
}

func (n *FizzBuzzNode) Read(ctx context.Context, f fs.FileHandle, dest []byte, off int64) (fuse.ReadResult, syscall.Errno) {
	// bs := n.readChars(off, len(dest))
	bs := n.readBytes(off, len(dest))
	return fuse.ReadResultData(bs), fs.OK
}

// func (n *FizzBuzzNode) readChars(off int64, destSize int) []byte {
// 	var bs []byte
// 	for i := 0; i < destSize && off <= n.size-int64(i); i++ {
// 		c := n.charAt(off + int64(i))
// 		bs = append(bs, c)
// 	}
// 	return bs
// }

// func (n *FizzBuzzNode) charAt(index int64) byte {
// 	l := n.lineBy(index)
// 	base := n.fizzbuzzLength(l - 1)
// 	v := n.lineAt(l)
// 	return v[uint64(index)-base]
// }

func (n *FizzBuzzNode) readBytes(off int64, destSize int) []byte {
	bs := []byte{}
	offset := off
	startLine := n.lineBy(offset)

	for i := startLine; ; i++ {
		v := n.lineAt(i)
		if i == startLine {
			pre := int64(n.fizzbuzzLength(startLine - 1))
			v = v[offset-pre:]
		}
		bs = append(bs, []byte(v)...)

		l := int64(len(v))
		if n.size-l < offset {
			bs = bs[:int64(len(bs))-(l-(n.size-offset))]
			break
		}
		offset += l

		if destSize <= len(bs) {
			break
		}
	}

	if destSize < len(bs) {
		bs = bs[:destSize]
	}

	return bs
}

func (n *FizzBuzzNode) lineBy(index int64) int64 {
	idx := uint64(index)
	var l, r int64 = 0, n.size
	for r-l > 1 {
		mid := (l + r) / 2
		v := n.fizzbuzzLength(mid)
		if idx < v {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func (*FizzBuzzNode) lineAt(n int64) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz\n"
	case n%3 == 0:
		return "Fizz\n"
	case n%5 == 0:
		return "Buzz\n"
	default:
		return fmt.Sprintf("%d\n", n)
	}
}

func (n *FizzBuzzNode) fizzbuzzLength(x int64) uint64 {
	const LF = 1
	var bytes, digits, cur uint64 = 0, 0, 0
	for {
		digits++
		pre := cur
		cur = min(cur*10+9, uint64(x))

		fizzbuzz := cur/15 - pre/15
		fizz := cur/3 - pre/3 - fizzbuzz
		buzz := cur/5 - pre/5 - fizzbuzz
		number := cur - pre - fizz - buzz - fizzbuzz

		bytes += number*(digits+LF) + fizz*(4+LF) + buzz*(4+LF) + fizzbuzz*(8+LF)

		if uint64(n.size) <= bytes || uint64(x) <= cur {
			break
		}
	}
	return bytes
}
