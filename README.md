# GoFizzBuzzFS

```sh
$ go run .
```

```sh
$ ls -l mnt/
total 18428729675200069632
-rw-r--r--  0 root  wheel  9223372036854775807 Jan  1  1970 fizzbuzz.txt

$ ls -lh mnt/
total 18428729675200069632
-rw-r--r--  0 root  wheel   8.0E Jan  1  1970 fizzbuzz.txt

$ head -n 15 mnt/fizzbuzz.txt
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz

$ dd if=mnt/fizzbuzz.txt ibs=512 count=1 skip=999999
1
Fizz
69989923
69989924
FizzBuzz
69989926
69989927
Fizz
69989929
Buzz
Fizz
69989932
69989933
Fizz
Buzz
69989936
Fizz
69989938
69989939
FizzBuzz
69989941
69989942
```

## References

* [FizzBuzz.txt(8エクサバイト)](https://zenn.dev/todesking/articles/c5ee080c6cb4db)
* [go-fuse](https://github.com/hanwen/go-fuse)
