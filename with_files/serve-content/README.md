# http.ServeContent
<!-- markdownlint-disable -->

```go
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content. io.ReadSeeker)

```

<h2> name string & modtime time.Time params</h2>

We can get these by using:

```go
func (os *File) Stat() (fs.FileInfo, error)
```
After we open the file, we can use the file pointer (*File) for the method .Stat() which returns a FileInfo structure and an error

type FS (file system):

type fs.FileInfo:

```go
type FileInfo interface {
    Name() string //base name of the file
    Size() int64 //length in bytes for regular files; system-dependent for others
    Mode() FileMode //file mode bits
    ModTime() time.Time // modification time
    IsDir() bool //abbreviation for Mode().IsDir()
    Sys() any //underlying data source (can return nil)
}
```

