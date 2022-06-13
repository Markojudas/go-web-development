# io.Copy

Here I am trying to my luck reading the docs. I will try to use the io.Copy method to read from stdin and copy to stdout

io.Copy
func Copy

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

dst Writer: we need a Type Writer:
whatever type has the "Write method" it is also type Writer

```go
type Writer interface{
    Write(p []byte) (n int, err error)
}
```

src Reader: we need a Type Reader:
whatever type has the "Read method" is is also type Reader

```go
type Reader interface{
    Read(p []byte) (n int, err error)
}
```

Standard Input and Standard Output :

from the os package:

```go
// os package's variables:
var (
    Stdin = NewFile(uintptr(syscall, Stdin), "/dev/stdin")
    Stdout = NewFile(uintptr(syscall, Stdout), "/dev/stdout")
    Stderr = NewFile(uintptr(syscall, Stderr), "/dev/stdout")
)

func NewFile(fd uintptr, name string) *File // Returns a pointer to a File

//Type File has the method Write and Read!
//from os.File
func (*File) Read(b []byte) (n int, err error)
func (*File) Write(b []byte) (n int, err error)
```
