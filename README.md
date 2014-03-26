gogetpanic
==========

```
$ git clone https://github.com/toffaletti/gogetpanic.git
Cloning into 'gogetpanic'...
remote: Counting objects: 9, done.
remote: Compressing objects: 100% (4/4), done.
remote: Total 9 (delta 0), reused 9 (delta 0)
Unpacking objects: 100% (9/9), done.
Checking connectivity... done.
```
```
$ cd gogetpanic/
```
```
$ make
GOPATH=/Users/jason/gogetpanic go get -t -d -v ./...
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xb code=0x1 addr=0x58 pc=0x25b83]

goroutine 1 [running]:
runtime.panic(0x35c9e0, 0x77ff99)
	/usr/local/go/src/pkg/runtime/panic.c:266 +0xb6
main.downloadPackage(0xc210151c00, 0xc210085600, 0xc2100b9c01)
	/usr/local/go/src/cmd/go/get.go:259 +0x33
main.download(0xc2100b9c01, 0x15, 0xc210095600, 0x0)
	/usr/local/go/src/cmd/go/get.go:166 +0x9ea
main.download(0xc2100aa320, 0x14, 0xc210095600, 0xc2100b9301)
	/usr/local/go/src/cmd/go/get.go:232 +0x70a
main.runGet(0x77d100, 0xc21000a050, 0x1, 0x1)
	/usr/local/go/src/cmd/go/get.go:72 +0xc8
main.main()
	/usr/local/go/src/cmd/go/main.go:161 +0x4f1

goroutine 3 [syscall]:
os/signal.loop()
	/usr/local/go/src/pkg/os/signal/signal_unix.go:21 +0x1e
created by os/signal.init·1
	/usr/local/go/src/pkg/os/signal/signal_unix.go:27 +0x31
make: *** [default] Error 2
```
