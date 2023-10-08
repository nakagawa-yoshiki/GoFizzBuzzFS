package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestFizzBuzzNode_fizzbuzzLength(t *testing.T) {
	_fizzbuzzLength := func(n int64) uint64 {
		fizz, buzz, fizzbuzz, number, bytes := 0, 0, 0, 0, 0
		for i := 1; i <= int(n); i++ {
			if i%15 == 0 {
				fizzbuzz++
				bytes += 8
			} else if i%3 == 0 {
				fizz++
				bytes += 4
			} else if i%5 == 0 {
				buzz++
				bytes += 4
			} else {
				number++
				bytes += len(fmt.Sprint(i))
			}
			bytes++ // LF
		}
		return uint64(bytes)
	}

	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "0", args: args{n: 0}, want: 0},
		{name: "1", args: args{n: 1}, want: 2},
		{name: "2", args: args{n: 2}, want: 4},
		{name: "3", args: args{n: 3}, want: 9},
		{name: "15", args: args{n: 15}, want: _fizzbuzzLength(15)},
		{name: "123456", args: args{n: 123456}, want: _fizzbuzzLength(123456)},
		{name: "max", args: args{n: 1e18}, want: 12674074074074074068},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &FizzBuzzNode{size: math.MaxInt64}
			if got := n.fizzbuzzLength(tt.args.n); got != tt.want {
				t.Errorf("FizzBuzzNode.fizzbuzzLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFizzBuzzNode_lineBy(t *testing.T) {
	type args struct {
		index int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		/*
			1: 1.      0  1
			2: 2.      2  3
			3: Fizz.   4  5  6  7  8
		*/
		{name: "0", args: args{0}, want: 1},
		{name: "1", args: args{1}, want: 1},
		{name: "2", args: args{2}, want: 2},
		{name: "3", args: args{3}, want: 2},
		{name: "4", args: args{4}, want: 3},
		{name: "8", args: args{8}, want: 3},
		{name: "max", args: args{math.MaxInt64}, want: 729002457810002753},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &FizzBuzzNode{size: math.MaxInt64}
			if got := n.lineBy(tt.args.index); got != tt.want {
				t.Errorf("FizzBuzzNode.lineBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFizzBuzzNode_lineAt(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{1}, want: "1\n"},
		{name: "3(Fizz)", args: args{3}, want: "Fizz\n"},
		{name: "5(Buzz)", args: args{5}, want: "Buzz\n"},
		{name: "15(FizzBuzz)", args: args{15}, want: "FizzBuzz\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &FizzBuzzNode{size: math.MaxInt64}
			if got := n.lineAt(tt.args.n); got != tt.want {
				t.Errorf("FizzBuzzNode.lineAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFizzBuzzNode_ReadBytes(t *testing.T) {
	type args struct {
		off      int64
		destSize int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "1.2.Fizz.", args: args{off: 0, destSize: 9}, want: []byte("1\n2\nFizz\n")},
		{name: "1.2.Fi", args: args{off: 0, destSize: 6}, want: []byte("1\n2\nFi")},
		{name: "zz.4.Buzz", args: args{off: 6, destSize: 9}, want: []byte("zz\n4\nBuzz")},
		{name: "maxSize", args: args{off: math.MaxInt64 - 40, destSize: 65535}, want: []byte("Fizz\n729002457810002752\n7290024578100027")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &FizzBuzzNode{size: math.MaxInt64}
			if got := n.readBytes(tt.args.off, tt.args.destSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FizzBuzzNode.ReadBytes() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

// func TestFizzBuzzNode_charAt(t *testing.T) {
// 	type args struct {
// 		index int64
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want byte
// 	}{
// 		/*
// 		   1: 0 1
// 		   2: 2 3
// 		   Fizz: 4 5 6 7 8
// 		   4: 9 10
// 		   Buzz: 11 12 13 14 15
// 		*/
// 		{name: "index:0", args: args{0}, want: '1'},
// 		{name: "index:2", args: args{2}, want: '2'},
// 		{name: "index:4", args: args{4}, want: 'F'},
// 		{name: "index:11", args: args{11}, want: 'B'},
// 		{name: "index:15", args: args{15}, want: '\n'},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			n := &FizzBuzzNode{size: math.MaxInt64}
// 			if got := n.charAt(tt.args.index); got != tt.want {
// 				t.Errorf("FizzBuzzNode.charAt() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
