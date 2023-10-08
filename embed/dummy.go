package embed

import (
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/time"
)

type file struct{}

func (*file) Name() string
func (*file) Size() int64
func (*file) Mode() fs.FileMode
func (*file) ModTime() time.Time
func (*file) IsDir() bool
func (*file) Sys() any
func (*file) Type() fs.FileMode
func (*file) Info() (fs.FileInfo, error)

type openFile struct{}

func (*openFile) Seek(offset int64, whence int) (int64, error)
func (*openFile) ReadAt(p []byte, off int64) (n int, err error)
